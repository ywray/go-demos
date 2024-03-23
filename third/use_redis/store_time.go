package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

func ExampleStoreTime() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	t1 := time.Now().Unix()
	err = client.Set("test_time", t1, 3*time.Second).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("001")
	time.Sleep(5 * time.Second)
	//读取数据
	t11, err := client.Get("test_time").Int64()
	if err != nil {
		// 如果返回的错误是key不存在
		if errors.Is(err, redis.Nil) {
			fmt.Println("use_redis key not exist")
		} else {
			// 出其他错了
			return
		}
	}
	t2 := time.Now().Unix()
	sub := t2 - t11
	fmt.Printf("sub:%v, t1:%v, t11:%v, t2:%v", sub, t1, t11, t2)

}
