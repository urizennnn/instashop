package redis

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/urizennnn/instashop/internal/config"
	"github.com/urizennnn/instashop/pkg/repository/storage"
	"github.com/urizennnn/instashop/utility"
	"github.com/go-redis/redis/v8"
)

var (
	Ctx     = context.Background()
	KeyName = "study-sync"
)

func ConnectToRedis(logger *utility.Logger, configDatabases config.Redis) *redis.Client {
	dbsCV := configDatabases
	utility.LogAndPrint(logger, "connecting to redis server")
	connectedServer := connectToDb(dbsCV.REDIS_HOST, dbsCV.REDIS_PORT, dbsCV.REDIS_DB, logger)
if connectedServer == nil {
		utility.LogAndPrint(logger, "failed to connect to redis server")
		panic("failed to connect to redis server")
	}
	utility.LogAndPrint(logger, "connected to redis server")

	storage.DB.Redis = connectedServer

	return connectedServer
}

func connectToDb(host, port, db string, logger *utility.Logger) *redis.Client {
	if _, err := strconv.Atoi(port); err != nil {
		u, err := url.Parse(port)
		if err != nil {
			utility.LogAndPrint(logger, fmt.Sprintf("parsing url %v to get port failed with: %v", port, err))
			panic(err)
		}

		detectedPort := u.Port()
		if detectedPort == "" {
			utility.LogAndPrint(logger, fmt.Sprintf("detecting port from url %v failed with: %v", port, err))
			return nil
		}
		port = detectedPort
	}
	dbInst, err := strconv.Atoi(db)
	if err != nil {
		utility.LogAndPrint(logger, fmt.Sprintf("Failed to parse db from url %v failed with %v ", db, err))
		panic(err)
	}

	addr := fmt.Sprintf("%v:%v", host, port)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       dbInst,
	})

	if err := redisClient.Ping(Ctx).Err(); err != nil {
		utility.LogAndPrint(logger, fmt.Sprintln(addr))
		utility.LogAndPrint(logger, fmt.Sprintln("Redis db error: ", err))
	}

	pong, _ := redisClient.Ping(Ctx).Result()
	utility.LogAndPrint(logger, fmt.Sprintln("Redis says: ", pong))

	return redisClient
}
