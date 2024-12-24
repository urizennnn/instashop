package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/urizennnn/instashop/internal/config"
	"github.com/urizennnn/instashop/internal/models/migrations"
	"github.com/urizennnn/instashop/internal/models/migrations/seed"
	"github.com/urizennnn/instashop/pkg/repository/storage"
	postgresql "github.com/urizennnn/instashop/pkg/repository/storage/pg"
	"github.com/urizennnn/instashop/pkg/repository/storage/redis"
	"github.com/urizennnn/instashop/pkg/router"
	"github.com/urizennnn/instashop/utility"
)

func main() {
	logger := utility.NewLogger()
	if logger == nil {
		log.Fatal("Error creating logger")
	}

	var name string
	flag := os.Getenv("ENV")
	if flag == "false" || flag == "" {
		name = "./app"
	} else {
		name = "./local"
	}
	configuration := config.Setup(logger, name)

	validatorRef := validator.New()

	postgresql.ConnectToDatabase(configuration.Database)

	redis.ConnectToRedis(logger, configuration.Redis)

	db := storage.Connection()

	if configuration.Database.Migrate {
		logger.Info("Running all migrations and seeding roles")
		migrations.RunAllMigrations(db)
		seed.SeedRoles(logger, db.Postgresql)
	}

	fmt.Println("Server is running on port: ", configuration.Server.Port)

	r := router.Setup(logger, validatorRef, db, &configuration.App)
	if err := r.Run("0.0.0.0:" + configuration.Server.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
