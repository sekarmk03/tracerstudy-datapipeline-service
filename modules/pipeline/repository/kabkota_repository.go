package repository

import (
	"context"
	"log"
	"tracerstudy-datapipeline-service/modules/pipeline/entity"

	"github.com/go-sql-driver/mysql"
	"go.opencensus.io/trace"
	"gorm.io/gorm"
)

type KabKotaRepository struct {
	db1 *gorm.DB
	db2 *gorm.DB
}

func NewKabKotaRepository(db1, db2 *gorm.DB) *KabKotaRepository {
	return &KabKotaRepository{
		db1: db1,
		db2: db2,
	}
}

type KabKotaRepositoryUseCase interface {
	FindAll(ctx context.Context) ([]*entity.OldKabkota, error)
	BulkInsert(ctx context.Context, kabkota []*entity.NewKabkota) error
}

func (k *KabKotaRepository) FindAll(ctx context.Context) ([]*entity.OldKabkota, error) {
	ctxSpan, span := trace.StartSpan(ctx, "KabkotaRepository - BulkInsert")
	defer span.End()

	var kabkota []*entity.OldKabkota
	if err := k.db1.Debug().WithContext(ctxSpan).Find(&kabkota).Error; err != nil {
		log.Println("ERROR: [KabKotaRepository - FindAll] Internal error:", err)
		return nil, err
	}

	return kabkota, nil
}

func (k *KabKotaRepository) BulkInsert(ctx context.Context, kabkota []*entity.NewKabkota) error {
	ctxSpan, span := trace.StartSpan(ctx, "KabkotaRepository - BulkInsert")
	defer span.End()

	tx := k.db2.Begin()
	if err := tx.Error; err != nil {
		log.Println("ERROR: [KabKotaRepository - BulkInsert] Failed to start transaction:", err)
		return err
	}

	for _, kk := range kabkota {
		if err := tx.Debug().WithContext(ctxSpan).Create(kk).Error; err != nil {
			if gormErr := err.(*mysql.MySQLError); gormErr.Number == 1062 {
				log.Printf("INFO: [KabKotaRepository - BulkInsert] Duplicate entry for id_wil: %s, skipping this entry\n", kk.IdWil)
				continue
			}
			tx.Rollback()
			log.Println("ERROR: [KabKotaRepository - BulkInsert] Failed to insert Kabkota data:", err)
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("ERROR: [KabKotaRepository - BulkInsert] Failed to commit transaction:", err)
		return err
	}

	return nil
}
