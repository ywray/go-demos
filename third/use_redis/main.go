package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

func ExampleClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("feekey", "examples", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("feekey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("feekey", val)

	val2, err := client.Get("feekey2").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("feekey does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("feekey", val2)
	}
}

func incr() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	var lockKey = "counter_lock"
	var counterKey = "counter"

	// lock
	resp := client.SetNX(lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()

	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result:", lockSuccess)
		return
	}

	// counter ++
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil || errors.Is(err, redis.Nil) {
		cntValue++
		response := client.Set(counterKey, cntValue, 0)
		_, err = response.Result()
		if err != nil {
			// log err
			println("set value error!")
			return
		}
	}
	println("current counter is", cntValue)

	delResp := client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}
}

func redisSetNXTest() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}

func main() {
	//ExampleClient()
	//redisSetNXTest()
	//ExampleJson()
	ExampleStoreTime()
}
