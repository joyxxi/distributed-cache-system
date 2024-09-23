package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash Hash // hash function
	replicas int // For virtual nodes
	keys []int // Hash ring, sorted
	hashMap map[int]string // Virtual nodes map to real nodes
}

// New initializes a Map instance
func New(replicas int, fn Hash) *Map {
	m := &Map {
		replicas: replicas,
		hash: fn,
		hashMap: make(map[int]string),
	}

	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}

	return m
}

// Add adds some nodes(keys) to the hash
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

// Get gets the closest item in the hash to the provided key
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	// Binary search for appropriate replica
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	// Ring structure, so we use % to get the key
	return m.hashMap[m.keys[idx % len(m.keys)]]
}