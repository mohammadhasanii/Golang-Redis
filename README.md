
# Golang-Redis

This package was developed for redis by the Go language




![Logo](https://blog.petehouston.com/wp-content/uploads/2022/06/blog.petehouston.com-use-redis-to-cache-data-in-go.jpg)


## Usage/Examples

```go
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

```


## Run Application

Install my-project with npm

```bash
  go run app.go
```
    
