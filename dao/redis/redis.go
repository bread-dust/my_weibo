/*
@author:Deng.l.w
@version:1.20
@date:2023-03-04 0:03
@file:redis.go
*/

package redis

import (
	"settings"
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize, //连接池大小
	})

	_, err = rdb.Ping().Result() //ping redis 判断连接成功
	return
}

func Close() {
	_ = rdb.Close()

}
