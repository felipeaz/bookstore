package redis

import (
	"fmt"
	"log"

	"bookstore/internal/app/logger"
	"github.com/garyburd/redigo/redis"
)

const (
	ClosingConnectionErrorMessage = "Failed to close redis connection %s"
)

// Cache implements Redis functions
type Cache struct {
	Host   string
	Port   string
	Expire string
	Logger logger.LogInterface
}

// NewCache returns an instance of Cache
func NewCache(host, port, expire string, logger logger.LogInterface) (*Cache, error) {
	cache := &Cache{
		Host:   host,
		Port:   port,
		Expire: expire,
		Logger: logger,
	}

	_, err := cache.connect()
	return cache, err
}

// Connect initialize the cache.
func (c *Cache) connect() (redis.Conn, error) {
	h := fmt.Sprintf("%s:%s", c.Host, c.Port)

	conn, err := redis.Dial("tcp", h)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return conn, nil
}

// Set inserts a value into the Cache
func (c *Cache) Set(key string, value []byte) error {
	conn, err := c.connect()
	if err != nil {
		c.Logger.Error(err)
		return err
	}
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			c.Logger.Warn(fmt.Sprintf(ClosingConnectionErrorMessage, err.Error()))
		}
	}(conn)

	_, err = conn.Do("SET", key, value)
	if err != nil {
		c.Logger.Error(err)
		return err
	}

	_, err = conn.Do("EXPIRE", key, c.Expire)
	if err != nil {
		c.Logger.Error(err)
		return err
	}

	return nil
}

// Get returns a value from Cache
func (c *Cache) Get(key string) ([]byte, error) {
	conn, err := c.connect()
	if err != nil {
		c.Logger.Error(err)
		return nil, err
	}
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			c.Logger.Warn(fmt.Sprintf(ClosingConnectionErrorMessage, err.Error()))
		}
	}(conn)

	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			return nil, nil
		}
		c.Logger.Error(err)
		return nil, err
	}

	return data, nil
}

// Flush removes a value from Cache
func (c *Cache) Flush(key string) error {
	conn, err := c.connect()
	if err != nil {
		c.Logger.Error(err)
		return err
	}
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			c.Logger.Warn(fmt.Sprintf(ClosingConnectionErrorMessage, err.Error()))
		}
	}(conn)

	_, err = conn.Do("DEL", key)
	return err
}
