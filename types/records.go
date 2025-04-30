package types

import (
	"time"

	"github.com/google/uuid"
)

type RecordPermission int8

const (
	RecordPermissionPublic  RecordPermission = 1
	RecordPermissionPrivate RecordPermission = 2
)

type RecordType int8

const (
	RecordTypeString RecordType = 1
	RecordTypeImage  RecordType = 2
	RecordTypeVideo  RecordType = 3
	RecordTypeAudio  RecordType = 4
	RecordTypeFile   RecordType = 5
	RecordTypeEvent  RecordType = 6
)

type Record struct {
	Id        uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	Network   string     `json:"network" gorm:"size:32;not null"`
	UserId    uuid.UUID  `json:"user_id" gorm:"type:uuid"`
	Type      RecordType `json:"type" gorm:"type:numeric"`
	Name      string     `json:"name" gorm:"size:64"`
	Data      string     `json:"data" gorm:"type:text"`
	ExpiresAt *time.Time `json:"expires_at,omitempty" gorm:"type:TIMESTAMP WITH TIME ZONE"`
	CreateAt  time.Time  `json:"created_at" gorm:"type:TIMESTAMP WITH TIME ZONE"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"type:TIMESTAMP WITH TIME ZONE"`
}
