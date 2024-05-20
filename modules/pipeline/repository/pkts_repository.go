package repository

import (
	"context"
	"log"
	"tracerstudy-datapipeline-service/modules/pipeline/entity"

	"go.opencensus.io/trace"
	"gorm.io/gorm"
)

type PktsRepository struct {
	db1 *gorm.DB
	db2 *gorm.DB
}

func NewPktsRepository(db1, db2 *gorm.DB) *PktsRepository {
	return &PktsRepository{
		db1: db1,
		db2: db2,
	}
}

type PktsRepositoryUseCase interface {
	FindAll(ctx context.Context) ([]*entity.OldPkts, error)
	CheckExist(ctx context.Context, nim string) (bool, error)
	BulkInsert(ctx context.Context, pkts []*entity.NewPkts) (uint64, error)
}

func (p *PktsRepository) FindAll(ctx context.Context) ([]*entity.OldPkts, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PktsRepository - FindAll")
	defer span.End()

	var pkts []*entity.OldPkts
	if err := p.db1.Debug().WithContext(ctxSpan).Order("created_at asc").Find(&pkts).Error; err != nil {
		log.Println("ERROR: [PktsRepository - FindAll] Internal server error:", err)
		return nil, err
	}

	return pkts, nil
}

func (p *PktsRepository) BulkInsert(ctx context.Context, pkts []*entity.NewPkts) (uint64, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PktsRepository - BulkInsert")
	defer span.End()

	tx := p.db2.Begin()
	if err := tx.Error; err != nil {
		log.Println("ERROR: [PktsRepository - BulkInsert] Internal server error:", err)
		return 0, err
	}

	count := uint64(len(pkts))

	for _, pk := range pkts {
		exist, err := p.CheckExist(ctx, pk.Nim)
		if err != nil {
			tx.Rollback()
			log.Println("ERROR: [PktsRepository - BulkInsert] Internal server error:", err)
			count--
			return 0, err
		}
		if exist {
			log.Printf("INFO: [PktsRepository - BulkInsert] Data already exist for nim %s, skipping this entry\n", pk.Nim)
			count--
			continue
		}

		if err := tx.Debug().WithContext(ctxSpan).Create(&pk).Error; err != nil {
			tx.Rollback()
			log.Println("ERROR: [PktsRepository - BulkInsert] Internal server error:", err)
			count--
			return 0, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("ERROR: [PktsRepository - BulkInsert] Internal server error:", err)
		return 0, err
	}

	return count, nil
}

func (p *PktsRepository) CheckExist(ctx context.Context, nim string) (bool, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PktsRepository - CheckExist")
	defer span.End()

	var pkts entity.NewPkts
	if err := p.db2.Debug().WithContext(ctxSpan).Where("nim = ?", nim).First(&pkts).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		log.Println("ERROR: [PktsRepository - CheckExist] Internal server error:", err)
		return false, err
	}

	return true, nil
}
