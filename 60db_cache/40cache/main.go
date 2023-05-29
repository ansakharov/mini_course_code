package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrNotPositiveTTL = errors.New("ttl must be >= 0")

type item struct {
	value      any
	validUntil int64
}

type Cache struct {
	storage map[string]*item
	mu      sync.RWMutex
}

// Решение с TTL предусматривает очистку кэша посредством удаления старых записей по истечению времени их жизни
func New(ctx context.Context, cleanPeriod time.Duration) (*Cache, error) {
	if cleanPeriod == 0 {
		return nil, errors.New("cleanPeriod must be > 0")
	}
	cache := &Cache{
		storage: make(map[string]*item),
	}

	go cache.cleanOldCache(ctx, cleanPeriod)
	return cache, nil
}

func (c *Cache) Set(key string, value any, ttl time.Duration) error {
	if ttl <= 0 {
		return ErrNotPositiveTTL
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	expirationTimestamp := time.Now().Add(ttl).UnixNano()
	c.storage[key] = &item{
		value:      value,
		validUntil: expirationTimestamp,
	}

	return nil
}

func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	// из-за возможного удаления ключа нельзя использовать defer c.mu.RUnlock()

	item, ok := c.storage[key]
	if !ok {
		c.mu.RUnlock()
		return nil
	}

	if item.validUntil < time.Now().UnixNano() {
		c.mu.RUnlock()
		//
		c.mu.Lock()
		delete(c.storage, key)
		c.mu.Unlock()
		return nil
	}

	c.mu.RUnlock()
	return item.value
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.storage, key)
}

func (c *Cache) cleanOldCache(ctx context.Context, cleanPeriod time.Duration) {
	ticker := time.NewTicker(cleanPeriod)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			now := time.Now().UnixNano()
			// здесь есть пространство для улучшения:
			// можно хранить не один кэш, а бакет со многими кэшами, шардируя ключи в них
			// тогда мы не будем блокировать ВСЮ мапу на очистку,
			// а переменно будем чистить много маленьких мап параллельно
			c.mu.Lock()
			for key, item := range c.storage {
				if item.validUntil < now {
					delete(c.storage, key)
				}
			}
			c.mu.Unlock()
		}
	}
}

func main() {
	mainRace()
}

func mainRace() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	cache, err := New(ctx, time.Second)
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}

	fmt.Println(cache.Set("name", "Alex", time.Minute))
	fmt.Println(cache.Set("hobby", "BJJ", time.Second))
	fmt.Println(cache.Set("hobby3", "BJJ", time.Second*10))
	fmt.Println(cache.Set("hobby2", "BJJ", time.Second*20))
	fmt.Println(cache.Set("hobby4", "BJJ", time.Second*15))
	fmt.Println(cache.Set("name", "Alex", time.Second*6))

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(cache.Get("name"))
		fmt.Println(cache.Get("hobby"))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cache.Delete("hobby")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(cache.Get("name"))
		fmt.Println(cache.Get("hobby"))
	}()

	wg.Wait()

	for {
		cache.mu.Lock()
		for k, v := range cache.storage {
			fmt.Println(k, *v)
		}
		cache.mu.Unlock()
		time.Sleep(time.Second)
	}
}
