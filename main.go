package main

import (
	"fmt"

	"github.com/go-redis/redis"
)
func main(){

	client := newClient()

	err := ping(client)
	if err != nil {
		fmt.Println(err)
	}
	
}
//return redis client
func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}

// ping tests connectivity for redis (PONG should be returned)
func ping(client *redis.Client) error {
	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return nil
}

