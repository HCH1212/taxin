package dao

import (
	"context"
	"log"

	"github.com/HCH1212/taxin/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	RedisClient *redis.Client
)

func InitDB() {
	dns := config.GetConf().SQL.DSN
	// 连接pg
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func InitRedis() {
	redisConf := config.GetConf().Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisConf.Address,
		Password: redisConf.Password,
		DB:       redisConf.DB,
	})
	if _, err := RedisClient.Ping(context.Background()).Result(); err != nil {
		log.Fatal(err)
	}
}
