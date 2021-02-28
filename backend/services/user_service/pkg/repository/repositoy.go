package repository

// import (
// 	"fmt"
// 	"log"
// )

// // Repository defines methods to interact with the storage layer
// type Repository interface {
// 	Insert(data map[string]i)
// }

// // New returns a Repository interface of the given _type (mongo,..)
// func New(_type, connString string) (Repository, error) {
// 	switch _type {
// 	case "mongo":
// 		return newMongoClient(connString)
// 	default:
// 		log.Printf("could not find a Repository for: %s\n", _type)
// 		return nil, fmt.Errorf("could not find a Repository for: %s", _type)
// 	}
// }
