package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"ppp/config"
	"ppp/entity"
	"ppp/pb"
	"ppp/service"
)

type App struct {
	config *config.Config
}

func NewApp(config *config.Config) *App {
	return &App{config: config}
}

func (app *App) Start() {
	lsn, err := net.Listen("tcp", fmt.Sprintf(":%d", app.config.GrpcPort))
	if err != nil {
		panic(err)
	}

	opts := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)

	db, err := sqlx.Connect("pgxpool", app.config.DatabaseUrl)
	if err != nil {
		fmt.Println("Unable to connect to database: \n", err)
		panic(err)
	}

	warehouse := service.NewWarehouse(db)
	user := entity.NewUser("manager", &entity.TypeUser{})

	pb.RegisterProductServiceServer(grpcServer, service.NewProductService(warehouse, user))

	log.Println("starting server")

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lsn); err != nil {
		panic(err)
	}
}
