package di

import (
	"context"
	"crypto/tls"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewRedis(v *viper.Viper, logger *zap.Logger) *redis.Client {
	v.SetDefault("redis_address", "localhost:6379")
	v.SetDefault("redis_db", 0)
	v.SetDefault("redis_username", "")
	v.SetDefault("redis_password", "")
	v.SetDefault("redis_tls_enabled", false)

	address := v.GetString("redis_address")
	db := v.GetInt("redis_db")
	username := v.GetString("redis_username")
	password := v.GetString("redis_password")
	tlsEnabled := v.GetBool("redis_tls_enabled")

	opts := redis.Options{
		Addr:     address,
		DB:       db,
		Username: username,
		Password: password,
	}
	if tlsEnabled {
		opts.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	client := redis.NewClient(&opts)
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Sugar().Fatalf("failed to connect to redis: %v, addr: %s db: %d", err, address, db)
	}
	logger.Info("connected to redis", zap.String("addr", address), zap.Int("db", db))
	return client

}