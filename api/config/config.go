package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Env         string
	DatabaseUrl string
	Debug       bool
	GrpcPort    int
}

func NewConfig() *Config {
	debug := os.Getenv("DEBUG") != ""
	if debug {
		fmt.Println("Включен дебаг режим")
	}

	grpcPort, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		grpcPort = 50051
		fmt.Println("gRPC Port set default = 50051")
	}

	return &Config{
		Env:         os.Getenv("APP_ENV"),
		Debug:       debug,
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		GrpcPort:    grpcPort,
	}
}
