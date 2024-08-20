package models

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type GlobalVirtualGroup struct {
	ID                    uint64         `gorm:"column:id;primaryKey"`
	GlobalVirtualGroupId  uint32         `gorm:"column:global_virtual_group_id;index:idx_gvg_id"`
	FamilyId              uint32         `gorm:"column:family_id"`
	PrimarySpId           uint32         `gorm:"column:primary_sp_id;index:idx_primary_sp_id"`
	SecondarySpIds        []uint32       `gorm:"column:secondary_sp_ids;type:TEXT"`
	StoredSize            uint64         `gorm:"column:stored_size"`
	VirtualPaymentAddress common.Address `gorm:"column:virtual_payment_address;type:BINARY(20)"`
	TotalDeposit          big.Int        `gorm:"column:total_deposit"`
	CreateAt              int64          `gorm:"column:create_at"`
	CreateTxHash          common.Hash    `gorm:"column:create_tx_hash;type:BINARY(32);not null"`
	CreateTime            time.Time      `gorm:"column:create_time"` // seconds
	UpdateAt              int64          `gorm:"column:update_at"`
	UpdateTxHash          common.Hash    `gorm:"column:update_tx_hash;type:BINARY(32);not null"`
	UpdateTime            time.Time      `gorm:"column:update_time"` // seconds
	Removed               bool           `gorm:"column:removed;default:false"`
}

func (*GlobalVirtualGroup) TableName() string {
	return "global_virtual_groups"
}
