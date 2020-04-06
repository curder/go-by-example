package redis

import (
	"github.com/go-redis/redis"
)

// 使用Redis之前需要下载驱动，使用命令下载： go get -u github.com/go-redis/redis

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
}
