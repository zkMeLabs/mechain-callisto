package storage

import (
	"strconv"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/forbole/bdjuno/v4/types"
	"github.com/rs/zerolog/log"
)

func (m *Module) handleCreateGroup(height int64, values []abcitypes.EventAttribute) {
	msg := &createGroupEvent{}
	for _, v := range values {
		strV := ""
		if len(v.Value) > 2 {
			strV = v.Value[1 : len(v.Value)-1]
		}
		switch v.Key {
		case "group_id":
			msg.GroupId = strV
		case "group_name":
			msg.GroupName = strV
		case "owner":
			msg.Owner = strV
		case "source_type":
			msg.SourceType = strV
		case "extra":
			msg.Extra = strV
		}
	}

	groupId, _ := strconv.ParseUint(msg.GroupId, 10, 64)

	err := m.db.SaveStorageGroup(height, []types.StorageGroup{
		{
			GroupId:    groupId,
			GroupName:  msg.GroupName,
			Owner:      msg.Owner,
			SourceType: msg.SourceType,
			Extra:      msg.Extra,
		},
	})
	if err != nil {
		log.Error().Str("module", "storage").Err(err)
	}
}

func (m *Module) handleDeleteGroup(height int64, values []abcitypes.EventAttribute) {
	msg := &deleteGroupEvent{}
	for _, v := range values {
		strV := ""
		if len(v.Value) > 2 {
			strV = v.Value[1 : len(v.Value)-1]
		}
		switch v.Key {
		case "group_id":
			msg.GroupId = strV
		case "group_name":
			msg.GroupName = strV
		case "owner":
			msg.Owner = strV
		}
	}
}
