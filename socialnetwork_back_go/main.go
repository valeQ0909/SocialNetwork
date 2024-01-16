package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"socialnetwork_back_go/router"
	"socialnetwork_back_go/utils"
	"strings"
)
import pb "socialnetwork_back_go/proto"

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "Hello World !!!" + req.RequestName}, nil
}

func main() {
	//解析配置文件
	cfg, err := utils.ParseConfig("./config/engineConfig.json")
	if err != nil {
		panic(err.Error())
	}

	engine := gin.Default()
	// 设置跨域
	engine.Use(Cors())
	// 初始化路由
	router.InitRouter(engine)
	// 运行web服务
	pprof.Register(engine)
	// EngineHost是项目的运行地址， EnginePort是项目运行的端口
	engine.Run(cfg.EngineHost + ":" + cfg.EnginePort)

	//开启端口
	listen, _ := net.Listen("tcp", ":3001")
	// 创建grpc服务
	grpcServer := grpc.NewServer()
	// 在grpc服务端中注册服务
	pb.RegisterSayHelloServer(grpcServer, &server{})
	//启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}

// Cors 跨域访问
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {

		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")

		var headerKeys []string
		for key, _ := range context.Request.Header {
			headerKeys = append(headerKeys, key)
		}

		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			context.Writer.Header().Set("Access-Control-Allow-Methods", "*")
			context.Writer.Header().Set("Access-Control-Allow-Headers", "*")
			context.Writer.Header().Set("Access-Control-Max-Age", "3600")
			context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			//context.Set("content-type", "application/json") // 设置返回格式是json
		}

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		} else {
			// 处理请求
			context.Next()
		}

	}
}
