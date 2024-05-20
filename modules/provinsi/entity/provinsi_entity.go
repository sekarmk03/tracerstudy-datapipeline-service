package entity

import (
	"time"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const (
	ProvinsiTableName = "ref_provinsi"
)

type Provinsi struct {
	Id        uint32         `json:"id"`
	IdWil     string         `json:"id_wil"`
	Nama      string         `json:"nama"`
	Ump       uint64         `json:"ump"`
	UmpPkts   uint64         `json:"ump_pkts"`
	CreatedAt time.Time      `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamptz;not_null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func NewProvinsi(id uint32, idWil, nama string, ump, umpPkts uint64) *Provinsi {
	return &Provinsi{
		Id:        id,
		IdWil:     idWil,
		Nama:      nama,
		Ump:       ump,
		UmpPkts:   umpPkts,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (p *Provinsi) TableName() string {
	return ProvinsiTableName
}

func ConvertEntityToProto(p *Provinsi) *pb.Provinsi {
	return &pb.Provinsi{
		Id:        p.Id,
		IdWil:     p.IdWil,
		Nama:      p.Nama,
		Ump:       p.Ump,
		UmpPkts:   p.UmpPkts,
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(p.UpdatedAt),
	}
}
