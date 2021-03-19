package main

import (
	"fmt"

	"github.com/KonstantinGasser/clickstream/protobuf/test/test"
)

func main() {
	t := test.Test{
		Name: &test.Name{
			First: "Konstantin",
			Last:  "Gasser",
		},
	}
	fmt.Println(t)
}
