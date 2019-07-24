package main

import (
	"github.com/go-redis/redis"
)

const (
	imagesKey = "ori-tech-test-images"
)

// RedisClient is an interface which makes unit testing redis
// operations eaiser- it's to avoid having to run up a redis
// instance in that situation
type RedisClient interface {
	HGet(key, field string) *redis.StringCmd
	HExists(key, field string) *redis.BoolCmd
	HSet(key, field string, value interface{}) *redis.BoolCmd
	Ping() *redis.StatusCmd
}

// Redis holds connections and config for connecting to a redis
// instance for storing sine wave graphs
type Redis struct {
	c RedisClient
}

// NewRedis will initialise a Redis, ready for use
func NewRedis(conn string) (r Redis, err error) {
	r.c = redis.NewClient(&redis.Options{
		Addr: conn,
	})

	_, err = r.c.Ping().Result()

	return
}

// Exists returns true when a chart for a sine wave exists within redis
func (r Redis) Exists(s Sine) (bool, error) {
	return r.c.HExists(imagesKey, s.RequestID()).Result()
}

// Read returns a string representing a base64 encoded chart from redis,
// where redis contains a chart for this wave,
// or an error where the chart wasn't found
func (r Redis) Read(s Sine) (string, error) {
	return r.c.HGet(imagesKey, s.RequestID()).Result()
}

// Write will add a base64'd chart to redis
func (r Redis) Write(s Sine, c Chart) (bool, error) {
	return r.c.HSet(imagesKey, s.RequestID(), c.Base64()).Result()
}
