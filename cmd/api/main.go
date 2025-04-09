package main

import (
	pb "alerts/api/proto/alert"
	"alerts/internal/cache"
	"alerts/internal/handlers"
	"alerts/internal/kafka"
	"alerts/internal/repository"
	"alerts/internal/server"
	"alerts/internal/service"
	utils "alerts/internal/utils"
	"context"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	httpServer *http.Server
)

func main() {
	// Logger setup
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Kafka logger
	kafkaLogger := kafka.NewKafkaLogger([]string{os.Getenv("KAFKA_BROKER")}, "iot-logs", logger)
	kafkaLogger.LogInfo("✅ Kafka producer started")
	defer kafkaLogger.Close()

	// DB configuration
	portStr := utils.GetEnv("DB_PORT", "5432")
	portInt, err := strconv.Atoi(portStr)
	if err != nil {
		portInt = 5432
	}
	dbConfig := &repository.Config{
		Host:     utils.GetEnv("DB_HOST", "postgres"),
		Port:     portInt,
		User:     utils.GetEnv("DB_USER", "root"),
		Password: utils.GetEnv("DB_PASSWORD", "secret"),
		DBName:   utils.GetEnv("DB_NAME", "alerts"),
	}

	// Initialize DB connection
	queries, db, err := repository.NewDBConnection(dbConfig)
	if err != nil {
		logger.Fatal("❌ Database connection error", zap.Error(err))
	}
	defer db.Close()

	// Initialize Redis client (make sure Redis is running on the correct host and port)
	redisHost := utils.GetEnv("REDIS_HOST", "redis")
	redisClient := cache.NewRedisClient(redisHost, 6379, "", 0)

	// Initialize LRU Cache (size of the local cache)
	localCache := cache.NewLRUCache(100, 30*time.Second)

	// Initialize CacheManager with Redis and Local Cache
	cacheManager := cache.NewCacheManager(redisClient, localCache)

	// Initialize repository, service and handlers
	alertRepo := repository.NewAlertRepository(queries)
	alertService := service.NewAlertService(alertRepo, cacheManager, kafkaLogger)
	alertHandler := handlers.NewHandler(alertService)

	grpcPort := utils.GetEnv("GRPC_PORT", "50051")
	httpPort := utils.GetEnv("HTTP_PORT", "8080")

	// Setup Gin routes
	router := gin.Default()
	server.SetupRoutes(router, alertHandler)

	// Add health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Start gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterAlertServiceServer(grpcServer, alertService)

	// Start servers concurrently
	go startHTTPServer(router, httpPort, logger)
	go startGRPCServer(grpcServer, grpcPort, logger)
	// Wait for termination signal
	<-utils.GracefulShutdown()

	// Graceful shutdown
	shutdownServers(grpcServer, logger)
}

// startHTTPServer starts the Gin HTTP server
func startHTTPServer(router *gin.Engine, port string, logger *zap.Logger) {
	httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("✅ HTTP Server started", zap.String("port", port))

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("❌ HTTP server error", zap.Error(err))
	}
}

// startGRPCServer starts the gRPC server
func startGRPCServer(grpcServer *grpc.Server, port string, logger *zap.Logger) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Fatal("❌ Failed to start gRPC listener", zap.Error(err))
	}
	logger.Info("✅ gRPC Server started", zap.String("port", port))

	if err := grpcServer.Serve(listener); err != nil {
		logger.Error("❌ gRPC server error", zap.Error(err))
	}
}

// shutdownServers shuts down both HTTP and gRPC servers
func shutdownServers(grpcServer *grpc.Server, logger *zap.Logger) {
	grpcServer.GracefulStop()
	logger.Info("✅ gRPC Server stopped gracefully")

	if httpServer != nil {
		if err := httpServer.Shutdown(context.Background()); err != nil {
			logger.Error("❌ HTTP server shutdown error", zap.Error(err))
		} else {
			logger.Info("✅ HTTP Server stopped gracefully")
		}
	}
}
