package main

import (
	"log"

	"github.com/KonstantinGasser/datalabs/required"
)

type t struct {
	Member   []int  `required:"yes,min=6,max=12"`
	UserName string `required:"yes,min=6"`
}

func main() {
	i := new(t)
	i.Member = append(i.Member, 24)
	i.UserName = "Konstantin"
	if err := required.All(i); err != nil {
		log.Fatal(err)
	}
}
