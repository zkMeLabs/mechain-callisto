package models

type Event struct {
	ID           uint64 `gorm:"id;type:bigint(64);primaryKey"`
	ResourceType string `gorm:"resource_type;type:varchar(64);not null"`
	ResourceID   string `gorm:"resource_id;type:TEXT;not null"`
	BlockID      int64  `gorm:"column:block_id;not null"`
	TxHash       string `gorm:"column:tx_hash;type:TEXT;not null"`
	Action       string `gorm:"column:action;type:TEXT;not null"`
}

func (*Event) TableName() string {
	return "events"
}
