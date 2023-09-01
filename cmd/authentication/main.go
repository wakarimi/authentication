package main

import (
	"authentication/api"
	_ "authentication/docs"
	"authentication/internal/config"
	"authentication/internal/database"
	"log"
)

// @title Wakarimi Authentication API
// @version 0.0
// @description This is the authentication service for Wakarimi.
// @contact.name Zalimannard
// @contact.email zalimannard@mail.ru
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8020
// @BasePath /api
func main() {
	cfg, err := config.LoadConfiguration()
	if err != nil {
		log.Fatal("Failed to load configuration: %v", err)
	}

	log.Println(cfg.DatabaseConnectionString)
	db, err := database.ConnectDb(cfg.DatabaseConnectionString)
	if err != nil {
		log.Fatal("Failed to connect to the database: %v", err)
	}
	defer db.Close()
	database.SetDatabase(db)

	err = database.RunMigrations(db, "./internal/database/migrations")
	if err != nil {
		log.Fatal("Failed to run migrations: %v", err)
	}

	r := api.SetupRouter(cfg)
	r.Run(":" + cfg.Port)
}
