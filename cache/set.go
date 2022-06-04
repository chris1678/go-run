/*
@Date : 2022/5/31 15:32
@Author : cirss
集合成员是唯一的，这就意味着集合中不能出现重复的数据。
list是可以重复的
*/
package cache

/**
 * @Description SAdd 向集合添加一个或多个成员
 * @Param key
 * @Param members
 * @return error
 **/
func SAdd(key string, members ...interface{}) error {
	return _cache.SAdd(key, members...).Err()
}

/**
 * @Description SCard 获取集合的成员数
 * @Param key
 * @return int64
 * @return error
 **/
func SCard(key string) (int64, error) {
	return _cache.SCard(key).Result()
}
