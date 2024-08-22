package storage

import (
	"context"
	"errors"
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"
	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/forbole/bdjuno/v4/database/models"
	storagetypes "github.com/forbole/bdjuno/v4/modules/storage/types"
)

var (
	EventCreateGroup       = proto.MessageName(&storagetypes.EventCreateGroup{})
	EventDeleteGroup       = proto.MessageName(&storagetypes.EventDeleteGroup{})
	EventLeaveGroup        = proto.MessageName(&storagetypes.EventLeaveGroup{})
	EventUpdateGroupMember = proto.MessageName(&storagetypes.EventUpdateGroupMember{})
	EventRenewGroupMember  = proto.MessageName(&storagetypes.EventRenewGroupMember{})
)

var GroupEvents = map[string]bool{
	EventCreateGroup:       true,
	EventDeleteGroup:       true,
	EventLeaveGroup:        true,
	EventUpdateGroupMember: true,
	EventRenewGroupMember:  true,
}

func (m *Module) ExtractGroupEventStatements(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, event sdk.Event) (map[string][]interface{}, error) {
	typedEvent, err := sdk.ParseTypedEvent(abci.Event(event))
	if err != nil {
		return nil, err
	}

	switch event.Type {
	case EventCreateGroup:
		createGroup, ok := typedEvent.(*storagetypes.EventCreateGroup)
		if !ok {
			return nil, errors.New("create group event assert error")
		}
		return m.handleCreateGroup(ctx, block, txHash, evmTxHash, createGroup), nil
	case EventUpdateGroupMember:
		updateGroupMember, ok := typedEvent.(*storagetypes.EventUpdateGroupMember)
		if !ok {
			return nil, errors.New("update group member event assert error")
		}
		return m.handleUpdateGroupMember(ctx, block, txHash, evmTxHash, updateGroupMember), nil
	case EventDeleteGroup:
		deleteGroup, ok := typedEvent.(*storagetypes.EventDeleteGroup)
		if !ok {
			return nil, errors.New("delete group event assert error")
		}
		return m.handleDeleteGroup(ctx, block, txHash, evmTxHash, deleteGroup), nil
	case EventLeaveGroup:
		leaveGroup, ok := typedEvent.(*storagetypes.EventLeaveGroup)
		if !ok {
			return nil, errors.New("leave group event assert error")
		}
		return m.handleLeaveGroup(ctx, block, txHash, evmTxHash, leaveGroup), nil
	case EventRenewGroupMember:
		renewGroupMember, ok := typedEvent.(*storagetypes.EventRenewGroupMember)
		if !ok {
			return nil, errors.New("renew group member event assert error")
		}
		return m.handleRenewGroupMember(ctx, block, txHash, evmTxHash, renewGroupMember), nil
	}
	return nil, nil
}

func (m *Module) handleCreateGroup(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, createGroup *storagetypes.EventCreateGroup) map[string][]interface{} {
	var membersToAddList []*models.Group
	groupItem := &models.Group{
		OwnerAddress:    createGroup.Owner,
		GroupID:         createGroup.GroupId.BigInt().String(),
		GroupName:       createGroup.GroupName,
		SourceType:      createGroup.SourceType.String(),
		Extra:           createGroup.Extra,
		CreateAt:        block.Block.Height,
		CreateTxHash:    txHash,
		CreateEVMTxHash: evmTxHash,
		CreateTime:      block.Block.Time,
		UpdateAt:        block.Block.Height,
		UpdateTime:      block.Block.Time,
		UpdateTxHash:    txHash,
		UpdateEVMTxHash: evmTxHash,
		Removed:         false,
	}
	membersToAddList = append(membersToAddList, groupItem)
	k, v := m.db.CreateGroupToSQL(ctx, membersToAddList)
	return map[string][]interface{}{
		k: v,
	}
}

func (m *Module) handleDeleteGroup(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, deleteGroup *storagetypes.EventDeleteGroup) map[string][]interface{} {
	group := &models.Group{
		OwnerAddress:    deleteGroup.Owner,
		GroupID:         deleteGroup.GroupId.BigInt().String(),
		GroupName:       deleteGroup.GroupName,
		UpdateAt:        block.Block.Height,
		UpdateTime:      block.Block.Time,
		UpdateTxHash:    txHash,
		UpdateEVMTxHash: evmTxHash,
		Removed:         true,
	}
	res := make(map[string][]interface{})
	k, v := m.db.DeleteGroupToSQL(ctx, group)
	res[k] = v
	return res
}

func (m *Module) handleLeaveGroup(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, leaveGroup *storagetypes.EventLeaveGroup) map[string][]interface{} {
	group := &models.Group{
		OwnerAddress:    leaveGroup.Owner,
		GroupID:         leaveGroup.GroupId.BigInt().String(),
		GroupName:       leaveGroup.GroupName,
		AccountAddress:  leaveGroup.MemberAddress,
		UpdateAt:        block.Block.Height,
		UpdateTime:      block.Block.Time,
		UpdateTxHash:    txHash,
		UpdateEVMTxHash: evmTxHash,
		Removed:         true,
	}

	// update group item
	groupItem := &models.Group{
		GroupID:    leaveGroup.GroupId.BigInt().String(),
		UpdateAt:   block.Block.Height,
		UpdateTime: block.Block.Time,
		Removed:    false,
	}
	res := make(map[string][]interface{})
	k, v := m.db.UpdateGroupToSQL(ctx, groupItem)
	res[k] = v

	k, v = m.db.UpdateGroupToSQL(ctx, group)
	res[k] = v
	return res
}

func (m *Module) handleUpdateGroupMember(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, updateGroupMember *storagetypes.EventUpdateGroupMember) map[string][]interface{} {
	membersToAdd := updateGroupMember.MembersToAdd
	membersToDelete := updateGroupMember.MembersToDelete

	var membersToAddList []*models.Group
	res := make(map[string][]interface{})

	if len(membersToAdd) > 0 {
		for _, memberToAdd := range membersToAdd {
			groupItem := &models.Group{
				OwnerAddress:    updateGroupMember.Owner,
				GroupID:         updateGroupMember.GroupId.BigInt().String(),
				GroupName:       updateGroupMember.GroupName,
				AccountAddress:  memberToAdd.Member,
				OperatorAddress: updateGroupMember.Operator,
				CreateAt:        block.Block.Height,
				CreateTime:      block.Block.Time,
				CreateTxHash:    txHash,
				CreateEVMTxHash: evmTxHash,
				UpdateAt:        block.Block.Height,
				UpdateTime:      block.Block.Time,
				UpdateTxHash:    txHash,
				UpdateEVMTxHash: evmTxHash,
				Removed:         false,
			}
			if memberToAdd.ExpirationTime != nil {
				groupItem.ExpirationTime = *memberToAdd.ExpirationTime
			}
			membersToAddList = append(membersToAddList, groupItem)
		}
		k, v := m.db.CreateGroupToSQL(ctx, membersToAddList)
		res[k] = v
	}

	if len(membersToDelete) > 0 {
		groupItem := &models.Group{
			OperatorAddress: updateGroupMember.Operator,
			UpdateAt:        block.Block.Height,
			UpdateTime:      block.Block.Time,
			UpdateTxHash:    txHash,
			UpdateEVMTxHash: evmTxHash,
			Removed:         true,
		}
		accountAddresses := make([]string, 0, len(membersToDelete))
		for _, memberToDelete := range membersToDelete {
			accountAddresses = append(accountAddresses, memberToDelete)
		}
		k, v := m.db.BatchDeleteGroupMemberToSQL(ctx, groupItem, updateGroupMember.GroupId.BigInt().String(), accountAddresses)
		res[k] = v
	}

	// update group item
	groupItem := &models.Group{
		GroupID:         updateGroupMember.GroupId.BigInt().String(),
		UpdateAt:        block.Block.Height,
		UpdateTime:      block.Block.Time,
		UpdateTxHash:    txHash,
		UpdateEVMTxHash: evmTxHash,
		Removed:         false,
	}
	k, v := m.db.UpdateGroupToSQL(ctx, groupItem)
	res[k] = v

	return res
}

func (m *Module) handleRenewGroupMember(ctx context.Context, block *tmctypes.ResultBlock, txHash, evmTxHash string, renewGroupMember *storagetypes.EventRenewGroupMember) map[string][]interface{} {
	res := map[string][]interface{}{}
	for _, e := range renewGroupMember.Members {
		expirationTime := int64(0)
		if e.ExpirationTime != nil {
			expirationTime = e.ExpirationTime.Unix()
		}
		k := fmt.Sprintf("Update `groups` set expiration_time = ?, update_at = ?,update_time = ? where account_id = %s and group_id = ?", e.Member)
		v := []interface{}{expirationTime, block.Block.Height, block.Block.Time, renewGroupMember.GroupId.BigInt().String()}
		res[k] = v
	}

	return res
}
