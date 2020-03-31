package db

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

var RD *redis.Client

const (
	defaultRedisAddr = "localhost:6379"
)


func init (){

	addr := os.Getenv("REDIS_ADDR")
	if len(addr) <= 0 {
		addr = defaultRedisAddr
	}
	rd := redis.NewClient(&redis.Options{
		Network:            "tcp",		// default tcp conn
		Addr:               addr,
		Dialer:             nil,
		OnConnect:          nil,
		Password:           os.Getenv("REDIS_PASSWORD"),
		DB:                 0,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       10,
		MaxConnAge:         100,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})
	_,err := rd.Ping().Result()
	if err != nil {
		log.Println("redis connect fail err : ",err)
		os.Exit(1)
	}
	// log.Println("redis connect success")

	RD = rd
}
