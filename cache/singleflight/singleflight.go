package singleflight

import (
	"sync"
)

type call struct {
	wg sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex // protexts m
	m map[string]*call
}

// If multiple goroutines attempt to call Do,
// only one will actually perform the work, while the others will wait for the result.
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	// Lock the mutex to safely access the map
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait() // Wait if there's already a call for the key
		return c.val, c.err
	}

	// Create a new call and add it to the map
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock() // Unlock after modifying the map

	c.val, c.err = fn()
	c.wg.Done()

	// Lock again to remove the key from the map once the request is complete
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
	
}