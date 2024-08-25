package sp

import (
	"context"
	"errors"

	abci "github.com/cometbft/cometbft/abci/types"
	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/gogoproto/proto"
	"github.com/forbole/bdjuno/v4/database/models"
	sptypes "github.com/forbole/bdjuno/v4/modules/sp/types"
	vgtypes "github.com/forbole/bdjuno/v4/modules/virtualgroup/types"
	junotypes "github.com/forbole/juno/v5/types"
	"github.com/rs/zerolog/log"
)

var (
	EventCreateStorageProvider = proto.MessageName(&sptypes.EventCreateStorageProvider{})
	EventEditStorageProvider   = proto.MessageName(&sptypes.EventEditStorageProvider{})
	EventSpStoragePriceUpdate  = proto.MessageName(&sptypes.EventSpStoragePriceUpdate{})
	EventCompleteSpExit        = proto.MessageName(&vgtypes.EventCompleteStorageProviderExit{})
)

var StorageProviderEvents = map[string]bool{
	EventCreateStorageProvider: true,
	EventEditStorageProvider:   true,
	EventSpStoragePriceUpdate:  true,
	EventCompleteSpExit:        true,
}

// HandleBlock implements modules.BlockModule
func (m *Module) HandleBlock(
	block *tmctypes.ResultBlock, results *tmctypes.ResultBlockResults, txs []*junotypes.Tx, vals *tmctypes.ResultValidators,
) error {
	// actual use "block.Block.Height == 1"
	if block.Block.Height%10 == 0 {
		req := query.PageRequest{
			Key:        nil,
			Offset:     0,
			Limit:      100,
			CountTotal: false,
			Reverse:    false,
		}
		sps, pageResponse, err := m.source.StorageProviders(block.Block.Height, req) // actual use 1 instead of block.Block.Height
		if err == nil && len(sps) > 0 {
			log.Error().Str("module", "sp").Str("page response", pageResponse.String()).Str("first storage provide", sps[0].String())
		}
	}
	ctx := context.Background()
	statements, err := m.ExportEventsInTxs(ctx, block, txs)
	if err != nil {
		return err
	}
	return m.db.ExecuteStatements(statements)
}

// ExportEventsInTxs accepts a slice of events in tx in order to save in database.
func (m Module) ExportEventsInTxs(ctx context.Context, block *tmctypes.ResultBlock, txs []*junotypes.Tx) (map[string][]interface{}, error) {
	allSQL := make(map[string][]interface{}, 0)
	for _, tx := range txs {
		sqls, err := m.ExtractEvent(ctx, block, tx)
		if err != nil {
			log.Err(err)
			continue
		}
		for k, v := range sqls {
			allSQL[k] = v
		}
	}
	return allSQL, nil
}

// ExtractEvent accepts the transaction and handles events contained inside the transaction.
func (m *Module) ExtractEvent(ctx context.Context, block *tmctypes.ResultBlock, tx *junotypes.Tx) (map[string][]interface{}, error) {
	allSQL := make(map[string][]interface{})
	for _, event := range tx.Events {
		sqls, err := m.ExtractEventStatements(ctx, block, tx.TxHash, sdk.Event(event))
		if err != nil {
			log.Err(err)
			continue
		}
		for k, v := range sqls {
			allSQL[k] = v
		}
	}
	return allSQL, nil
}

func (m *Module) ExtractEventStatements(ctx context.Context, block *tmctypes.ResultBlock, txHash string, event sdk.Event) (map[string][]interface{}, error) {
	if !StorageProviderEvents[event.Type] {
		return nil, nil
	}

	if !StorageProviderEvents[event.Type] {
		return nil, nil
	}

	typedEvent, err := sdk.ParseTypedEvent(abci.Event(event))
	if err != nil {
		return nil, err
	}

	switch event.Type {
	case EventCreateStorageProvider:
		createStorageProvider, ok := typedEvent.(*sptypes.EventCreateStorageProvider)
		if !ok {
			return nil, errors.New("create storage provider event assert error")
		}
		return m.handleCreateStorageProvider(ctx, block, txHash, createStorageProvider), nil
	case EventEditStorageProvider:
		editStorageProvider, ok := typedEvent.(*sptypes.EventEditStorageProvider)
		if !ok {
			return nil, errors.New("edit storage provider event assert error")
		}
		return m.handleEditStorageProvider(ctx, block, txHash, editStorageProvider), nil
	case EventSpStoragePriceUpdate:
		spStoragePriceUpdate, ok := typedEvent.(*sptypes.EventSpStoragePriceUpdate)
		if !ok {
			return nil, errors.New("storage provider price update event assert error")
		}
		return m.handleSpStoragePriceUpdate(ctx, block, txHash, spStoragePriceUpdate), nil
	case EventCompleteSpExit:
		completeSpExit, ok := typedEvent.(*vgtypes.EventCompleteStorageProviderExit)
		if !ok {
			return nil, errors.New("storage provider exit event assert error")
		}
		return m.handleCompleteStorageProviderExit(ctx, block, txHash, completeSpExit), nil
	}

	return nil, nil
}

func (m *Module) handleCreateStorageProvider(ctx context.Context, block *tmctypes.ResultBlock, txHash string, createStorageProvider *sptypes.EventCreateStorageProvider) map[string][]interface{} {
	storageProvider := &models.StorageProvider{
		SpID:            createStorageProvider.SpId,
		OperatorAddress: createStorageProvider.SpAddress,
		FundingAddress:  createStorageProvider.FundingAddress,
		SealAddress:     createStorageProvider.SealAddress,
		ApprovalAddress: createStorageProvider.ApprovalAddress,
		GcAddress:       createStorageProvider.GcAddress,
		TotalDeposit:    *createStorageProvider.TotalDeposit.Amount.BigInt(),
		Status:          createStorageProvider.Status.String(),
		Endpoint:        createStorageProvider.Endpoint,
		Moniker:         createStorageProvider.Description.Moniker,
		Identity:        createStorageProvider.Description.Identity,
		Website:         createStorageProvider.Description.Website,
		SecurityContact: createStorageProvider.Description.SecurityContact,
		Details:         createStorageProvider.Description.Details,
		BlsKey:          createStorageProvider.BlsKey,
		CreateTxHash:    txHash,
		CreateAt:        block.Block.Height,
		UpdateAt:        block.Block.Height,
		UpdateTxHash:    txHash,
		Removed:         false,
	}

	k, v := m.db.CreateStorageProviderToSQL(ctx, storageProvider)
	return map[string][]interface{}{
		k: v,
	}
}

func (m *Module) handleEditStorageProvider(ctx context.Context, block *tmctypes.ResultBlock, txHash string, editStorageProvider *sptypes.EventEditStorageProvider) map[string][]interface{} {
	storageProvider := &models.StorageProvider{
		SpID:            editStorageProvider.SpId,
		OperatorAddress: editStorageProvider.SpAddress,
		SealAddress:     editStorageProvider.SealAddress,
		ApprovalAddress: editStorageProvider.ApprovalAddress,
		GcAddress:       editStorageProvider.GcAddress,
		Endpoint:        editStorageProvider.Endpoint,
		Moniker:         editStorageProvider.Description.Moniker,
		Identity:        editStorageProvider.Description.Identity,
		Website:         editStorageProvider.Description.Website,
		SecurityContact: editStorageProvider.Description.SecurityContact,
		Details:         editStorageProvider.Description.Details,
		BlsKey:          editStorageProvider.BlsKey,

		UpdateAt:     block.Block.Height,
		UpdateTxHash: txHash,
		Removed:      false,
	}

	k, v := m.db.UpdateStorageProviderToSQL(ctx, storageProvider)
	return map[string][]interface{}{
		k: v,
	}
}

func (m *Module) handleSpStoragePriceUpdate(ctx context.Context, block *tmctypes.ResultBlock, txHash string, spStoragePriceUpdate *sptypes.EventSpStoragePriceUpdate) map[string][]interface{} {
	storageProvider := &models.StorageProvider{
		SpID:          spStoragePriceUpdate.SpId,
		UpdateTimeSec: spStoragePriceUpdate.UpdateTimeSec,
		ReadPrice:     *spStoragePriceUpdate.ReadPrice.BigInt(),
		FreeReadQuota: spStoragePriceUpdate.FreeReadQuota,
		StorePrice:    *spStoragePriceUpdate.StorePrice.BigInt(),

		UpdateAt:     block.Block.Height,
		UpdateTxHash: txHash,
		Removed:      false,
	}

	k, v := m.db.UpdateStorageProviderToSQL(ctx, storageProvider)
	return map[string][]interface{}{
		k: v,
	}
}

func (m *Module) handleCompleteStorageProviderExit(ctx context.Context, block *tmctypes.ResultBlock, txHash string, completeStorageProviderExit *vgtypes.EventCompleteStorageProviderExit) map[string][]interface{} {
	data := &models.StorageProvider{
		SpID: completeStorageProviderExit.StorageProviderId,

		UpdateAt:     block.Block.Height,
		UpdateTxHash: txHash,
		Removed:      true,
	}
	k, v := m.db.UpdateStorageProviderToSQL(ctx, data)
	return map[string][]interface{}{
		k: v,
	}
}
