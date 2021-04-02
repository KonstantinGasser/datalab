package storage

import (
	"context"
	"fmt"
	"sync"
)

type inMem struct {
	sync.Mutex
	// store is a map holding key:value pairs
	store map[string]interface{}
}

func (store *inMem) Put(ctx context.Context, key string, value interface{}) error {
	// lock the store for the time of the operation
	store.Lock()
	// unlock store after operation is done
	defer store.Unlock()
	// key:value already exists abort insert!
	if _, ok := store.store[key]; ok {
		return fmt.Errorf("could not add key: %v - already in store", key)
	}
	// assign value to key in store
	store.store[key] = value
	return nil
}

func (store *inMem) Get(ctx context.Context, key string) (interface{}, error) {
	// lock the store for the time of the operation
	store.Lock()
	// unlock store after operation is done
	defer store.Unlock()
	if _, ok := store.store[key]; !ok {
		return nil, fmt.Errorf("could not find value for key:%v", key)
	}
	return store.store[key], nil
}

// Exists checks if a key:value exists in the in-memory map
func (store *inMem) Exists(ctx context.Context, key string) bool {
	// lock the store for the time of the operation
	store.Lock()
	// unlock store after operation is done
	defer store.Unlock()
	_, ok := store.store[key]
	return ok
}
