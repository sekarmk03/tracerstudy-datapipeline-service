package repository

import (
	"context"
	"errors"
	"log"
	"time"
	"tracerstudy-tracer-service/modules/provinsi/entity"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProvinsiRepository struct {
	db *gorm.DB
}

func NewProvinsiRepository(db *gorm.DB) *ProvinsiRepository {
	return &ProvinsiRepository{
		db: db,
	}
}

type ProvinsiRepositoryUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Provinsi, error)
	FindByIdWil(ctx context.Context, idWil string) (*entity.Provinsi, error)
	Create(ctx context.Context, req *entity.Provinsi) (*entity.Provinsi, error)
	Update(ctx context.Context, provinsi *entity.Provinsi, updatedFields map[string]interface{}) (*entity.Provinsi, error)
	Delete(ctx context.Context, idWil string) error
}

func (p *ProvinsiRepository) FindAll(ctx context.Context, req any) ([]*entity.Provinsi, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProvinsiRepository - FindAll")
	defer span.End()

	var provinsi []*entity.Provinsi
	if err := p.db.Debug().WithContext(ctxSpan).Find(&provinsi).Error; err != nil {
		log.Println("ERROR: [ProvinsiRepository - FindAll] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return provinsi, nil
}

func (p *ProvinsiRepository) FindByIdWil(ctx context.Context, idWil string) (*entity.Provinsi, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProvinsiRepository - FindByIdWil")
	defer span.End()

	var provinsi entity.Provinsi
	if err := p.db.Debug().WithContext(ctxSpan).Where("id_wil = ?", idWil).First(&provinsi).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [ProvinsiRepository - FindByIdWil] Record not found for idWil", idWil)
			return nil, status.Errorf(codes.NotFound, "record not found for idWil %s", idWil)
		}
		log.Println("ERROR: [ProvinsiRepository - FindByIdWil] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &provinsi, nil
}

func (p *ProvinsiRepository) Create(ctx context.Context, req *entity.Provinsi) (*entity.Provinsi, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProvinsiRepository - Create")
	defer span.End()

	if err := p.db.Debug().WithContext(ctxSpan).Create(req).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Println("ERROR: [ProvinsiRepository - Create] Duplicated:", err)
			return nil, status.Errorf(codes.AlreadyExists, "duplicate key error: %v", err)
		}
		log.Println("ERROR: [ProvinsiRepository - Create] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return req, nil
}

func (p *ProvinsiRepository) Update(ctx context.Context, provinsi *entity.Provinsi, updatedFields map[string]interface{}) (*entity.Provinsi, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ProvinsiRepository - Update")
	defer span.End()

	updatedFields["updated_at"] = time.Now()
	if err := p.db.Debug().WithContext(ctxSpan).Model(&provinsi).Where("id_wil", provinsi.IdWil).Updates(updatedFields).Error; err != nil {
		if errors.Is(err, gorm.ErrInvalidValue) {
			log.Println("ERROR: [ProvinsiRepository - Update] Invalid value:", err)
			return nil, status.Errorf(codes.InvalidArgument, "invalid value: %v", err)
		}
		log.Println("ERROR: [ProvinsiRepository - Update] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return provinsi, nil
}

func (p *ProvinsiRepository) Delete(ctx context.Context, idWil string) error {
	ctxSpan, span := trace.StartSpan(ctx, "ProvinsiRepository - Delete")
	defer span.End()

	if err := p.db.Debug().WithContext(ctxSpan).Where("id_wil = ?", idWil).Delete(&entity.Provinsi{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [ProvinsiRepository - Delete] Record not found for idWil", idWil)
			return status.Errorf(codes.NotFound, "record not found for idWil %s", idWil)
		}
		log.Println("ERROR: [ProvinsiRepository - Delete] Internal server error:", err)
		return status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return nil
}
