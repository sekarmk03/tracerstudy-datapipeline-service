package entity

import (
	"time"

	"gorm.io/gorm"
)

const (
	OldUserStudyTableName = "user_study"
	NewUserStudyTableName = "user_study"
)

type OldUserStudy struct {
	ID                                uint32         `json:"id"`
	NamaResponden                     string         `json:"nama_responden"`
	EmailResponden                    string         `json:"email_responden"`
	HpResponden                       string         `json:"hp_responden"`
	NamaInstansi                      string         `json:"nama_instansi"`
	Jabatan                           string         `json:"jabatan"`
	AlamatInstansi                    string         `json:"alamat_instansi"`
	NimLulusan                        string         `json:"nim_lulusan"`
	NamaLulusan                       string         `json:"nama_lulusan"`
	ProdiLulusan                      string         `json:"prodi_lulusan"`
	TahunLulusan                      string         `json:"tahun_lulusan"`
	LamaMengenalLulusan               string         `json:"lama_mengenal_lulusan"`
	Etika                             string         `json:"etika"`
	KeahlianBidIlmu                   string         `json:"keahlian_bid_ilmu"`
	BahasaInggris                     string         `json:"bahasa_inggris"`
	PenggunaanTi                      string         `json:"penggunaan_ti"`
	Komunikasi                        string         `json:"komunikasi"`
	KerjasamaTim                      string         `json:"kerjasama_tim"`
	PengembanganDiri                  string         `json:"pengembangan_diri"`
	KesiapanTerjunMasy                string         `json:"kesiapan_terjun_masy"`
	KeunggulanLulusan                 string         `json:"keunggulan_lulusan"`
	KelemahanLulusan                  string         `json:"kelemahan_lulusan"`
	SaranPeningkatanKompetensiLulusan string         `json:"saran_peningkatan_kompetensi_lulusan"`
	SaranPerbaikanKurikulum           string         `json:"saran_peningkatan_kurikulum"`
	CreatedBy                         string         `json:"created_by"`
	UpdatedBy                         string         `json:"updated_by"`
	CreatedAt                         time.Time      `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt                         time.Time      `gorm:"type:timestamptz;not_null" json:"updated_at"`
	DeletedAt                         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (OldUserStudy) TableName() string {
	return OldUserStudyTableName
}

type NewUserStudy struct {
	Id                                uint64         `json:"id"`
	NamaResponden                     string         `json:"nama_responden"`
	EmailResponden                    string         `json:"email_responden"`
	HpResponden                       string         `json:"hp_responden"`
	NamaInstansi                      string         `json:"nama_instansi"`
	Jabatan                           string         `json:"jabatan"`
	AlamatInstansi                    string         `json:"alamat_instansi"`
	NimLulusan                        string         `json:"nim_lulusan"`
	NamaLulusan                       string         `json:"nama_lulusan"`
	ProdiLulusan                      string         `json:"prodi_lulusan"`
	TahunLulusan                      string         `json:"tahun_lulusan"`
	LamaMengenalLulusan               uint32         `json:"lama_mengenal_lulusan"`
	Etika                             string         `json:"etika"`
	KeahlianBidIlmu                   string         `json:"keahlian_bid_ilmu"`
	BahasaInggris                     string         `json:"bahasa_inggris"`
	PenggunaanTi                      string         `json:"penggunaan_ti"`
	Komunikasi                        string         `json:"komunikasi"`
	KerjasamaTim                      string         `json:"kerjasama_tim"`
	PengembanganDiri                  string         `json:"pengembangan_diri"`
	KesiapanTerjunMasy                string         `json:"kesiapan_terjun_masy"`
	KeunggulanLulusan                 string         `json:"keunggulan_lulusan"`
	KelemahanLulusan                  string         `json:"kelemahan_lulusan"`
	SaranPeningkatanKompetensiLulusan string         `json:"saran_peningkatan_kompetensi_lulusan"`
	SaranPerbaikanKurikulum           string         `json:"saran_peningkatan_kurikulum"`
	CreatedAt                         time.Time      `json:"created_at"`
	UpdatedAt                         time.Time      `json:"updated_at"`
	DeletedAt                         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (NewUserStudy) TableName() string {
	return NewUserStudyTableName
}
