package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/charmingruby/swrc/config"
	"github.com/charmingruby/swrc/internal/example/domain/example_usecase"
	"github.com/charmingruby/swrc/pkg/mongodb"
	"github.com/charmingruby/swrc/test/inmemory"
	"github.com/joho/godotenv"
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

	_, err = mongodb.NewMongoConnection(cfg.MongoConfig.URL, cfg.MongoConfig.Database)
	if err != nil {
		slog.Error(fmt.Sprintf("MONGO CONNECTION: %s", err.Error()))
		os.Exit(1)
	}

	initDependencies()
}

func initDependencies() {
	exampleRepo := inmemory.NewInMemoryExampleRepository()
	// if err != nil {
	// 	slog.Error(fmt.Sprintf("DATABASE REPOSITORY: %s", err.Error()))
	// 	os.Exit(1)
	// }

	_ = example_usecase.NewExampleUseCaseRegistry(exampleRepo)

}
