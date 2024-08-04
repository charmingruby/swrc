package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	"github.com/charmingruby/swrc/config"
	"github.com/charmingruby/swrc/internal/account"
	"github.com/charmingruby/swrc/internal/account/domain/usecase"
	"github.com/charmingruby/swrc/internal/account/infra/database/mongo_repository"
	accountInterceptor "github.com/charmingruby/swrc/internal/account/infra/transport/grpc/auth"
	"github.com/charmingruby/swrc/internal/common"
	"github.com/charmingruby/swrc/internal/common/infra/auth"
	"github.com/charmingruby/swrc/internal/common/infra/auth/interceptor"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/charmingruby/swrc/pkg/bcrypt"
	"github.com/charmingruby/swrc/pkg/jwt"
	"github.com/charmingruby/swrc/pkg/mongodb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	grpcAddr := fmt.Sprintf("%s:%s", cfg.ServerConfig.Host, cfg.ServerConfig.Port)
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal(err)
	}

	db, err := mongodb.NewMongoConnection(cfg.MongoConfig.URL, cfg.MongoConfig.Database)
	if err != nil {
		slog.Error(fmt.Sprintf("MONGO CONNECTION: %s", err.Error()))
		os.Exit(1)
	}

	jwtSvc := jwt.NewJWTService(cfg.JWTConfig.Issuer, cfg.JWTConfig.SecretKey)
	interceptor := interceptor.NewGRPCInterceptor(
		*jwtSvc,
		accountInterceptor.AccountMethodsToBypass,
		accountInterceptor.AccountRBACEnsuredMethods,
	)

	unaryInterceptor := grpc.ChainUnaryInterceptor(
		interceptor.AuthInterceptor,
		interceptor.RBACInterceptor,
	)

	server := grpc.NewServer(unaryInterceptor)
	initDependencies(server, *db, jwtSvc)
	reflection.Register(server)

	slog.Info("Starting gRPC server on port " + cfg.ServerConfig.Port + "...")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("GRPC SERVER: Failed to start server: %v", err)
	}
}

func initDependencies(server *grpc.Server, db mongo.Database, authSvc auth.TokenService) {
	accountSvc := usecase.NewAccountUseCaseRegistry(
		mongo_repository.NewAccountMongoRepository(&db),
		bcrypt.NewBcryptService(),
	)

	common.NewCommonGRPCHandlerSetup(server)
	account.NewAccountGRPCHandlerSetup(server, accountSvc, authSvc)
}
