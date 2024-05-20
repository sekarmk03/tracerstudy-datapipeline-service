package entity

import (
	"time"

	"gorm.io/gorm"
)

const (
	OldRespondenTableName = "responden"
	NewRespondenTableName = "responden"
)

type OldResponden struct {
	Id           uint32         `json:"id"`
	ThnAk        string         `json:"thn_ak"`
	Semester     uint32         `json:"semester"`
	Type         string         `json:"type"`
	StatusUpdate string         `json:"status_update"`
	Jlrmasuk     string         `json:"jlrmasuk"`
	Thnmasuk     string         `json:"thnmasuk"`
	Lamastd      string         `json:"lamastd"`
	Kodefak      string         `json:"kodefak"`
	Namafak      string         `json:"namafak"`
	Nim          string         `json:"nim"`
	Nama         string         `json:"nama"`
	Kodeprodi    string         `json:"kodeprodi"`
	Kodeprodi2   string         `json:"kodeprodi2"`
	Namaprodi    string         `json:"namaprodi"`
	Namaprodi2   string         `json:"namaprodi2"`
	Kodedikti    string         `json:"kodedikti"`
	Jenjang      string         `json:"jenjang"`
	JK           string         `json:"jk"`
	Email        string         `json:"email"`
	Hp           string         `json:"hp"`
	Ipk          string         `json:"ipk"`
	TglSidang    string         `json:"tgl_sidang"`
	ThnSidang    string         `json:"thn_sidang"`
	TglWisuda    string         `json:"tgl_wisuda"`
	Nik          string         `json:"nik"`
	Npwp         string         `json:"npwp"`
	CreatedBy    string         `json:"created_by"`
	UpdatedBy    string         `json:"updated_by"`
	CreatedAt    time.Time      `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"type:timestamptz;not_null" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (OldResponden) TableName() string {
	return OldRespondenTableName
}

type NewResponden struct {
	Id            uint32         `json:"id"`
	Nim           string         `json:"nim"`
	Nama          string         `json:"nama"`
	StatusUpdate  uint32         `json:"status_update"`
	JalurMasuk    string         `json:"jalur_masuk"`
	TahunMasuk    string         `json:"tahun_masuk"`
	LamaStudi     uint32         `json:"lama_studi"`
	KodeFakultas  string         `json:"kode_fakultas"`
	KodeProdi     string         `json:"kode_prodi"`
	JenisKelamin  string         `json:"jenis_kelamin"`
	Email         string         `json:"email"`
	Hp            string         `json:"hp"`
	Ipk           string         `json:"ipk"`
	TanggalSidang string         `json:"tanggal_sidang"`
	TahunSidang   string         `json:"tahun_sidang"`
	TanggalWisuda string         `json:"tanggal_wisuda"`
	Nik           string         `json:"nik"`
	Npwp          string         `json:"npwp"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

func (NewResponden) TableName() string {
	return NewRespondenTableName
}
