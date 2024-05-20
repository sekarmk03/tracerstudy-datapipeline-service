package repository

import (
	"context"
	"log"
	"tracerstudy-datapipeline-service/modules/pipeline/entity"

	"github.com/go-sql-driver/mysql"
	"go.opencensus.io/trace"
	"gorm.io/gorm"
)

type ProvinsiRepository struct {
	db1 *gorm.DB
	db2 *gorm.DB
}

func NewProvinsiRepository(db1, db2 *gorm.DB) *ProvinsiRepository {
	return &ProvinsiRepository{
		db1: db1,
		db2: db2,
	}
}

type ProvinsiRepositoryUseCase interface {
	FindAll(ctx context.Context) ([]*entity.OldProvinsi, error)
	BulkInsert(ctx context.Context, provinsi []*entity.NewProvinsi) (uint64, error)
}

func (p *ProvinsiRepository) FindAll(ctx context.Context) ([]*entity.OldProvinsi, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProvinsiRepository - BulkInsert")
	defer span.End()

	var provinsi []*entity.OldProvinsi
	if err := p.db1.Debug().WithContext(ctxSpan).Find(&provinsi).Error; err != nil {
		log.Println("ERROR: [ProvinsiRepository - FindAll] Internal error:", err)
		return nil, err
	}

	return provinsi, nil
}

func (p *ProvinsiRepository) BulkInsert(ctx context.Context, provinsi []*entity.NewProvinsi) (uint64, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProvinsiRepository - BulkInsert")
	defer span.End()

	tx := p.db2.Begin()
	if err := tx.Error; err != nil {
		log.Println("ERROR: [ProvinsiRepository - BulkInsert] Failed to start transaction:", err)
		return 0, err
	}

	count := uint64(len(provinsi))

	for _, p := range provinsi {
		if err := tx.Debug().WithContext(ctxSpan).Create(p).Error; err != nil {
			count--
			if gormErr := err.(*mysql.MySQLError); gormErr.Number == 1062 {
				log.Printf("INFO: [ProvinsiRepository - BulkInsert] Duplicate entry for id_wil: %s, skipping this entry\n", p.IdWil)
				continue
			}
			tx.Rollback()
			log.Println("ERROR: [ProvinsiRepository - BulkInsert] Failed to insert Provinsi data:", err)
			return 0, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("ERROR: [ProvinsiRepository - BulkInsert] Failed to commit transaction:", err)
		return 0, err
	}

	return count, nil
}
