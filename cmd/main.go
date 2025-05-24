package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"net/http"
	_ "net/http/pprof"

	"github.com/HCH1212/taxin/api/pb/system"
	"github.com/HCH1212/taxin/api/pb/user"
	"github.com/HCH1212/taxin/internal/dao"
	"github.com/HCH1212/taxin/internal/middleware"
	"github.com/HCH1212/taxin/internal/service"
	"github.com/HCH1212/taxin/internal/tracing"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	_ = godotenv.Load()

	// 启动 pprof 服务
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	app := fx.New(
		// 初始化数据库和 Redis
		fx.Invoke(func() {
			dao.InitDB()
			dao.InitRedis()
		}),
		// 提供 Jaeger 追踪器
		fx.Provide(
			newTracerProvider,
		),
		// 提供监听套接字和 gRPC 服务器
		fx.Provide(
			newListener,
			newGRPCServer,
		),
		// 触发服务器启动
		fx.Invoke(func(*grpc.Server) {}),
		// 禁用日志
		fx.NopLogger,
	)

	app.Run()
}

// 创建监听套接字
func newListener() (net.Listener, error) {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	return lis, nil
}

// 创建 gRPC 服务器
func newGRPCServer(lc fx.Lifecycle, lis net.Listener) *grpc.Server {
	s := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.UnaryInterceptor(middleware.AuthInterceptor()), // 认证拦截器
	)

	user.RegisterUserServiceServer(s, &service.UserService{})
	system.RegisterSystemServiceServer(s, &service.SystemService{})
	reflection.Register(s)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Server is listening on port 50052")
			go func() {
				if err := s.Serve(lis); err != nil {
					log.Printf("failed to serve: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			s.GracefulStop()
			return nil
		},
	})

	return s
}

// 创建 Jaeger 追踪器
func newTracerProvider(lc fx.Lifecycle) (func(context.Context) error, error) {
	ctx := context.Background()
	tp, err := tracing.InitTracer(ctx, "taxin")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize tracer: %w", err)
	}

	// 设置全局的追踪传播器
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// 注册生命周期钩子
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			fmt.Println("Shutting down tracer provider")
			return tp.Shutdown(ctx)
		},
	})

	return tp.Shutdown, nil
}
