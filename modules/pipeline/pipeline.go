package provinsi

import (
	"tracerstudy-datapipeline-service/common/config"
	"tracerstudy-datapipeline-service/modules/pipeline/builder"
	"tracerstudy-datapipeline-service/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db1, db2 *gorm.DB, grpcConn *grpc.ClientConn) {
	pipeline := builder.BuildPipelineHandler(cfg, db1, db2, grpcConn)
	pb.RegisterPipelineServiceServer(server, pipeline)
}
