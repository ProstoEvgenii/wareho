package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"warehouse/config"
	"warehouse/internal/controller"
	db "warehouse/internal/materials/repository/postgres"
	"warehouse/internal/materials/usecase"
	"warehouse/migrations"
	psql "warehouse/pkg/db/postgres"
	server "warehouse/pkg/fasthttpServer"
)

func main() {
	cfg := config.ParseEnv()

	log.Log().Msg("Connection to psql")
	psqlDB, err := psql.New(cfg)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer psqlDB.Conn.Close(context.Background())

	log.Log().Msg("Run Migration...")
	err = migrations.Run(context.Background(), psqlDB)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	materialsRepo := db.NewMaterialsRepository(psqlDB)

	materialsUC := usecase.NewMaterialsUseCase(materialsRepo)

	router := controller.NewRouter(cfg, materialsUC)

	server.Start(cfg, router.Handler)
}
