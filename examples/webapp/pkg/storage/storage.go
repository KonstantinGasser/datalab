package storage

import (
	"context"
	"log"
)

// Storage represents types which can add and get data from a given data store
type Storage interface {
	Put(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (interface{}, error)
	Exists(ctx context.Context, key string) bool
}

func New(_type string) Storage {
	switch _type {
	case "in-memory":
		return &inMem{
			store: make(map[string]interface{}),
		}
	default:
		log.Fatalf("could not find: %v as storage type\n", _type)
		return nil
	}
}
