package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// OpenTelemetry 초기화 함수
func initTracerProvider() (*trace.TracerProvider, error) {
	ctx := context.Background()
	otlpEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if otlpEndpoint == "" {
		otlpEndpoint = "tempo:4317"
	}

	log.Println("Attempting to connect to OTLP endpoint:", otlpEndpoint)

	// gRPC 연결 설정 (WithBlock 제거)
	conn, err := grpc.DialContext(ctx, otlpEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.WithBlock(), // 블로킹 제거
	)
	if err != nil {
		log.Printf("Failed to dial OTLP endpoint: %v", err) // 에러 로그 추가
		return nil, err
	}

	log.Println("Creating OTLP exporter...")
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		log.Printf("Failed to create OTLP exporter: %v", err)
		return nil, err
	}

	log.Println("Creating Tracer resource...")
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName("dogu-app"),
		),
	)
	if err != nil {
		log.Printf("Failed to create Tracer resource: %v", err)
		return nil, err
	}

	log.Println("Creating Tracer provider...")
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)

	log.Println("Setting global Tracer provider...")
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	log.Println("Tracer provider initialized successfully.")
	return tp, nil
}

func main() {
	log.Println("Starting main function...")

	// Tracer Provider 초기화
	log.Println("Initializing Tracer provider...")
	tp, err := initTracerProvider()
	if err != nil {
		// Fatal 대신 에러만 기록하고 계속 진행 (서버는 떠야 하므로)
		log.Printf("WARN: failed to initialize tracer provider: %v. Tracing will be disabled.", err)
	} else {
		// 성공 시에만 Shutdown 등록
		defer func() {
			if err := tp.Shutdown(context.Background()); err != nil {
				log.Printf("Error shutting down tracer provider: %v", err)
			}
		}()
	}

	log.Println("Initializing Gin router...")
	r := gin.Default()

	log.Println("Setting up CORS...")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 실제 운영 시에는 더 제한적으로 설정
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "traceparent", "tracestate", "baggage"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	log.Println("Setting up static files and templates...")
	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/template/*")

	log.Println("Setting up routes...")
	r.GET("/", func(c *gin.Context) {
		// Span 생성 (Tracer Provider 초기화 실패 시에도 기본 NoopTracer 사용)
		tracer := otel.Tracer("gin-server")
		ctx, span := tracer.Start(c.Request.Context(), "mainPageHandler")
		defer span.End()

		c.Request = c.Request.WithContext(ctx)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Dogu",
		})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	log.Println("Starting Gin server on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start Gin server: %v", err)
	}
}
