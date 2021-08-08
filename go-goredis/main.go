package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {

	// 连接redis
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		// redis 服务的地址和端口号
		Addr:     "139.155.77.236:6678",
		Password: "",
		DB:       0,
	})

	/* 设置username的值为bob
	expiration 过期时间， 0代表不设置过期时间
	过期时间到达后，key删除
	 */
   // 保存数据
	err := rdb.Set(ctx, "username", "bob", 0).Err()
	if err != nil {
		// 错误信息输出，控制台或终端
		panic(err)
	}

	// 获取数据
	val, err := rdb.Get(ctx, "username").Result()
	// 判断获取数据是否失败
	if err != nil {
		panic(err)
	}
	fmt.Println("username", val)

	val2, err := rdb.Get(ctx, "key2").Result()

	// key不存在
	if err == redis.Nil {
		// 用户没有被禁用
		fmt.Println("key2 does not exists")
	}else if err != nil {
		// redis 出现异常
		panic(err)
	}else {
		// 获取到用户禁用数据
		fmt.Println("key2",val2)
	}


	defer rdb.Close()

}
