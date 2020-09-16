package session

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/scys12/clean-architecture-golang/pkg/config"
)

type redisClient struct {
	pool *redis.Pool
	conn redis.Conn
}

func NewRedisPool(cfg *config.Config) SessionStore {
	redisPool := &redis.Pool{
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
	}
	return &redisClient{
		pool: redisPool,
		conn: redisPool.Get(),
	}
}

func Ping(s SessionStore) error {
	_, err := redis.String(s.Connect().Do("PING"))
	if err != nil {
		return err
	}
	return nil
}
