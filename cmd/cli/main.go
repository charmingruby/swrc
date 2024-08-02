package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/charmingruby/swrc/config"
	"github.com/charmingruby/swrc/internal/common/infra/cli"
	"github.com/charmingruby/swrc/internal/common/infra/transport/grpc/client"
	"github.com/charmingruby/swrc/proto/pb"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("CONFIGURATION: .env file not found")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("CONFIGURATION: %s", err.Error()))
		os.Exit(1)
	}

	grpcServerAddr := fmt.Sprintf("%s:%s", cfg.ServerConfig.Host, cfg.ServerConfig.Port)
	conn, err := grpc.NewClient(grpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != err {
		slog.Error(fmt.Sprintf("GRPC SERVER CONNECTION: %s", err.Error()))
		os.Exit(1)
	}

	// health check
	grpcClient := client.NewCommonClientHandler(conn)
	_, err = grpcClient.HealthCheck(
		context.Background(),
		&pb.PingMessage{Greeting: "health check"})
	if err != nil {
		slog.Error(fmt.Sprintf("GRPC SERVER HEALTH CHECK: %s", err.Error()))
		os.Exit(1)
	}

	var rootCommand = &cobra.Command{}
	cli := cli.NewCLI(rootCommand, grpcClient)
	cli.Register()
	cli.Start()
}
