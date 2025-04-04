package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "alerts/api/proto/alert" // Import generated gRPC code
	"alerts/internal/handlers"
	"alerts/internal/repository"
	"alerts/internal/server"
	"alerts/internal/service"
	utils "alerts/internal/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// Load Configuration (Port Numbers)
	dbConfig := &repository.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "root",
		Password: "secret",
		DBName:   "alerts",
	}

	// Initialize Database Connection
	queries, db, err := repository.NewDBConnection(dbConfig)
	if err != nil {
		fmt.Printf("❌ Database connection error: %v\n", err)
		return
	}

	defer db.Close()

	// Load Environment Variables
	utils.LoadConfig(".env") // Load environment variables from .env file
	// Initialize Repository and Service

	alertRepo := repository.NewAlertRepository(queries) // Create a new AlertRepository instance
	alertService := service.NewAlertService(alertRepo)  // Create a new AlertService instance
	// Create a new handler instance with the AlertService.
	alertHandler := handlers.NewHandler(alertService)

	grpcPort := getEnv("GRPC_PORT", "50051")
	httpPort := getEnv("HTTP_PORT", "8080")
	// Channel to listen for termination signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Create Gin Router
	router := gin.Default()
	server.SetupRoutes(router, alertHandler) // Set up HTTP routes

	// Start gRPC Server
	grpcServer := grpc.NewServer()
	pb.RegisterAlertServiceServer(grpcServer, alertService) // Register gRPC service

	// Start HTTP & gRPC Servers
	go startHTTPServer(router, httpPort)
	go startGRPCServer(grpcServer, grpcPort)

	// Block until termination signal is received
	<-stop

	// Graceful Shutdown
	shutdownServers(grpcServer)
}

// startHTTPServer starts the Gin HTTP server
func startHTTPServer(router *gin.Engine, port string) {
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("✅ HTTP Server started on port %s\n", port)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("❌ HTTP server error: %v", err)
	}
}

// startGRPCServer starts the gRPC server
func startGRPCServer(grpcServer *grpc.Server, port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("❌ gRPC server error: %v", err)
	}
	fmt.Printf("✅ gRPC Server started on port %s\n", port)

	if err := grpcServer.Serve(listener); err != nil {
		fmt.Printf("❌ gRPC server error: %v", err)
	}
}

// shutdownServers shuts down the gRPC server
func shutdownServers(grpcServer *grpc.Server) {
	grpcServer.GracefulStop()
	fmt.Println("✅ gRPC Server stopped gracefully")
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
