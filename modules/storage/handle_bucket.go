package storage

import (
	"strconv"
	"time"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/forbole/bdjuno/v4/database/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (m *Module) handleCreateBucket(block *tmctypes.ResultBlock, txHash string, values []abcitypes.EventAttribute) {
	b := &models.Bucket{
		Removed:      false,
		CreateAt:     block.Block.Height,
		CreateTxHash: txHash,
		UpdateAt:     block.Block.Height,
	}
	for _, v := range values {
		strV := ""
		if len(v.Value) > 2 {
			strV = v.Value[1 : len(v.Value)-1]
		}
		switch v.Key {
		case "bucket_id":
			b.BucketID = strV
		case "bucket_name":
			b.BucketName = strV
		case "owner":
			b.OwnerAddress = strV
		case "payment_address":
			b.PaymentAddress = strV
		case "global_virtual_group_family_id":
			globalVirtualGroupFamilyId, _ := strconv.ParseUint(strV, 10, 32)
			b.GlobalVirtualGroupFamilyId = uint32(globalVirtualGroupFamilyId)
		case "source_type":
			b.SourceType = strV
		case "charged_read_quota":
			chargedReadQuota, _ := strconv.ParseUint(strV, 10, 64)
			b.ChargedReadQuota = chargedReadQuota
		case "visibility":
			b.Visibility = strV
		case "status":
			b.Status = strV
		case "create_at":
			createTime, _ := strconv.ParseInt(strV, 10, 64)
			b.CreateTime = time.Unix(createTime, 0)

		}
	}
	b.OperatorAddress = b.OwnerAddress
	b.UpdateTxHash = b.CreateTxHash
	b.UpdateTime = b.CreateTime

	dsn := "postgresql://postgres:postgres_mechain@localhost:5432/bdjuno?sslmode=disable&search_path=public"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.Bucket{})
	if err != nil {
		log.Fatal().Str("module", "storage").Err(err)
	}
	result := m.db.G.Create(b)
	if result.Error != nil {
		log.Fatal().Str("module", "storage").Err(result.Error)
	}
}

func (m *Module) handleDeleteBucket(block *tmctypes.ResultBlock, txHash string, values []abcitypes.EventAttribute) {
	b := &models.Bucket{
		Removed:      true,
		UpdateAt:     block.Block.Height,
		UpdateTxHash: txHash,
		UpdateTime:   time.Unix(block.Block.Time.Unix(), 0),
	}
	for _, v := range values {
		strV := ""
		if len(v.Value) > 2 {
			strV = v.Value[1 : len(v.Value)-1]
		}
		switch v.Key {
		case "bucket_id":
			b.BucketID = strV
		case "bucket_name":
			b.BucketName = strV
		case "owner":
			b.OwnerAddress = strV
		case "operator":
			b.OperatorAddress = strV
		case "global_virtual_group_family_id":
			globalVirtualGroupFamilyId, _ := strconv.ParseUint(strV, 10, 32)
			b.GlobalVirtualGroupFamilyId = uint32(globalVirtualGroupFamilyId)
		}
	}
	result := m.db.G.Where("bucket_id = ?", b.BucketID).Updates(b)
	if result.Error != nil {
		log.Fatal().Str("module", "storage").Err(result.Error)
	}
}
