package provinsi

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/provinsi/builder"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	provinsi := builder.BuildProvinsiHandler(cfg, db, grpcConn)
	pb.RegisterProvinsiServiceServer(server, provinsi)
}
