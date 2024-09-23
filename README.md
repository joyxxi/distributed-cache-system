# distributed-cache-system

Distributed cache system modeled after groupcache

## Project Structure

distributed-cache-system/
|--cache
|--consistenthash
|--|--consistenthash.go // Consisten hashing implementation
|--lru/
|--|--lru.go // LRU cache implementation
|--byteview.go // Encapsulation and abstraction of cache values
|--cache.go // Core caching logic, hendling concurrent cache operations.
|--cache_controller.go // Controller for the flow of the caching system.
|--http.go // Provide HTTP-based access to the cache.
