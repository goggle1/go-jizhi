package db

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
)

var RedisConnPool *redis.Pool

func InitRedis(conntime, readtime, writetime, maxidle int, addr, password string, db int) {
	connTimeout := time.Duration(conntime) * time.Millisecond
	readTimeout := time.Duration(readtime) * time.Millisecond
	writeTimeout := time.Duration(writetime) * time.Millisecond

	RedisConnPool = &redis.Pool{
		MaxIdle:     maxidle,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr,
				redis.DialConnectTimeout(connTimeout),
				redis.DialReadTimeout(readTimeout),
				redis.DialWriteTimeout(writeTimeout),
				redis.DialDatabase(db),
				redis.DialPassword(password))

			if err != nil {
				logrus.WithField("redisConfig.Addr", addr).Errorln("open redis fail:", err)
				return nil, err
			}

			return c, err
		},
		TestOnBorrow: PingRedis,
	}
	logrus.Info("redis db run")

}

func PingRedis(c redis.Conn, t time.Time) error {
	_, err := c.Do("ping")
	if err != nil {
		logrus.Error("[ERROR] ping redis fail", err)
	}
	return err
}
