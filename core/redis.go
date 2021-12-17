package core

import (
	"context"
	"fmt"
	"motor-mini-backend/config"
	"runtime"

	"motor-mini-backend/core/logger"
	log "motor-mini-backend/core/logger"
	redisHook "motor-mini-backend/core/redis"

	"github.com/go-redis/redis/v8"
)

// Redis client instance to use redis
var Redis *redis.Client

//ConnectRedis connect to redis
func ConnectRedis() {
	redis.SetLogger(&redisLogger{})
	Redis = redis.NewClient(&redis.Options{
		Addr: config.Redis.Addr,
		DB:   config.Redis.DB, // use default DB
	})
	Redis.AddHook(&redisHook.LogHook{})
}

type redisLogger struct {
}

func (l *redisLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(3)
	message := fmt.Sprintf(format, v...)
	logData := log.CustomLog{
		File:    file,
		Line:    line,
		Message: message,
	}
	logger.Warn(ctx, "redis", logData)
}
