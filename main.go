package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

func main() {

	var port string
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	client := newClient()

	err := ping(client)
	if err != nil {
		fmt.Println(err)
	}
	err = set(client)
	if err != nil {
		fmt.Println(err)
	}

	err = get(client)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Starting Test HTTP Server.........")
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}

}

//return redis client

func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:      "master.cf-8477718e-2180-43f.vgi9g4.usw2.cache.amazonaws.com:6379",
		Password:  "196e76b243b8cf8b1748c5142c2df0e3a1fed8448a13061f7d6fbf4eec2692a6", // no password set - d1ea6f732e97324b6ae7abea9580d76de8cefb288cec6eca3d36d3cee995ab7e
		DB:        0,                                                                  // use default DB
		TLSConfig: &tls.Config{},                                                      //for at transit encryption
	})

	return client
}

// ping tests connectivity for redis (PONG should be returned)
func ping(client *redis.Client) error {
	fmt.Println("executing Ping Pong....")
	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>
	fmt.Println("PING PONG operation is done")
	return nil
}

// set executes the redis Set command
func set(client *redis.Client) error {
	err := client.Set("key", "Rajendra", 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func get(client *redis.Client) error {
	val, err := client.Get("key").Result()
	if err != nil {
		return (err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

	return nil
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}


