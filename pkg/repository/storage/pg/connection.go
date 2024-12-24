package postgresql

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"log"

	"github.com/urizennnn/instashop/internal/config"
	"github.com/urizennnn/instashop/pkg/repository/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	lg "gorm.io/gorm/logger"
)

func ConnectToDatabase(configDatabases config.Database) *gorm.DB {
	dbsCV := configDatabases
	connectedDB := connectToDb(dbsCV.DB_HOST, dbsCV.USERNAME, dbsCV.PASSWORD, dbsCV.DB_NAME, dbsCV.DB_PORT, dbsCV.SSLMODE, dbsCV.TIMEZONE)

	println(dbsCV.DB_NAME, "database connected")
	storage.DB.Postgresql = connectedDB

	return connectedDB
}

func connectToDb(host, user, password, dbname, port, sslmode, timezone string) *gorm.DB {
	if _, err := strconv.Atoi(port); err != nil {
		u, err := url.Parse(port)
		if err != nil {
			panic(err)
		}

		detectedPort := u.Port()
		if detectedPort == "" {
			panic("Port is required")
		}
		port = detectedPort
	}
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", host, user, password, dbname, port, sslmode, timezone)

	newLogger := lg.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		lg.Config{
			LogLevel:                  lg.Error, // Log level
			IgnoreRecordNotFoundError: true,     // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)

	}

	log.Println("Connected to database")
	// db = db.Debug() //database debug mode
	return db
}
