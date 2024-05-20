package main

import (
	"fmt"
	"tracerstudy-datapipeline-service/common/config"

	gormConn "tracerstudy-datapipeline-service/common/gorm"
	commonJwt "tracerstudy-datapipeline-service/common/jwt"
	"tracerstudy-datapipeline-service/common/mysql"
	"tracerstudy-datapipeline-service/server"

	pipelineModule "tracerstudy-datapipeline-service/modules/pipeline"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	cfg, cerr := config.NewConfig(".env")
	checkError(cerr)

	splash(cfg)

	dsn1, derr1 := mysql.NewPool(cfg.MySQL1.User, cfg.MySQL1.Password, cfg.MySQL1.Host, cfg.MySQL1.Port, cfg.MySQL1.Name)
	checkError(derr1)

	db1, gerr1 := gormConn.NewMySQLGormDB(dsn1)
	checkError(gerr1)

	dsn2, derr2 := mysql.NewPool(cfg.MySQL2.User, cfg.MySQL2.Password, cfg.MySQL2.Host, cfg.MySQL2.Port, cfg.MySQL2.Name)
	checkError(derr2)

	db2, gerr2 := gormConn.NewMySQLGormDB(dsn2)
	checkError(gerr2)

	jwtManager := commonJwt.NewJWT(cfg.JWT.JwtSecretKey, cfg.JWT.TokenDuration)

	grpcServer := server.NewGrpcServer(cfg.Port.GRPC, jwtManager)
	grpcConn := server.InitGRPCConn(fmt.Sprintf("127.0.0.1:%v", cfg.Port.GRPC), false, "")

	registerGrpcHandlers(grpcServer.Server, *cfg, db1, db2, grpcConn)

	_ = grpcServer.Run()
	_ = grpcServer.AwaitTermination()
}

func registerGrpcHandlers(server *grpc.Server, cfg config.Config, db1, db2 *gorm.DB, grpcConn *grpc.ClientConn) {
	pipelineModule.InitGrpc(server, cfg, db1, db2, grpcConn)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func splash(cfg *config.Config) {
	version := "1.0.0"
	colorReset := "\033[0m"
	// colorBlue := "\033[34m"
	colorCyan := "\033[36m"

	fmt.Printf(`
    __                                                    __                    __
  _/  |_ _______ ______   ____  ______ _______     ______/  |__  __   __       |  \ __   __
  \   __\\_  __ \\___  \_/ ___\/  __  \\_  __ \   /  ___/\   __\|  | |  |   ___|  ||  | |  |
   |  |   |  | \/  / ___\  \___\   ___/ |  | \/   \___ \  |  |  |  |_|  | /  __   ||  |_|  | >>
   |__|   |__|    (____  /\__  >\___  > |__|     /____  > |__|  | ______|(  ____  )\__   __/
                       \/    \/     \/                \/        \/        \/    \/ _ /  /    v%s
                                                                                  / ___/
	`, version)

	// fmt.Println(colorBlue, fmt.Sprintf(`⇨ REST server started on port :%s`, cfg.Port.REST))
	fmt.Println(colorCyan, fmt.Sprintf(`⇨ GRPC data pipeline service server started on port :%s`, cfg.Port.GRPC))
	fmt.Println(colorReset, "")
}
