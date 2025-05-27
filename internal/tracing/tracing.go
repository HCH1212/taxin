package tracing

import (
	"context"
	"log"

	"github.com/HCH1212/taxin/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// InitTracer 初始化 OpenTelemetry 追踪器
func InitTracer(ctx context.Context, serviceName string) (*sdktrace.TracerProvider, error) {
	log.Println(config.GetConf().Jeager.Address)

	// 连接 Jaeger OTLP 端口
	conn, err := grpc.DialContext(
		ctx,
		config.GetConf().Jeager.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Printf("Failed to connect to Jaeger: %v", err)
		return nil, err
	}

	// 创建 OTLP 导出器
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		log.Printf("Failed to create exporter: %v", err)
		return nil, err
	}

	// 创建资源
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String("1.0.0"),
			attribute.String("environment", "development"),
		),
	)
	if err != nil {
		log.Printf("Failed to create resource: %v", err)
		return nil, err
	}

	// 创建追踪器提供者
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()), // 100% 采样率
	)

	// 设置全局追踪器
	otel.SetTracerProvider(tp)
	return tp, nil
}
