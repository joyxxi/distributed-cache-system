package main

/*
Test commands:
$ curl http://localhost:9999/_cache/scores/Tom
630

$ curl http://localhost:9999/_cache/scores/kkk
kkk not exist
*/

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joyxxi/distributed-cache-system/cache"
)

// Database source
var db = map[string]string {
	"Tom": "630",
	"Jack": "589",
	"Sam": "567",
}

func main() {
	cache.NewGroup("scores", 2<<10, cache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := cache.NewHTTPPool(addr)
	log.Println("cache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}