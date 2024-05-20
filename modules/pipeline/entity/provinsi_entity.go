package entity

import (
	"time"

	"gorm.io/gorm"
)

const (
	OldProvinsiTableName = "ref_provinsi"
	NewProvinsiTableName = "ref_provinsi"
)

type OldProvinsi struct {
	IdWil     string         `json:"id_wil"`
	NmWil     string         `json:"nm_wil"`
	UMP12     string         `json:"12ump"`
	CreatedAt time.Time      `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamptz;not_null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (OldProvinsi) TableName() string {
	return OldProvinsiTableName
}

type NewProvinsi struct {
	Id        uint32         `json:"id"`
	IdWil     string         `json:"id_wil"`
	Nama      string         `json:"nama"`
	Ump       uint64         `json:"ump"`
	UmpPkts   uint64         `json:"ump_pkts"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (NewProvinsi) TableName() string {
	return NewProvinsiTableName
}
