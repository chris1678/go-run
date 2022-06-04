/*
@Date : 2022/5/31 14:52
@Author : cirss
*/
package cache

import "time"

/**
 * @Description Set  value with key and expire time 会覆盖以前的值
 * @Param key
 * @Param val
 * @Param expire
 * @return error
 **/
func Set(key string, val interface{}, expire int) error {
	return _cache.Set(key, val, time.Duration(expire)*time.Second).Err()
}

/**
 * @Description SetNX set no exit 设置前判断是否存在或者是否过期
 * @Param key
 * @Param value
 * @Param expiration
 **/
func SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return _cache.SetNX(key, value, expiration).Result()
}

/**
 * @Description Get
 * @Param key
 * @return string
 * @return error
 **/
func Get(key string) (string, error) {
	return _cache.Get(key).Result()
}

/**
 * @Description GetRange 遍历索引区间key的值 下标可以是负数 代表反向 -1开始
 * @Param key
 * @Param start
 * @Param end
 * @return string
 * @return error
 **/
func GetRange(key string, start int64, end int64) (string, error) {
	return _cache.GetRange(key, start, end).Result()
}

/**
 * @Description GetSet 将给定 key 的值设为 value ，并返回 key 的旧值(old value)。
 * @Param key
 * @Param value
 * @return string
 * @return error
 **/
func GetSet(key string, value interface{}) (string, error) {
	return _cache.GetSet(key, value).Result()
}

/**
 * @Description MGet 批量获取数据
 * @Param keys
 * @return []interface{}
 * @return error
 **/
func MGet(keys ...string) ([]interface{}, error) {
	return _cache.MGet(keys...).Result()
}

/**
 * @Description MSet  MSet("key1", "value1", "key2", "value2")
 * @Param values
 * @return error
 **/
func MSet(values ...interface{}) error {
	return _cache.MSet(values...).Err()
}

/**
 * @Description SetRange 从某个下标开始覆盖字符串
 * @Param key
 * @Param offset
 * @Param value
 * @return error
 **/
func SetRange(key string, offset int64, value string) error {
	return _cache.SetRange(key, offset, value).Err()
}

/**
 * @Description MSetNX 批量设置前判断是否存在或者是否过期 nx not exit
 * @Param values
 * @return error
 **/
func MSetNX(values ...interface{}) error {
	return _cache.MSetNX(values...).Err()
}

/**
 * @Description Incr 自增 不存在key 则创建
 * @Param key
 * @return int64
 * @return error
 **/
func Incr(key string) (int64, error) {
	return _cache.Incr(key).Result()
}

/**
 * @Description IncrBy 按照指定的数字增
 * @Param key
 * @return int64
 * @return error
 **/
func IncrBy(key string, value int64) (int64, error) {
	return _cache.IncrBy(key, value).Result()
}

/**
 * @Description Decr 按照指定的数字减
 * @Param key
 * @return int64
 * @return error
 **/
func Decr(key string) (int64, error) {
	return _cache.Decr(key).Result()
}

/**
 * @Description DecrBy 按照指定的数字减
 * @Param key
 * @return int64
 * @return error
 **/
func DecrBy(key string, value int64) (int64, error) {
	return _cache.DecrBy(key, value).Result()
}

/**
 * @Description Append 如果key不存在则新创建
 * @Param key
 * @Param value
 * @return int64
 * @return error
 **/
func Append(key, value string) (int64, error) {
	return _cache.Append(key, value).Result()
}
