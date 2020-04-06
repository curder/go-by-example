package redis

import (
	"github.com/go-redis/redis"
	"testing"
)

// 设置和获取Redis
func TestSetAndGetString(t *testing.T) {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		t.Logf("set score failed, err:%v\n", err)
		return
	}

	val, err := rdb.Get("score").Result()
	if err != nil {
		t.Logf("get score failed, err:%v\n", err)
		return
	}
	t.Log("score", val)

	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		t.Log("name does not exist")
	} else if err != nil {
		t.Logf("get name failed, err:%v\n", err)
		return
	} else {
		t.Log("name", val2)
	}
}

// 复杂数据结构的操作
func TestZsetAndZget(t *testing.T) {
	zsetKey := "language_rank"
	languages := []*redis.Z{
		&redis.Z{Score: 90.0, Member: "Golang"},
		&redis.Z{Score: 98.0, Member: "Java"},
		&redis.Z{Score: 95.0, Member: "Python"},
		&redis.Z{Score: 97.0, Member: "JavaScript"},
		&redis.Z{Score: 99.0, Member: "C/C++"},
	}
	// ZADD
	num, err := rdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		t.Logf("zadd failed, err:%v\n", err)
		return
	}
	t.Logf("zadd %d succ.\n", num)

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		t.Logf("zincrby failed, err:%v\n", err)
		return
	}
	t.Logf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		t.Logf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		t.Log(z.Member, z.Score)
	}

	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		t.Logf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		t.Log(z.Member, z.Score)
	}
}
