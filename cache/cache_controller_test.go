package cache

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

var db = map[string]string{
	"Tom": "630",
	"Jack": "589",
	"Sam": "567",
}

func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Fatal("callback failed")
	}
}

func TestGet(t *testing.T) {
	loadCounts := make(map[string]int, len(db))
	testGroup := NewGroup("scores", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				if _, ok := loadCounts[key]; !ok {
					loadCounts[key] = 0
				}
				loadCounts[key] += 1
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
	

		for k, v := range db {
			if view, err := testGroup.Get(k); err != nil || view.String() != v {
				t.Fatal("Failed to get value of Tom")
			} // load from callback function
			if _, err := testGroup.Get(k); err != nil || loadCounts[k] > 1 {
				t.Fatalf("cache %s miss", k)
			} // cache hit
		}

		if view, err := testGroup.Get("unknown"); err == nil {
			t.Fatalf("The value of unknow should be empty, but got %s.", view)
		}
		
}