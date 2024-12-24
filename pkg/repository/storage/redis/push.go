package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/urizennnn/instashop/utility"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func PushtoRedis(ctx context.Context, key string, value interface{}, expiry time.Duration, client *redis.Client, logger *utility.Logger) error {
	err := client.Set(ctx, key, value, expiry).Err()
	if err != nil {
		logger.Error("Failed to push to cache")
		utility.LogAndPrint(logger, fmt.Sprintf("Code failed with %v", err))
		return err
	}
	logger.Info(fmt.Sprintf("Pushed key '%s' to Redis with expiry %s", key, expiry))
	return nil
}

func GetfromRedis(ctx context.Context, key string, client *redis.Client, logger *utility.Logger) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		logger.Info(fmt.Sprintf("Key '%s' does not exist in Redis", key))
		return "", nil
	} else if err != nil {
		logger.Error("Failed to get from cache")
		utility.LogAndPrint(logger, fmt.Sprintf("Code failed with %v", err))
		return "", err
	}

	logger.Info(fmt.Sprintf("Retrieved key '%s' from Redis", key))
	return val, nil
}
