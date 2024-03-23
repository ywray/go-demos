package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type fileManager struct {
	// 记录写入字符总数（中文算一个）
	TotalWriteCount int `json:"TotalWriteCount"`
	// 记录上个文件写入了多少字符（中文算一个）
	LastWriteCount int `json:"LastWriteCount"`
	// 记录已经产生了多少文件
	NexSuffix int `json:"NexSuffix"`
}

func ExampleJson() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	TestStruct := fileManager{
		TotalWriteCount: 1,
		LastWriteCount:  22,
		NexSuffix:       333,
	}
	//json序列化
	datas, _ := json.Marshal(TestStruct)
	//缓存数据
	err = client.Set("filename", datas, 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("001")

	//读取数据
	val, err := client.Get("filename").Result()
	if err != nil {
		panic(err)
	}
	//json反序列化
	object := &fileManager{}
	err = json.Unmarshal([]byte(val), object)
	if err != nil {
		panic(err)
	}
	fmt.Println("002", object.TotalWriteCount, object.LastWriteCount, object.NexSuffix)

	// 删除缓存
	client.Del("filename")

	//读取数据
	val, err = client.Get("filename").Result()
	if err != nil {
		panic(err)
	}
	//json反序列化
	object = &fileManager{}
	err = json.Unmarshal([]byte(val), object)
	if err != nil {
		panic(err)
	}
	fmt.Println("003", object.TotalWriteCount, object.LastWriteCount, object.NexSuffix)

}
