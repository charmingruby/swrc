package config

import (
	env "github.com/caarlos0/env/v6"
)

type environment struct {
	ServerHost    string `env:"SERVER_HOST,required"`
	ServerPort    string `env:"SERVER_PORT,required"`
	MongoURL      string `env:"MONGO_URL,required"`
	MongoDatabase string `env:"MONGO_DB,required"`
	ClientHost    string `env:"CLIENT_HOST,required"`
	ClientPort    string `env:"CLIENT_PORT,required"`
	JWTIssuer     string `env:"JWT_ISSUER,required"`
	JWTSecretKey  string `env:"JWT_SECRET_KEY,required"`
}

func NewConfig() (*Config, error) {
	environment := environment{}
	if err := env.Parse(&environment); err != nil {
		return nil, err
	}

	cfg := Config{
		ClientConfig: &clientConfig{
			Host: environment.ClientHost,
			Port: environment.ClientPort,
		},
		ServerConfig: &serverConfig{
			Host: environment.ServerHost,
			Port: environment.ServerPort,
		},
		MongoConfig: &mongoConfig{
			URL:      environment.MongoURL,
			Database: environment.MongoDatabase,
		},
		JWTConfig: &jwtConfig{
			SecretKey: environment.JWTSecretKey,
			Issuer:    environment.JWTIssuer,
		},
	}

	return &cfg, nil
}

type Config struct {
	ClientConfig *clientConfig
	ServerConfig *serverConfig
	MongoConfig  *mongoConfig
	JWTConfig    *jwtConfig
}

type serverConfig struct {
	Host string
	Port string
}

type mongoConfig struct {
	URL      string
	Database string
}

type clientConfig struct {
	Host string
	Port string
}

type jwtConfig struct {
	SecretKey string
	Issuer    string
}
