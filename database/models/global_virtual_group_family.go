package models

import (
	"time"
)

type GlobalVirtualGroupFamily struct {
	ID                         uint64    `gorm:"column:id;primaryKey"`
	GlobalVirtualGroupFamilyId uint32    `gorm:"column:global_virtual_group_family_id;index:idx_vgf_id"`
	PrimarySpID                uint32    `gorm:"column:primary_sp_id;index:idx_primary_sp_id"`
	GlobalVirtualGroupIDs      []uint32  `gorm:"column:global_virtual_group_ids;type:TEXT"`
	VirtualPaymentAddress      string    `gorm:"column:virtual_payment_address;type:TEXT"`
	CreateAt                   int64     `gorm:"column:create_at"`
	CreateTxHash               string    `gorm:"column:create_tx_hash;type:TEXT;not null"`
	CreateTime                 time.Time `gorm:"column:create_time"`
	UpdateAt                   int64     `gorm:"column:update_at"`
	UpdateTxHash               string    `gorm:"column:update_tx_hash;type:TEXT;not null"`
	UpdateTime                 time.Time `gorm:"column:update_time"`
}

func (*GlobalVirtualGroupFamily) TableName() string {
	return "global_virtual_group_families"
}
