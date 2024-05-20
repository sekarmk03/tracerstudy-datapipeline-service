package handler

import (
	"context"
	"log"
	"net/http"
	"tracerstudy-datapipeline-service/common/config"
	"tracerstudy-datapipeline-service/common/errors"
	"tracerstudy-datapipeline-service/modules/pipeline/service"
	"tracerstudy-datapipeline-service/pb"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PipelineHandler struct {
	pb.UnimplementedPipelineServiceServer
	config      config.Config
	pipelineSvc service.PipelineServiceUseCase
}

func NewPipelineHandler(config config.Config, pipelineService service.PipelineServiceUseCase) *PipelineHandler {
	return &PipelineHandler{
		config:      config,
		pipelineSvc: pipelineService,
	}
}

func (ph *PipelineHandler) KabKotaPipeline(ctx context.Context, req *emptypb.Empty) (*pb.PipelineServiceResponse, error) {
	rows, err := ph.pipelineSvc.KabKotaPipeline(ctx)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PipelineHandler - KabKotaPipeline] Error while run KabKota pipeline: ", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.PipelineServiceResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
			Rows: 0,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.PipelineServiceResponse{
		Code:    uint32(http.StatusOK),
		Message: "Pipeline for KabKota has been successfully executed",
		Rows: rows,
	}, nil
}

func (ph *PipelineHandler) ProvinsiPipeline(ctx context.Context, req *emptypb.Empty) (*pb.PipelineServiceResponse, error) {
	err := ph.pipelineSvc.ProvinsiPipeline(ctx)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PipelineHandler - ProvinsiPipeline] Error while run Provinsi pipeline: ", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.PipelineServiceResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.PipelineServiceResponse{
		Code:    uint32(http.StatusOK),
		Message: "Pipeline for Provinsi has been successfully executed",
	}, nil
}

func (ph *PipelineHandler) ProdiPipeline(ctx context.Context, req *emptypb.Empty) (*pb.PipelineServiceResponse, error) {
	rows, err := ph.pipelineSvc.ProdiPipeline(ctx)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PipelineHandler - ProdiPipeline] Error while run Prodi pipeline: ", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.PipelineServiceResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
			Rows: 0,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.PipelineServiceResponse{
		Code:    uint32(http.StatusOK),
		Message: "Pipeline for Prodi has been successfully executed",
		Rows: rows,
	}, nil
}

func (ph *PipelineHandler) UserStudyPipeline(ctx context.Context, req *emptypb.Empty) (*pb.PipelineServiceResponse, error) {
	err := ph.pipelineSvc.UserStudyPipeline(ctx)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PipelineHandler - UserStudyPipeline] Error while run User Study pipeline: ", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.PipelineServiceResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.PipelineServiceResponse{
		Code:    uint32(http.StatusOK),
		Message: "Pipeline for User Study has been successfully executed",
	}, nil
}

func (ph *PipelineHandler) SiakUpdateRespondenPipeline(ctx context.Context, req *emptypb.Empty) (*pb.PipelineServiceResponse, error) {
	err := ph.pipelineSvc.SiakUpdateRespondenPipeline(ctx)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PipelineHandler - SiakUpdateRespondenPipeline] Error while update Responden data from Siak API: ", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.PipelineServiceResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.PipelineServiceResponse{
		Code:    uint32(http.StatusOK),
		Message: "Pipeline for Siak Update Responden has been successfully executed by 500 rows per batch",
	}, nil
}

func (ph *PipelineHandler) RespondenPipeline(ctx context.Context, req *emptypb.Empty) (*pb.PipelineServiceResponse, error) {
	err := ph.pipelineSvc.RespondenPipeline(ctx)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PipelineHandler - RespondenPipeline] Error while run Responden pipeline: ", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.PipelineServiceResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.PipelineServiceResponse{
		Code:    uint32(http.StatusOK),
		Message: "Pipeline for Responden has been successfully executed",
	}, nil
}

func (ph *PipelineHandler) PKTSPipeline(ctx context.Context, req *emptypb.Empty) (*pb.PipelineServiceResponse, error) {
	rows, err := ph.pipelineSvc.PKTSPipeline(ctx)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PipelineHandler - PKTSPipeline] Error while run PKTS pipeline:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.PipelineServiceResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
			Rows: 0,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.PipelineServiceResponse{
		Code:    uint32(http.StatusOK),
		Message: "Pipeline for PKTS has been successfully executed by 500 rows per batch",
		Rows: rows,
	}, nil
}
