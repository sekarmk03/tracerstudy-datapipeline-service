package entity

import (
	"time"

	"gorm.io/gorm"
)

const (
	OldKabKotaTableName = "ref_kabkota"
	NewKabKotaTableName = "ref_kabkota"
)

type OldKabkota struct {
	IdWil          string         `json:"id_wil"`
	NmWil          string         `json:"nm_wil"`
	IdIndukWilayah string         `json:"id_induk_wilayah"`
	CreatedAt      time.Time      `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"type:timestamptz;not_null" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (OldKabkota) TableName() string {
	return OldKabKotaTableName
}

type NewKabkota struct {
	Id         uint32         `json:"id"`
	IdWil      string         `json:"id_wil"`
	Nama       string         `json:"nama"`
	IdIndukWil string         `json:"id_induk_wil"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func (NewKabkota) TableName() string {
	return NewKabKotaTableName
}
