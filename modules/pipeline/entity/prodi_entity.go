package entity

import (
	"time"

	"gorm.io/gorm"
)

const (
	OldProdiTableName = "ref_prodi"
	NewProdiTableName = "ref_prodi"
)

type OldProdi struct {
	Kode          string         `json:"kode"`
	KodeDikti     string         `json:"kode_dikti"`
	KodeFak       string         `json:"kode_fak"`
	KodeIntegrasi string         `json:"kode_integrasi"`
	Nama          string         `json:"nama"`
	Jenjang       string         `json:"jenjang"`
	NamaFak       string         `json:"nama_fak"`
	CreatedAt     time.Time      `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"type:timestamptz;not_null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (OldProdi) TableName() string {
	return OldProdiTableName
}

type NewProdi struct {
	Id              uint32         `json:"id"`
	Kode            string         `json:"kode"`
	KodeDikti       string         `json:"kode_dikti"`
	KodeIntegrasi   string         `json:"kode_integrasi"`
	Nama            string         `json:"nama"`
	Jenjang         string         `json:"jenjang"`
	KodeFakultas    string         `json:"kode_fakultas"`
	NamaFakultas    string         `json:"nama_fakultas"`
	AkronimFakultas string         `json:"akronim_fakultas"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

func (NewProdi) TableName() string {
	return NewProdiTableName
}
