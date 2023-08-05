package csredis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type Client interface {
	R() *redis.Client
	Close()
	Set(key string, value interface{}, expire ...time.Duration) error
	Get(key string) (string, error)
}

/*
Option
Addr connection dsn
Password optional
DB select database/channel
RedisExpire set default redis expire
Timeout set default context timeout
*/
type Option struct {
	Addr        string `json:"addr"`
	Password    string `json:"password"`
	DB          int    `json:"DB"`
	RedisExpire int    `json:"redis_expire"`
	Timeout     int    `json:"timeout"`
}

type defaultClient struct {
	rdc     *redis.Client
	timeout int
	expire  time.Duration
}

func NewClient(opt *Option) Client {
	rdc := redis.NewClient(&redis.Options{
		Addr:     opt.Addr,
		Password: opt.Password,
		DB:       opt.DB,
	})

	cmd := rdc.Ping(context.Background())
	_, err := cmd.Result()
	if err != nil {
		log.Fatalf("can't connect to csredis: %v", cmd.Err())
	}

	dc := &defaultClient{rdc: rdc, timeout: opt.Timeout}

	if opt.RedisExpire != 0 {
		dc.expire = time.Duration(opt.RedisExpire) * time.Minute
	}

	return dc
}

func (c *defaultClient) R() *redis.Client {
	return c.rdc
}

func (c *defaultClient) Close() {
	c.rdc.Close()
}

func (c *defaultClient) Set(key string, value interface{}, expire ...time.Duration) error {
	ctx, cancel := c.context()
	defer cancel()
	exp := c.expire
	if len(expire) > 0 {
		exp = expire[0]
	}

	err := c.rdc.Set(ctx, key, value, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *defaultClient) Get(key string) (string, error) {
	ctx, cancel := c.context()
	defer cancel()

	val, err := c.rdc.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("")
	} else if err != nil {
		return "", err
	}

	return val, nil
}

func (c *defaultClient) context() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.TODO(), time.Duration(c.timeout)*time.Second)
}
