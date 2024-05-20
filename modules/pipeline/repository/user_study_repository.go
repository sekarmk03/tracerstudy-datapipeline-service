package repository

import (
	"context"
	"log"
	"tracerstudy-datapipeline-service/modules/pipeline/entity"

	"go.opencensus.io/trace"
	"gorm.io/gorm"
)

type UserStudyRepository struct {
	db1 *gorm.DB
	db2 *gorm.DB
}

func NewUserStudyRepository(db1, db2 *gorm.DB) *UserStudyRepository {
	return &UserStudyRepository{
		db1: db1,
		db2: db2,
	}
}

type UserStudyRepositoryUseCase interface {
	FindAll(ctx context.Context) ([]*entity.OldUserStudy, error)
	CheckExist(ctx context.Context, emailResponden, hpResponden, nimLulusan string) (bool, error)
	BulkInsert(ctx context.Context, userStudy []*entity.NewUserStudy) (uint64, error)
}

func (us *UserStudyRepository) FindAll(ctx context.Context) ([]*entity.OldUserStudy, error) {
	ctxSpan, span := trace.StartSpan(ctx, "UserStudyRepository - FindAll")
	defer span.End()

	var userStudy []*entity.OldUserStudy
	if err := us.db1.Debug().WithContext(ctxSpan).Order("created_at asc").Find(&userStudy).Error; err != nil {
		log.Println("ERROR: [UserStudyRepository - FindAll] Internal server error:", err)
		return nil, err
	}

	return userStudy, nil
}

func (us *UserStudyRepository) BulkInsert(ctx context.Context, userStudy []*entity.NewUserStudy) (uint64, error) {
	ctxSpan, span := trace.StartSpan(ctx, "UserStudyRepository - BulkInsert")
	defer span.End()

	tx := us.db2.Begin()
	if err := tx.Error; err != nil {
		log.Println("ERROR: [UserStudyRepository - BulkInsert] Internal server error:", err)
		return 0, err
	}

	count := uint64(len(userStudy))

	for _, u := range userStudy {
		exist, err := us.CheckExist(ctx, u.EmailResponden, u.HpResponden, u.NimLulusan)
		if err != nil {
			count--
			tx.Rollback()
			log.Println("ERROR: [UserStudyRepository - BulkInsert] Internal server error:", err)
			return 0, err
		}
		if exist {
			count--
			log.Printf("INFO: [PktsRepository - BulkInsert] Data already exist for user  %s and lulusan %s, skipping this entry\n", u.EmailResponden, u.NimLulusan)
			continue
		}

		if err := tx.Debug().WithContext(ctxSpan).Create(&u).Error; err != nil {
			count--
			tx.Rollback()
			log.Println("ERROR: [UserStudyRepository - BulkInsert] Internal server error:", err)
			return 0, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("ERROR: [UserStudyRepository - BulkInsert] Internal server error:", err)
		return 0, err
	}

	return count, nil
}

func (us *UserStudyRepository) CheckExist(ctx context.Context, emailResponden, hpResponden, nimLulusan string) (bool, error) {
	ctxSpan, span := trace.StartSpan(ctx, "UserStudyRepository - CheckExist")
	defer span.End()

	var userStudy entity.NewUserStudy
	if err := us.db2.Debug().WithContext(ctxSpan).
		Where("(email_responden = ? AND nim_lulusan = ?) OR (hp_responden = ? AND nim_lulusan = ?)", emailResponden, nimLulusan, hpResponden, nimLulusan).
		First(&userStudy).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		log.Println("ERROR: [UserStudyRepository - CheckExist] Internal server error:", err)
		return false, err
	}

	return true, nil
}
