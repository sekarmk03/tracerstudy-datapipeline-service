package builder

import (
	"tracerstudy-datapipeline-service/common/config"
	"tracerstudy-datapipeline-service/modules/pipeline/handler"
	"tracerstudy-datapipeline-service/modules/pipeline/repository"
	"tracerstudy-datapipeline-service/modules/pipeline/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildPipelineHandler(cfg config.Config, db1, db2 *gorm.DB, grpcConn *grpc.ClientConn) *handler.PipelineHandler {
	kabKotaRepo := repository.NewKabKotaRepository(db1, db2)
	provinsiRepo := repository.NewProvinsiRepository(db1, db2)
	prodiRepo := repository.NewProdiRepository(db1, db2)
	mhsBiodataSvc := service.NewMhsBiodataService(cfg)
	respondenRepo := repository.NewRespondenRepository(db1, db2)
	userStudyRepo := repository.NewUserStudyRepository(db1, db2)
	pktsRepo := repository.NewPktsRepository(db1, db2)

	pipelineSvc := service.NewPipelineService(
		cfg,
		kabKotaRepo,
		provinsiRepo,
		prodiRepo,
		mhsBiodataSvc,
		respondenRepo,
		userStudyRepo,
		pktsRepo,
	)

	return handler.NewPipelineHandler(cfg, pipelineSvc)
}
