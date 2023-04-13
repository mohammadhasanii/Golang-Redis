package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	context := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	keyName := "lastname"
	value := "Ahamadi"
	err := rdb.Set(context, keyName, value, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(context, keyName).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keyName, val)

}
