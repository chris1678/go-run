/*
@Date : 2022/5/30 10:06
@Author : cirss
*/
package cache

import (
	"fmt"
	"github.com/chris1678/go-run/config"
	"github.com/go-redis/redis/v7"
)

var _cache *redis.Client

func Initialize() {
	//初始化redis
	c := config.RedisConfig
	_cache = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Addr, c.Port),
		Password: c.Password, // no password set
		DB:       0,          // use default DB
	})
	//设置验证码缓存

}

/**
 * @Description GetClient 暴露原生client
 * @return *redis.Client
 **/
func GetClient() *redis.Client {
	return _cache
}
