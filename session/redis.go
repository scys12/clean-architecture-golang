package session

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/scys12/clean-architecture-golang/config"
)

type RedisClient struct {
	Pool *redis.Pool
	Conn redis.Conn
}

func NewRedisPool(cfg *config.Config) (rdb *RedisClient) {
	return &RedisClient{
		Pool: &redis.Pool{
			MaxIdle:   80,
			MaxActive: 12000,
			Dial: func() (redis.Conn, error) {
				port := fmt.Sprintf("%v:%v", cfg.RedisHost, cfg.RedisPort)
				c, err := redis.Dial("tcp", port)
				if err != nil {
					return nil, err
				}
				return c, nil
			},
		},
	}
}

func Ping(c redis.Conn) error {
	_, err := redis.String(c.Do("PING"))
	if err != nil {
		return err
	}
	return nil
}
