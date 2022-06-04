/*
@Date : 2022/5/31 14:19
@Author : cirss
*/
package cache

import "time"

/**
 * @Description Del
 * @Param key
 * @return error
 **/
func Del(key string) error {
	return _cache.Del(key).Err()
}

/**
 * @Description Dump 序列化给定 key ，并返回被序列化的值。
 * @Param key
 * @return error
 **/
func Dump(key string) (string, error) {
	return _cache.Dump(key).Result()
}

/**
 * @Description Exists  判断是否存在keys
 * @Param keys
 * @return error
 **/
func Exists(keys ...string) error {
	return _cache.Exists(keys...).Err()
}

/**
 * @Description Expire
 * @Param key
 * @Param dur
 * @return error
 **/
func Expire(key string, dur time.Duration) error {
	return _cache.Expire(key, dur).Err()
}

/**
 * @Description Keys  根据key查询 key vcode_1_*
 * @Param pattern
 * @return []string
 * @return error
 **/
func Keys(pattern string) ([]string, error) {
	return _cache.Keys(pattern).Result()
}

/**
 * @Description Move 移动key到另外一个数据库
 * @Param key
 * @Param db
 * @return error
 **/
func Move(key string, db int) error {
	return _cache.Move(key, db).Err()
}

/**
 * @Description PExpire 重新设置过期时间
 * @Param key
 * @Param expiration
 **/
func PExpire(key string, expiration time.Duration) (bool, error) {
	return _cache.PExpire(key, expiration).Result()
}

/**
 * @Description PExpire 重新设置过期时间
 * @Param key
 * @Param expiration
 **/
func PExpireAt(key string, tm time.Time) (bool, error) {
	return _cache.PExpireAt(key, tm).Result()
}

/**
 * @Description Persist 删除过期时间 永久保存
 * @Param key
 * @return bool
 * @return error
 **/
func Persist(key string) (bool, error) {
	return _cache.Persist(key).Result()
}

/**
 * @Description TTL 剩余时间
 * @Param key
 * @return time.Duration -1 没设置过期时间 -2 已过期删除 其他显示剩余时间 秒
 * @return error
 **/
func TTL(key string) (time.Duration, error) {
	return _cache.TTL(key).Result()
}

/**
 * @Description PTTL 剩余时间
 * @Param key
 * @return time.Duration -1 没设置过期时间 -2 已过期删除 其他显示剩余时间 毫秒
 * @return error
 **/
func PTTL(key string) (time.Duration, error) {
	return _cache.PTTL(key).Result()
}

/**
 * @Description RandomKey 从当前数据库中随机返回一个 key 。
 * @return string
 * @return error
 **/
func RandomKey() (string, error) {
	return _cache.RandomKey().Result()
}

/**
 * @Description Rename 重新命名key 当存在这个key 会覆盖
 * @Param key
 * @Param newkey
 * @return error
 **/
func Rename(key, newkey string) error {
	return _cache.Rename(key, newkey).Err()
}

/**
 * @Description Rename 重新命名key not exit 当不存在这个键值才能改成功
 * @Param key
 * @Param newkey
 * @return error
 **/
func RenameNX(key, newkey string) error {
	return _cache.RenameNX(key, newkey).Err()
}

/**
 * @Description Type 获取键值对应的类型
 * @Param key
 * @return string
 * @return error
 **/
func Type(key string) (string, error) {
	return _cache.Type(key).Result()
}
