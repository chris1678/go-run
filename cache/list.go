/*
@Date : 2022/5/31 14:19
@Author : cirss
*/
package cache

//list

/**
 * @Description LPush List左插入
 * @Param key
 * @Param values
 * @return error
 **/
func LPush(key string, values ...interface{}) error {
	return _cache.LPush(key, values...).Err()
}

/**
 * @Description LPushX List 插入左侧 x exit 先判断是否存在这个key
 * @Param key
 * @Param values
 * @return error
 **/
func LPushX(key string, values ...interface{}) error {
	return _cache.LPushX(key, values...).Err()
}

/**
 * @Description RPush List右插入
 * @Param key
 * @Param values
 * @return error
 **/
func RPush(key string, values ...interface{}) error {
	return _cache.RPush(key, values...).Err()
}

/**
 * @Description RPushX  List插入右侧 x exit 先判断是否存在这个key
 * @Param key
 * @Param values
 * @return error
 **/
func RPushX(key string, values ...interface{}) error {
	return _cache.RPushX(key, values...).Err()
}

/**
 * @Description LPop List 左弹出
 * @Param key
 * @return string
 * @return error
 **/
func LPop(key string) (string, error) {
	return _cache.LPop(key).Result()
}

/**
 * @Description RPop List 右弹出
 * @Param key
 * @return string
 * @return error
 **/
func RPop(key string) (string, error) {
	return _cache.RPop(key).Result()
}

/**
 * @Description LRange list 遍历
 * @Param key
 * @Param start
 * @Param stop
 * @return []string
 * @return error
 **/
func LRange(key string, start int64, stop int64) ([]string, error) {
	return _cache.LRange(key, start, stop).Result()
}

/**
 * @Description LSet 设置某个键值
 * @Param key
 * @Param index
 * @Param value
 * @return error
 **/
func LSet(key string, index int64, value interface{}) error {
	return _cache.LSet(key, index, value).Err()
}

/**
 * @Description LIndex 获取某个键值的值
 * @Param key
 * @Param index
 * @return string
 * @return error
 **/
func LIndex(key string, index int64) (string, error) {
	return _cache.LIndex(key, index).Result()
}

/**
 * @Description LTrim 截取指定键值范围内的数据
 * @Param key
 * @Param start
 * @Param stop
 * @return string
 * @return error
 **/
func LTrim(key string, start int64, stop int64) (string, error) {
	return _cache.LTrim(key, start, stop).Result()
}

/**
 * @Description LLen 获取list的长度
 * @Param key
 * @return int64
 * @return error
 **/
func LLen(key string) (int64, error) {
	return _cache.LLen(key).Result()
}
