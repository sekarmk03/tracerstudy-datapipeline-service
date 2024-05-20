package builder

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/provinsi/handler"
	"tracerstudy-tracer-service/modules/provinsi/repository"
	"tracerstudy-tracer-service/modules/provinsi/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildProvinsiHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.ProvinsiHandler {
	provinsiRepo := repository.NewProvinsiRepository(db)
	provinsiSvc := service.NewProvinsiService(cfg, provinsiRepo)

	return handler.NewProvinsiHandler(cfg, provinsiSvc)
}
