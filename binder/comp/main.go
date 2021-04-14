package main

import (
	"log"

	"github.com/KonstantinGasser/datalabs/binder"
)

type test struct {
	Field string `bind:"yes" json:"field"`
}

func main() {
	s := test{
		Field: "hello",
	}

	if err := binder.MustBind(&s); err != nil {
		log.Fatal(err)
	}
}
