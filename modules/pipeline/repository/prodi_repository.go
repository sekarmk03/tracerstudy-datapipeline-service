package repository

import (
	"context"
	"log"
	"tracerstudy-datapipeline-service/modules/pipeline/entity"

	"github.com/go-sql-driver/mysql"
	"go.opencensus.io/trace"
	"gorm.io/gorm"
)

type ProdiRepository struct {
	db1 *gorm.DB
	db2 *gorm.DB
}

func NewProdiRepository(db1, db2 *gorm.DB) *ProdiRepository {
	return &ProdiRepository{
		db1: db1,
		db2: db2,
	}
}

type ProdiRepositoryUseCase interface {
	FindAll(ctx context.Context) ([]*entity.OldProdi, error)
	BulkInsert(ctx context.Context, prodi []*entity.NewProdi) error
}

func (p *ProdiRepository) FindAll(ctx context.Context) ([]*entity.OldProdi, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProdiRepository - BulkInsert")
	defer span.End()

	var prodi []*entity.OldProdi
	if err := p.db1.Debug().WithContext(ctxSpan).Find(&prodi).Error; err != nil {
		log.Println("ERROR: [ProdiRepository - FindAll] Internal error:", err)
		return nil, err
	}

	return prodi, nil
}

func (p *ProdiRepository) BulkInsert(ctx context.Context, prodi []*entity.NewProdi) error {
	ctxSpan, span := trace.StartSpan(ctx, "ProdiRepository - BulkInsert")
	defer span.End()

	tx := p.db2.Begin()
	if err := tx.Error; err != nil {
		log.Println("ERROR: [ProdiRepository - BulkInsert] Failed to start transaction:", err)
		return err
	}

	for _, p := range prodi {
		if err := tx.Debug().WithContext(ctxSpan).Create(p).Error; err != nil {
			if gormErr := err.(*mysql.MySQLError); gormErr.Number == 1062 {
				log.Printf("INFO: [ProdiRepository - BulkInsert] Duplicate entry for kode: %s, skipping this entry\n", p.Kode)
				continue
			}
			tx.Rollback()
			log.Println("ERROR: [ProdiRepository - BulkInsert] Failed to insert Prodi data:", err)
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("ERROR: [ProdiRepository - BulkInsert] Failed to commit transaction:", err)
		return err
	}

	return nil
}
