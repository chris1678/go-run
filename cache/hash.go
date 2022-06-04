/*
@Date : 2022/5/31 14:53
@Author : cirss
*/
package cache

/**
 * @Description HashGet
 * @Param hk
 * @Param key
 * @return string
 * @return error
 **/
func HGet(hk, key string) (string, error) {
	return _cache.HGet(hk, key).Result()
}

/**
 * @Description HMGet 获取所有给定字段的值
 * @Param key
 * @Param fields
 * @return []interface{}
 * @return error
 **/
func HMGet(key string, fields ...string) ([]interface{}, error) {
	return _cache.HMGet(key, fields...).Result()
}

/**
 * @Description HMSet 同时将多个 field-value (域-值)对设置到哈希表 key 中。
 * @Param key
 * @Param values
 * @return bool
 * @return error
 **/
func HMSet(key string, values ...interface{}) (bool, error) {
	return _cache.HMSet(key, values...).Result()
}

/**
 * @Description HSet 将哈希表 key 中的字段 field 的值设为 value 。
 * @Param key
 * @Param values
 * @return error
 **/
func HSet(key string, values ...interface{}) error {
	return _cache.HSet(key, values...).Err()
}

/**
 * @Description HSetNX 只有在字段 field 不存在时，设置哈希表字段的值。
 * @Param key
 * @Param field
 * @Param value
 * @return error
 **/
func HSetNX(key string, field string, value interface{}) error {
	return _cache.HSetNX(key, field, value).Err()
}

/**
 * @Description HExists 查看哈希表 key 中，指定的字段是否存在。
 * @Param key
 * @Param field
 * @return bool
 * @return error
 **/
func HExists(key string, field string) (bool, error) {
	return _cache.HExists(key, field).Result()
}

/**
 * @Description HGetAll 获取在哈希表中指定 key 的所有字段和值
 * @Param key
 * @return map[string]string
 * @return error
 **/
func HGetAll(key string) (map[string]string, error) {
	return _cache.HGetAll(key).Result()
}

/**
 * @Description HIncrBy 为哈希表 key 中的指定字段的整数值加上增量 increment 。
 * @Param key
 * @Param field
 * @Param incr
 * @return error
 **/
func HIncrBy(key string, field string, incr int64) error {
	return _cache.HIncrBy(key, field, incr).Err()
}

/**
 * @Description HIncrByFloat 为哈希表 key 中的指定字段的浮点数值加上增量 increment 。
 * @Param key
 * @Param field
 * @Param incr
 * @return error
 **/
func HIncrByFloat(key string, field string, incr float64) error {
	return _cache.HIncrByFloat(key, field, incr).Err()
}

/**
 * @Description HKeys 获取所有哈希表中的字段
 * @Param key
 * @return []string
 * @return error
 **/
func HKeys(key string) ([]string, error) {
	return _cache.HKeys(key).Result()
}

/**
 * @Description HLen	获取哈希表中字段的数量
 * @Param key
 * @return int64
 * @return error
 **/
func HLen(key string) (int64, error) {
	return _cache.HLen(key).Result()
}

/**
 * @Description HVals 获取哈希表中所有值。
 * @Param key
 * @return []string
 * @return error
 **/
func HVals(key string) ([]string, error) {
	return _cache.HVals(key).Result()
}

/**
 * @Description HDel
 * @Param hk
 * @Param key
 * @return error
 **/
func HDel(hk, key string) error {
	return _cache.HDel(hk, key).Err()
}
