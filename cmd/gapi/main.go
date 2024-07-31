package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	"github.com/charmingruby/swrc/config"
	accountGRPC "github.com/charmingruby/swrc/internal/account/transport/grpc"
	commonGRPC "github.com/charmingruby/swrc/internal/common/transport/grpc"

	"github.com/charmingruby/swrc/pkg/mongodb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		slog.Warn("CONFIGURATION: .env file not found")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("CONFIGURATION: %s", err.Error()))
		os.Exit(1)
	}

	_, err = mongodb.NewMongoConnection(cfg.MongoConfig.URL, cfg.MongoConfig.Database)
	if err != nil {
		slog.Error(fmt.Sprintf("MONGO CONNECTION: %s", err.Error()))
		os.Exit(1)
	}

	server := grpc.NewServer()
	commonGRPC.NewCommonGRPCHandler(server).Register()
	accountGRPC.NewAccountGRPCHandler(server).Register()
	reflection.Register(server)

	slog.Info("Starting gRPC server on port " + "9000...")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("GRPC SERVER: Failed to start server: %v", err)
	}
}
