package redisClient

import (
	"time"

	"github.com/go-redis/redis/v7"
)

type RedisStruct struct {
	r *Redis
}

func NewRedisStruct(r *Redis) RedisInterface {
	return &RedisStruct{
		r: r,
	}
}

// Set 命令用于设置给定 key 的值。如果 key 已经存储其他值， SET 就覆写旧值，且无视类型。	//SET KEY_NAME VALUE
func (r *RedisStruct) Set(key string, value interface{}, expiration time.Duration) (bool, error) {
	res, err := r.r.RE.Set(key, value, expiration).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}
	if res == "" {
		return false, nil
	}
	return true, nil
}

// SetNx 命令在指定的 key 不存在时，为 key 设置指定的值。 //SETNX KEY_NAME VALUE
func (r *RedisStruct) SetNx(key string, value interface{}, expiration time.Duration) (bool, error) {
	res, err := r.r.RE.SetNX(key, value, expiration).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}
	return res, nil
}

// Get 命令用于获取指定 key 的值。如果 key 不存在，返回 nil 。如果key 储存的值不是字符串类型，返回一个错误。	//GET KEY_NAME
func (r *RedisStruct) Get(key string) (string, error) {
	res, err := r.r.RE.Get(key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	//if res == "" {
	//	return "", nil
	//}
	return res, nil
}

// Del 命令用于删除已存在的键。不存在的 key 会被忽略。	//DEL KEY_NAME
func (r *RedisStruct) Del(key string) (int64, error) {
	res, err := r.r.RE.Del(key).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, nil
}

// Exists 命令用于检查给定 key 是否存在。	//EXISTS KEY_NAME
func (r *RedisStruct) Exists(key string) (int64, error) {
	res, err := r.r.RE.Exists(key).Result()
	if err != nil && err != redis.Nil {
		return res, err
	}
	return res, nil
}

// Incr 命令将 key 中储存的数字值增一,如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 INCR 操作 //INCR KEY_NAME
func (r *RedisStruct) Incr(key string) (int64, error) {
	res, err := r.r.RE.Incr(key).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, nil
}

//Decr decr:命令将 key 中储存的数字值减一。//DECR KEY_NAME
func (r *RedisStruct) Decr(key string) (int64, error) {
	res, err := r.r.RE.Decr(key).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, nil
}

// IncrBy 命令将 key 中储存的数字值增加任意数值
func (r *RedisStruct) IncrBy(key string, value int64) (int64, error) {
	res, err := r.r.RE.IncrBy(key, value).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, nil
}

//DecrBy decrBy:命令将 key 所储存的值减去指定的减量值。	//DECRBY KEY_NAME DECREMENT_AMOUNT
func (r *RedisStruct) DecrBy(key string, value int64) (int64, error) {
	res, err := r.r.RE.DecrBy(key, value).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, nil
}

// Expire 命令用于设置 key 的过期时间，key 过期后将不再可用。单位以秒计。//Expire KEY_NAME TIME_IN_SECONDS
func (r *RedisStruct) Expire(key string, expiration time.Duration) (bool, error) {
	res, err := r.r.RE.Expire(key, expiration).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}
	return res, nil
}

//HSet 用于为哈希表中的字段赋值,如果哈希表不存在,一个新的哈希表被创建并进行HSET操作。字段已经存在,旧值被覆盖。HSET KEY_NAME FIELD VALUE
func (r *RedisStruct) HSet(key string, field string, data interface{}) (bool, error) {
	_, err := r.r.RE.HSet(key, field, data).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}
	bo, err := r.Expire(key, time.Minute)
	if err != nil {
		return false, err
	}
	return bo, err
}

// HSetNX 命令用于为哈希表中不存在的的字段赋值 。如果字段已经存在于哈希表中，操作无效。
func (r *RedisStruct) HSetNX(key string, field string, value interface{}) (bool, error) {
	_, err := r.r.RE.HSetNX(key, field, value).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}
	bo, err := r.Expire(key, time.Minute)
	if err != nil {
		return false, err
	}
	return bo, err
}

// HDel 命令用于删除哈希表 key 中的一个或多个指定字段，不存在的字段将被忽略。
func (r *RedisStruct) HDel(key string, field ...string) (int64, error) {
	res, err := r.r.RE.HDel(key, field...).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, nil
}

// HExists 命令用于查看哈希表的指定字段是否存在。
func (r *RedisStruct) HExists(key string, field string) (bool, error) {
	res, err := r.r.RE.HExists(key, field).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}
	return res, nil
}

// HGet 命令用于返回哈希表中指定字段的值。	//HGET KEY_NAME FIELD_NAME
func (r *RedisStruct) HGet(key string, value string) (string, error) {
	res, err := r.r.RE.HGet(key, value).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return res, nil
}

// HGetAll 命令用于返回哈希表中，所有的字段和值。//HGETALL KEY_NAME
func (r *RedisStruct) HGetAll(key string) (map[string]string, error) {
	res, err := r.r.RE.HGetAll(key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return res, err
}

//Hkeys 命令用于获取哈希表中的所有域（field）
func (r *RedisStruct) Hkeys(key string) ([]string, error) {
	res, err := r.r.RE.HKeys(key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return res, err
}

// Hlen 命令用于获取哈希表中字段的数量。
func (r *RedisStruct) Hlen(key string) (int64, error) {
	res, err := r.r.RE.HLen(key).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, err
}

// HMGet 命令用于返回哈希表中，一个或多个给定字段的值。如果指定的字段不存在于哈希表，那么返回一个 nil 值。
func (r *RedisStruct) HMGet(key string, fields ...string) ([]interface{}, error) {
	res, err := r.r.RE.HMGet(key, fields...).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return res, err
}

// HVals 命令返回哈希表所有的值。
func (r *RedisStruct) HVals(key string) ([]string, error) {
	res, err := r.r.RE.HVals(key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return res, err
}

// SAdd 命令将一个或多个成员元素加入到集合中，已经存在于集合的成员元素将被忽略。//SADD KEY_NAME VALUE1..VALUEN
func (r *RedisStruct) SAdd(key string, members ...interface{}) (int64, error) {
	res, err := r.r.RE.SAdd(key, members...).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, err
}

// SRem 命令用于移除集合中的一个或多个成员元素，不存在的成员元素会被忽略。//SREM KEY MEMBER1..MEMBERN
func (r *RedisStruct) SRem(key string, members ...interface{}) (int64, error) {
	res, err := r.r.RE.SRem(key, members).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, err
}

// SMembers 命令返回集合中的所有的成员。 不存在的集合 key 被视为空集合。//SMEMBERS key
func (r *RedisStruct) SMembers(key string) ([]string, error) {
	res, err := r.r.RE.SMembers(key).Result()
	if err != nil && err != redis.Nil {
		return []string{}, err
	}
	return res, err
}

// SIsMembers 命令判断成员元素是否是集合的成员。//SISMEMBER KEY VALUE
func (r *RedisStruct) SIsMembers(key string, member string) (bool, error) {
	res, err := r.r.RE.SIsMember(key, member).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}
	return res, nil
}

// SRandMemberN 命令用于返回集合中的一个随机元素。//SRANDMEMBER KEY [count]
func (r *RedisStruct) SRandMemberN(key string, count int64) ([]string, error) {
	res, err := r.r.RE.SRandMemberN(key, count).Result()
	if err != nil && err != redis.Nil {
		return []string{}, err
	}
	return res, nil
}

// Llen 命令用于返回列表的长度	//LLEN KEY_NAME
func (r *RedisStruct) Llen(key string) (int64, error) {
	res, err := r.r.RE.LLen(key).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, nil
}

// Lpush 命令将一个或多个值插入到列表头部。 如果 key 不存在，一个空列表会被创建并执行 LPUSH 操作。
func (r *RedisStruct) Lpush(key string, value ...interface{}) (int64, error) {
	res, err := r.r.RE.LPush(key, value).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, nil
}

//LPushX 将一个值插入到已存在的列表头部，列表不存在时操作无效。
func (r *RedisStruct) LPushX(key string, value ...interface{}) (int64, error) {
	res, err := r.r.RE.LPushX(key, value).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, nil
}

// Rpush 命令用于将一个或多个值插入到列表的尾部(最右边)。//RPUSH KEY_NAME VALUE1..VALUEN
func (r *RedisStruct) Rpush(key string, value string) (int64, error) {
	res, err := r.r.RE.RPush(key, value).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, nil
}

// LRange 返回列表中指定区间内的元素，区间以偏移量 START 和 END 指定。 其中 0 表示列表的第一个元素， 1表示列表的第二个元素，以此类推。
func (r *RedisStruct) LRange(key string, start, stop int64) ([]string, error) {
	res, err := r.r.RE.LRange(key, start, stop).Result()
	if err != nil && err != redis.Nil {
		return []string{}, err
	}
	return res, nil
}

//RPop 命令用于移除列表的最后一个元素，返回值为移除的元素。	//RPOP KEY_NAME
func (r *RedisStruct) RPop(key string) (string, error) {
	res, err := r.r.RE.RPop(key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return res, nil
}

// BRpop 命令移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
func (r *RedisStruct) BRpop(timeout time.Duration, keys ...string) ([]string, error) {
	res, err := r.r.RE.BRPop(timeout, keys...).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return res, nil
}

//LPop 命令用于移除并返回列表的第一个元素。//Lpop KEY_NAME
func (r *RedisStruct) LPop(key string) (string, error) {
	res, err := r.r.RE.LPop(key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return res, nil
}

//LIndex 命令用于通过索引获取列表中的元素。你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。
func (r *RedisStruct) LIndex(key string, index int64) (string, error) {
	res, err := r.r.RE.LIndex(key, index).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return res, nil
}

func (r *RedisStruct) LInsertBefore(key string, pivot, value interface{}) (int64, error) {
	res, err := r.r.RE.LInsertBefore(key, pivot, value).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, nil
}

func (r *RedisStruct) LInsertAfter(key string, pivot, value interface{}) (int64, error) {
	res, err := r.r.RE.LInsertAfter(key, pivot, value).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, nil
}

// LRem 根据参数 COUNT 的值，移除列表中与参数 VALUE 相等的元素。count>0从头开始移除count，count<0从尾移除count，count=0，移除所有
func (r *RedisStruct) LRem(key string, count int64, value interface{}) (int64, error) {
	res, err := r.r.RE.LRem(key, count, value).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, nil
}

// LSet 通过索引来设置元素的值。 当索引参数超出范围，或对一个空列表进行 LSET 时，返回一个错误。 LSET KEY_NAME INDEX VALUE
func (r *RedisStruct) LSet(key string, index int64, value interface{}) (string, error) {
	res, err := r.r.RE.LSet(key, index, value).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return res, nil
}

//LTrim 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。
func (r *RedisStruct) LTrim(key string, start, stop int64) (string, error) {
	res, err := r.r.RE.LTrim(key, start, stop).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return res, nil
}

// Dump 命令用于序列化给定 key ，并返回被序列化的值。
func (r *RedisStruct) Dump(key string) (string, error) {
	res, err := r.r.RE.Dump(key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return res, nil
}

// Keys 命令用于查找所有符合给定模式 pattern 的 key   key*  *
func (r *RedisStruct) Keys(pattern string) ([]string, error) {
	res, err := r.r.RE.Keys(pattern).Result()
	if err != nil && err != redis.Nil {
		return []string{}, err
	}
	return res, nil
}

// Persist 命令用于移除给定 key 的过期时间，使得 key 永不过期。
func (r *RedisStruct) Persist(key string) (bool, error) {
	res, err := r.r.RE.Persist(key).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}
	return res, nil
}

// FlushDB 删除当前数据库所有 key
func (r *RedisStruct) FlushDB() (string, error) {
	res, err := r.r.RE.FlushDB().Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return res, nil
}

// MGet 命令返回所有(一个或多个)给定 key 的值。 //MGET KEY1 KEY2 .. KEYN
func (r *RedisStruct) MGet(keys ...string) ([]interface{}, error) {
	res, err := r.r.RE.MGet(keys...).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return res, nil
}

//MSet 命令用于同时设置一个或多个 key-value 对。//MSET key1 value1 key2 value2 .. keyN valueN
func (r *RedisStruct) MSet(values ...interface{}) (string, error) {
	res, err := r.r.RE.MSet(values...).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return res, nil
}

//MSetNX 命令用于所有给定 key 都不存在时，同时设置一个或多个 key-value 对。//MSETNX key1 value1 key2 value2 ..
func (r *RedisStruct) MSetNX(values ...interface{}) (bool, error) {
	res, err := r.r.RE.MSetNX(values).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}
	return res, nil
}

// RandomKey 命令从当前数据库中随机返回一个 key 。
func (r *RedisStruct) RandomKey() (string, error) {
	res, err := r.r.RE.RandomKey().Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return res, nil
}

// Rename 命令用于修改 key 的名称 。
func (r *RedisStruct) Rename(key, newkey string) (string, error) {
	res, err := r.r.RE.Rename(key, newkey).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return res, nil
}

// Type 命令用于返回 key 所储存的值的类型。
func (r *RedisStruct) Type(key string) (string, error) {
	res, err := r.r.RE.Type(key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return res, nil
}

//GetRange 命令用于获取存储在指定 key 中字符串的子字符串。字符串的截取范围由 start 和 end 两个偏移量决定(包括 start 和 end 在内)。
func (r *RedisStruct) GetRange(key string, start, end int64) (string, error) {
	res, err := r.r.RE.GetRange(key, start, end).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return res, nil
}

// Zadd 命令用于将一个或多个成员元素及其分数值加入到有序集当中。
func (r *RedisStruct) Zadd(key string, members ...*redis.Z) (int64, error) {
	res, err := r.r.RE.ZAdd(key, members...).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, nil
}

// ZRevrange Zrevrange 命令返回有序集中，指定区间内的成员。 其中成员的位置按分数值递减(从大到小)来排列。
func (r *RedisStruct) ZRevrange(key string, start, stop int64) ([]string, error) {
	res, err := r.r.RE.ZRevRange(key, start, stop).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return res, nil
}

// ZRange 命令返回有序集中，指定区间内的成员。 其中成员的位置按分数值递增(从小到大)来排列。
func (r *RedisStruct) ZRange(key string, start, stop int64) ([]string, error) {
	res, err := r.r.RE.ZRange(key, start, stop).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return res, nil
}

// ZRevRank 命令返回有序集中成员的排名。其中有序集成员按分数值递减(从大到小)排序。 排名以 0 为底，也就是说， 分数值最大的成员排名为 0 。
func (r *RedisStruct) ZRevRank(key, member string) (int64, error) {
	res, err := r.r.RE.ZRevRank(key, member).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	//这里加1，是因为返回从0开始
	return res + 1, nil
}

// ZRank 命令可以获得成员按分数值递增(从小到大)排列的排名。
func (r *RedisStruct) ZRank(key, member string) (int64, error) {
	res, err := r.r.RE.ZRank(key, member).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	count, err := r.zcard(key)
	if err != nil && err != redis.Nil {
		return int64(0), err
	}

	return count - res, nil
}

//Zcard 命令用于计算集合中元素的数量。
func (r *RedisStruct) zcard(key string) (int64, error) {
	res, err := r.r.RE.ZCard(key).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, nil
}

// GetSet 命令用于设置指定 key 的值，并返回 key 的旧值。
func (r *RedisStruct) GetSet(key string, value interface{}) (string, error) {
	res, err := r.r.RE.GetSet(key, value).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return res, nil
}

// Append 命令用于为指定的 key 追加值。
//如果 key 已经存在并且是一个字符串， APPEND 命令将 value 追加到 key 原来的值的末尾。
//如果 key 不存在， APPEND 就简单地将给定 key 设为 value ，就像执行 SET key value 一样。
func (r *RedisStruct) Append(key string, value string) (int64, error) {
	res, err := r.r.RE.Append(key, value).Result()
	if err != nil && err != redis.Nil {
		return int64(0), err
	}
	return res, nil
}
