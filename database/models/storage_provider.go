package models

import (
	"math/big"
)

type StorageProvider struct {
	ID              uint64  `gorm:"column:id;primaryKey"`
	SpID            uint32  `gorm:"column:sp_id;index:idx_sp_id"`
	OperatorAddress string  `gorm:"column:operator_address;type:TEXT;index:idx_operator_address"`
	FundingAddress  string  `gorm:"column:funding_address;type:TEXT"`
	SealAddress     string  `gorm:"column:seal_address;type:TEXT"`
	ApprovalAddress string  `gorm:"column:approval_address;type:TEXT"`
	GcAddress       string  `gorm:"column:gc_address;type:TEXT"`
	TotalDeposit    big.Int `gorm:"column:total_deposit"`
	Status          string  `gorm:"column:status;type:VARCHAR(50)"`
	Endpoint        string  `gorm:"column:endpoint;type:VARCHAR(256)"`
	Moniker         string  `gorm:"column:moniker;type:VARCHAR(128)"`
	Identity        string  `gorm:"column:identity;type:VARCHAR(256)"`
	Website         string  `gorm:"column:website;type:VARCHAR(128)"`
	SecurityContact string  `gorm:"column:security_contact;type:VARCHAR(128)"`
	Details         string  `gorm:"column:details;type:VARCHAR(256)"`
	BlsKey          string  `gorm:"column:bls_key;type:VARCHAR(96)"`
	UpdateTimeSec   int64   `gorm:"column:update_time_sec"`
	ReadPrice       big.Int `gorm:"column:read_price"`
	FreeReadQuota   uint64  `gorm:"column:free_read_quota"`
	StorePrice      big.Int `gorm:"column:store_price"`
	CreateAt        int64   `gorm:"column:create_at"`
	CreateTxHash    string  `gorm:"column:create_tx_hash;type:TEXT;not null"`
	UpdateAt        int64   `gorm:"column:update_at"`
	UpdateTxHash    string  `gorm:"column:update_tx_hash;type:TEXT;not null"`
	Removed         bool    `gorm:"column:removed;default:false"`
}

func (*StorageProvider) TableName() string {
	return "storage_providers"
}
