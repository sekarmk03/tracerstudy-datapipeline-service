package repository

import (
	"context"
	"log"
	"tracerstudy-datapipeline-service/modules/pipeline/entity"

	"github.com/go-sql-driver/mysql"
	"go.opencensus.io/trace"
	"gorm.io/gorm"
)

type RespondenRepository struct {
	db1 *gorm.DB
	db2 *gorm.DB
}

func NewRespondenRepository(db1, db2 *gorm.DB) *RespondenRepository {
	return &RespondenRepository{
		db1: db1,
		db2: db2,
	}
}

type RespondenRepositoryUseCase interface {
	FindUnupdated(ctx context.Context) ([]*entity.OldResponden, error)
	Update(ctx context.Context, nim string, responden *entity.OldResponden, updatedFields map[string]interface{}) error
	UpdateStatusUpdate(ctx context.Context, statusFrom, statusTo string) error
	FindAll(ctx context.Context) ([]*entity.OldResponden, error)
	BulkInsert(ctx context.Context, responden []*entity.NewResponden) error
}

func (r *RespondenRepository) FindUnupdated(ctx context.Context) ([]*entity.OldResponden, error) {
	ctxSpan, span := trace.StartSpan(ctx, "RespondenRepository - FindUnupdated")
	defer span.End()

	var responden []*entity.OldResponden
	if err := r.db1.Debug().WithContext(ctxSpan).Where("thnmasuk IS NULL OR lamastd IS NULL").Limit(500).Find(&responden).Error; err != nil {
		log.Println("ERROR: [RespondenRepository - FindUnupdated] Internal server error:", err)
		return nil, err
	}

	return responden, nil
}

func (r *RespondenRepository) Update(ctx context.Context, nim string, responden *entity.OldResponden, updatedFields map[string]interface{}) error {
	ctxSpan, span := trace.StartSpan(ctx, "RespondenRepository - Update")
	defer span.End()

	if err := r.db1.Debug().WithContext(ctxSpan).Model(&responden).Where("nim = ?", nim).Updates(updatedFields).Error; err != nil {
		log.Println("ERROR: [RespondenRepository - Update] Internal server error:", err)
		return err
	}

	return nil
}

func (r *RespondenRepository) UpdateStatusUpdate(ctx context.Context, statusFrom, statusTo string) error {
	ctxSpan, span := trace.StartSpan(ctx, "RespondenRepository - UpdateStatusUpdate")
	defer span.End()

	if err := r.db1.Debug().WithContext(ctxSpan).Model(&entity.OldResponden{}).Where("status_update = ?", statusFrom).Update("status_update", statusTo).Error; err != nil {
		log.Println("ERROR: [RespondenRepository - UpdateStatusUpdate] Internal server error:", err)
		return err
	}

	return nil
}

func (r *RespondenRepository) FindAll(ctx context.Context) ([]*entity.OldResponden, error) {
	ctxSpan, span := trace.StartSpan(ctx, "RespondenRepository - FindAll")
	defer span.End()

	var responden []*entity.OldResponden
	if err := r.db1.Debug().WithContext(ctxSpan).Order("created_at asc").Find(&responden).Error; err != nil {
		log.Println("ERROR: [RespondenRepository - FindAll] Internal server error:", err)
		return nil, err
	}

	return responden, nil
}

func (r *RespondenRepository) BulkInsert(ctx context.Context, responden []*entity.NewResponden) error {
	ctxSpan, span := trace.StartSpan(ctx, "RespondenRepository - BulkInsert")
	defer span.End()

	tx := r.db2.Begin()
	if err := tx.Error; err != nil {
		log.Println("ERROR: [RespondenRepository - BulkInsert] Failed to start transaction:", err)
		return err
	}

	for _, r := range responden {
		if err := tx.Debug().WithContext(ctxSpan).Create(r).Error; err != nil {
			if gormErr := err.(*mysql.MySQLError); gormErr.Number == 1062 {
				log.Printf("INFO: [RespondenRepository - BulkInsert] Duplicate entry for nim: %s, skipping this entry\n", r.Nim)
				continue
			}
			tx.Rollback()
			log.Println("ERROR: [RespondenRepository - BulkInsert] Failed to insert Responden data:", err)
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("ERROR: [RespondenRepository - BulkInsert] Failed to commit transaction:", err)
		return err
	}

	return nil
}
