package models

import (
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/lib/pq"
)

type Permission struct {
	ID             uint64      `gorm:"id;type:bigint(64);primaryKey"`
	PrincipalType  int32       `gorm:"principal_type;type:int;uniqueIndex:idx_policy,priority:1"`
	PrincipalValue string      `gorm:"principal_value;type:varchar(128);uniqueIndex:idx_policy,priority:2"`
	ResourceType   string      `gorm:"resource_type;type:varchar(64);uniqueIndex:idx_policy,priority:3"`
	ResourceID     common.Hash `gorm:"resource_id;type:BINARY(32);uniqueIndex:idx_policy,priority:4"`
	PolicyID       common.Hash `gorm:"policy_id;type:BINARY(32);index:idx_policy_id"`
	CreateTime     time.Time   `gorm:"create_timestamp;type:bigint(64)"`
	UpdateTime     time.Time   `gorm:"update_timestamp;type:bigint(64)"`
	ExpirationTime time.Time   `gorm:"expiration_time;type:bigint(64)"` // seconds
	Removed        bool        `gorm:"removed;"`
}

func (p Permission) TableName() string {
	return "permission"
}

type Statements struct {
	ID             uint64         `gorm:"id;type:bigint(64);primaryKey"`
	PolicyID       common.Hash    `gorm:"policy_id;type:BINARY(32);index:idx_policy_id"`
	Effect         string         `gorm:"effect;type:varchar(32)"`
	ActionValue    int            `gorm:"action_value;type:int"`
	Resources      pq.StringArray `gorm:"resources;type:text"`
	ExpirationTime time.Time      `gorm:"expiration_time;type:bigint(64)"`
	LimitSize      uint64         `gorm:"limit_size;type:bigint(64)"`
	Removed        bool           `gorm:"removed;"`
}

func (s Statements) TableName() string {
	return "statements"
}

type GroupMember struct {
	ID      uint64         `gorm:"id;type:bigint(64);primaryKey"`
	GroupID uint64         `gorm:"id;type:bigint(64);index:idx_group_id"`
	Account common.Address `gorm:"account;type:BINARY(20);index:idx_account"`
	Removed bool           `gorm:"removed"`
}

func (g GroupMember) TableName() string {
	return "group_member"
}