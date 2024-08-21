package storage

import (
	"strconv"
	"time"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/forbole/bdjuno/v4/database/models"
	junotypes "github.com/forbole/juno/v5/types"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (m *Module) handleCreateBucket(tx *junotypes.Tx, values []abcitypes.EventAttribute) {
	b := &models.Bucket{
		Removed:      false,
		CreateAt:     tx.Height,
		CreateTxHash: common.HexToHash(tx.TxHash),
		UpdateAt:     tx.Height,
	}
	for _, v := range values {
		strV := ""
		if len(v.Value) > 2 {
			strV = v.Value[1 : len(v.Value)-1]
		}
		switch v.Key {
		case "bucket_id":
			b.BucketID = common.HexToHash(strV)
		case "bucket_name":
			b.BucketName = strV
		case "owner":
			b.OwnerAddress = common.HexToAddress(strV)
		case "payment_address":
			b.PaymentAddress = common.HexToAddress(strV)
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

	// 迁移 schema (自动创建表)
	dsn := "postgresql://postgres:postgres_mechain@localhost:5432/bdjuno?sslmode=disable&search_path=public"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.Bucket{})
	if err != nil {
		log.Fatal().Str("module", "storage").Err(err)
	}
	result := db.Create(b)
	if result.Error != nil {
		log.Fatal().Str("module", "storage").Err(result.Error)
	}
}

func (m *Module) handleDeleteBucket(height int64, values []abcitypes.EventAttribute) {
	msg := &deleteBucketEvent{}
	for _, v := range values {
		strV := ""
		if len(v.Value) > 2 {
			strV = v.Value[1 : len(v.Value)-1]
		}
		switch v.Key {
		case "bucket_id":
			msg.BucketId = strV
		case "bucket_name":
			msg.BucketName = strV
		case "owner":
			msg.Owner = strV
		case "operator":
			msg.Operator = strV
		case "global_virtual_group_family_id":
			globalVirtualGroupFamilyId, _ := strconv.ParseUint(strV, 10, 32)
			msg.GlobalVirtualGroupFamilyId = uint32(globalVirtualGroupFamilyId)
		}
	}
}
