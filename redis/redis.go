package redis

import (
	"github.com/go-redis/redis"
	"gopackage/loger"
)

var Redis *redis.Client

func NewRedisConnect(address string, auth string, dbIndex int)(*redis.Client){

	Redis = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: auth,
		DB:       dbIndex,
	})
	_, err := Redis.Ping().Result()
	if err != nil {
		panic("Redis Client SetUp Failed!" + err.Error())
		return nil
	}

	loger.Info("# Redis Connect Success.")
	return Redis
}