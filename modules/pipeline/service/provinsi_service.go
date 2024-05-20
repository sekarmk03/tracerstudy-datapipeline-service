package service

import (
	"context"
	"log"
	"time"
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/common/errors"
	"tracerstudy-tracer-service/common/utils"
	"tracerstudy-tracer-service/modules/provinsi/entity"
	"tracerstudy-tracer-service/modules/provinsi/repository"
)

type ProvinsiService struct {
	cfg                config.Config
	provinsiRepository repository.ProvinsiRepositoryUseCase
}

type ProvinsiServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Provinsi, error)
	FindByIdWil(ctx context.Context, idWil string) (*entity.Provinsi, error)
	Create(ctx context.Context, idWil, nama string, ump uint64) (*entity.Provinsi, error)
	Update(ctx context.Context, idWil string, fields *entity.Provinsi) (*entity.Provinsi, error)
	Delete(ctx context.Context, idWil string) error
}

func NewProvinsiService(cfg config.Config, provinsiRepository repository.ProvinsiRepositoryUseCase) *ProvinsiService {
	return &ProvinsiService{
		cfg:                cfg,
		provinsiRepository: provinsiRepository,
	}
}

func (svc *ProvinsiService) FindAll(ctx context.Context, req any) ([]*entity.Provinsi, error) {
	res, err := svc.provinsiRepository.FindAll(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProvinsiService - FindAll] Error while find all provinsi: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ProvinsiService) FindByIdWil(ctx context.Context, idWil string) (*entity.Provinsi, error) {
	res, err := svc.provinsiRepository.FindByIdWil(ctx, idWil)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProvinsiService - FindByIdWil] Error while find provinsi by IdWil: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ProvinsiService) Create(ctx context.Context, idWil, nama string, ump uint64) (*entity.Provinsi, error) {
	provinsi := &entity.Provinsi{
		IdWil:     idWil,
		Nama:      nama,
		Ump:       ump,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res, err := svc.provinsiRepository.Create(ctx, provinsi)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProvinsiService - Create] Error while create provinsi: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ProvinsiService) Update(ctx context.Context, idWil string, fields *entity.Provinsi) (*entity.Provinsi, error) {
	provinsi, err := svc.provinsiRepository.FindByIdWil(ctx, idWil)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProvinsiService - Update] Error while find provinsi by IdWil: ", parseError.Message)
		return nil, err
	}

	updatedMap := make(map[string]interface{})

	utils.AddItemToMap(updatedMap, "nama", fields.Nama)
	utils.AddItemToMap(updatedMap, "ump", fields.Ump)

	res, err := svc.provinsiRepository.Update(ctx, provinsi, updatedMap)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProvinsiService - Update] Error while update provinsi: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ProvinsiService) Delete(ctx context.Context, idWil string) error {
	err := svc.provinsiRepository.Delete(ctx, idWil)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProvinsiService - Delete] Error while delete provinsi: ", parseError.Message)
		return err
	}

	return nil
}
