package main

import (
	"authentication/internal/config"
	"authentication/internal/database"
	"log"
)

func main() {
	cfg, err := config.LoadConfiguration()
	if err != nil {
		log.Fatal("Failed to load configuration: %v", err)
	}

	log.Println(cfg.DatabaseConnectionString)
	db, err := database.ConnectDb(cfg.DatabaseConfiguration.DatabaseConnectionString)
	if err != nil {
		log.Fatal("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	err = database.RunMigrations(db, "./internal/database/migrations")
	if err != nil {
		log.Fatal("Failed to run migrations: %v", err)
	}
}
