package main

import (
	"fmt"
	"log"

	"github.com/aver343/blog/pkg/config"
	"github.com/aver343/blog/pkg/db"
	"github.com/aver343/blog/pkg/db/repository"
)

func main() {
	cfg, err := config.SetupConfig()
	if err != nil {
		log.Fatal("Error loading env vars")
	}
	dbConfig := *cfg.DbConfig
	database, err := db.New(dbConfig.Addr,
		dbConfig.MaxOpenConns,
		dbConfig.MaxIdleConns,
		dbConfig.MaxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer database.Close()
	repo := repository.NewRepository(database)
	app := NewApplication(cfg, repo)

	handler := app.Mount()
	fmt.Printf("Server is running at %s !", app.Config.Addr)
	app.Run(handler)
}
