package entity

import (
	"time"

	"gorm.io/gorm"
)

const (
	OldPktsTableName = "pkts"
	NewPktsTableName = "pkts"
)

type OldPkts struct {
	ID                  uint32         `json:"id"`
	Nim                 string         `json:"nim"`
	Kodeprodi           string         `json:"kodeprodi"`
	ThnSidang           string         `json:"thn_sidang"`
	F8                  uint32         `json:"f8"`
	F5_04               uint32         `json:"f5_04"`
	F5_02               string         `json:"f5_02"`
	F5_06               string         `json:"f5_06"`
	F5_05               string         `json:"f5_05"`
	F5a1                string         `json:"f5a1"`
	F5a2                string         `json:"f5a2"`
	F11_01              uint32         `json:"f11_01"`
	F11_02              string         `json:"f11_02"`
	F5b                 string         `json:"f5b"`
	F5c                 uint32         `json:"f5c"`
	F5d                 uint32         `json:"f5d"`
	F18a                uint32         `json:"f18a"`
	F18b                string         `json:"f18b"`
	F18c                string         `json:"f18c"`
	F18d                string         `json:"f18d"`
	F12_01              uint32         `json:"f12_01"`
	F12_02              string         `json:"f12_02"`
	F14                 uint32         `json:"f14"`
	F15                 uint32         `json:"f15"`
	F1761               uint32         `json:"f1761"`
	F1762               uint32         `json:"f1762"`
	F1763               uint32         `json:"f1763"`
	F1764               uint32         `json:"f1764"`
	F1765               uint32         `json:"f1765"`
	F1766               uint32         `json:"f1766"`
	F1767               uint32         `json:"f1767"`
	F1768               uint32         `json:"f1768"`
	F1769               uint32         `json:"f1769"`
	F1770               uint32         `json:"f1770"`
	F1771               uint32         `json:"f1771"`
	F1772               uint32         `json:"f1772"`
	F1773               uint32         `json:"f1773"`
	F1774               uint32         `json:"f1774"`
	F21                 uint32         `json:"f21"`
	F22                 uint32         `json:"f22"`
	F23                 uint32         `json:"f23"`
	F24                 uint32         `json:"f24"`
	F25                 uint32         `json:"f25"`
	F26                 uint32         `json:"f26"`
	F27                 uint32         `json:"f27"`
	F301                uint32         `json:"f301"`
	F302                string         `json:"f302"`
	F303                string         `json:"f303"`
	F4_01               string         `json:"f4_01"`
	F4_02               string         `json:"f4_02"`
	F4_03               string         `json:"f4_03"`
	F4_04               string         `json:"f4_04"`
	F4_05               string         `json:"f4_05"`
	F4_06               string         `json:"f4_06"`
	F4_07               string         `json:"f4_07"`
	F4_08               string         `json:"f4_08"`
	F4_09               string         `json:"f4_09"`
	F4_10               string         `json:"f4_10"`
	F4_11               string         `json:"f4_11"`
	F4_12               string         `json:"f4_12"`
	F4_13               string         `json:"f4_13"`
	F4_14               string         `json:"f4_14"`
	F4_15               string         `json:"f4_15"`
	F4_16               string         `json:"f4_16"`
	F6                  string         `json:"f6"`
	F7                  string         `json:"f7"`
	F7a                 string         `json:"f7a"`
	F10_01              uint32         `json:"f10_01"`
	F10_02              string         `json:"f10_02"`
	F16_01              string         `json:"f16_01"`
	F16_02              string         `json:"f16_02"`
	F16_03              string         `json:"f16_03"`
	F16_04              string         `json:"f16_04"`
	F16_05              string         `json:"f16_05"`
	F16_06              string         `json:"f16_06"`
	F16_07              string         `json:"f16_07"`
	F16_08              string         `json:"f16_08"`
	F16_09              string         `json:"f16_09"`
	F16_10              string         `json:"f16_10"`
	F16_11              string         `json:"f16_11"`
	F16_12              string         `json:"f16_12"`
	F16_13              string         `json:"f16_13"`
	F16_14              string         `json:"f16_14"`
	NamaAtasan          string         `json:"nama_atasan"`
	HpAtasan            string         `json:"hp_atasan"`
	EmailAtasan         string         `json:"email_atasan"`
	TinggalSelamaKuliah string         `json:"tinggal_selama_kuliah"`
	Code                string         `json:"code"`
	MailSent            string         `json:"mail_sent"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (OldPkts) TableName() string {
	return OldPktsTableName
}

type NewPkts struct {
	Id                  uint64         `json:"id"`
	Nim                 string         `json:"nim"`
	KodeProdi           string         `json:"kode_prodi"`
	TahunSidang         string         `json:"tahun_sidang"`
	F8                  uint16         `json:"f8"`
	F504                uint16         `json:"f504"`
	F502                uint32         `json:"f502"`
	F506                uint32         `json:"f506"`
	F505                uint64         `json:"f505"`
	F5a1                string         `json:"f5a1"`
	F5a2                string         `json:"f5a2"`
	F1101               uint16         `json:"f1101"`
	F1102               string         `json:"f1102"`
	F5b                 string         `json:"f5b"`
	F5c                 uint16         `json:"f5c"`
	F5d                 uint16         `json:"f5d"`
	F18a                uint16         `json:"f18a"`
	F18b                string         `json:"f18b"`
	F18c                string         `json:"f18c"`
	F18d                string         `json:"f18d"`
	F1201               uint16         `json:"f1201"`
	F1202               string         `json:"f1202"`
	F14                 uint16         `json:"f14"`
	F15                 uint16         `json:"f15"`
	F1761               uint16         `json:"f1761"`
	F1762               uint16         `json:"f1762"`
	F1763               uint16         `json:"f1763"`
	F1764               uint16         `json:"f1764"`
	F1765               uint16         `json:"f1765"`
	F1766               uint16         `json:"f1766"`
	F1767               uint16         `json:"f1767"`
	F1768               uint16         `json:"f1768"`
	F1769               uint16         `json:"f1769"`
	F1770               uint16         `json:"f1770"`
	F1771               uint16         `json:"f1771"`
	F1772               uint16         `json:"f1772"`
	F1773               uint16         `json:"f1773"`
	F1774               uint16         `json:"f1774"`
	F21                 uint16         `json:"f21"`
	F22                 uint16         `json:"f22"`
	F23                 uint16         `json:"f23"`
	F24                 uint16         `json:"f24"`
	F25                 uint16         `json:"f25"`
	F26                 uint16         `json:"f26"`
	F27                 uint16         `json:"f27"`
	F301                uint16         `json:"f301"`
	F302                uint32         `json:"f302"`
	F303                uint32         `json:"f303"`
	F401                uint8          `json:"f401"`
	F402                uint8          `json:"f402"`
	F403                uint8          `json:"f403"`
	F404                uint8          `json:"f404"`
	F405                uint8          `json:"f405"`
	F406                uint8          `json:"f406"`
	F407                uint8          `json:"f407"`
	F408                uint8          `json:"f408"`
	F409                uint8          `json:"f409"`
	F410                uint8          `json:"f410"`
	F411                uint8          `json:"f411"`
	F412                uint8          `json:"f412"`
	F413                uint8          `json:"f413"`
	F414                uint8          `json:"f414"`
	F415                uint8          `json:"f415"`
	F416                string         `json:"f416"`
	F6                  uint32         `json:"f6"`
	F7                  uint32         `json:"f7"`
	F7a                 uint32         `json:"f7a"`
	F1001               uint16         `json:"f1001"`
	F1002               string         `json:"f1002"`
	F1601               uint16         `json:"f1601"`
	F1602               uint16         `json:"f1602"`
	F1603               uint16         `json:"f1603"`
	F1604               uint16         `json:"f1604"`
	F1605               uint16         `json:"f1605"`
	F1606               uint16         `json:"f1606"`
	F1607               uint16         `json:"f1607"`
	F1608               uint16         `json:"f1608"`
	F1609               uint16         `json:"f1609"`
	F1610               uint16         `json:"f1610"`
	F1611               uint16         `json:"f1611"`
	F1612               uint16         `json:"f1612"`
	F1613               uint16         `json:"f1613"`
	F1614               string         `json:"f1614"`
	NamaAtasan          string         `json:"nama_atasan"`
	HpAtasan            string         `json:"hp_atasan"`
	EmailAtasan         string         `json:"email_atasan"`
	TinggalSelamaKuliah string         `json:"tinggal_selama_kuliah"`
	Code                string         `json:"code"`
	MailSent            time.Time      `json:"mail_sent"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (NewPkts) TableName() string {
	return NewPktsTableName
}
