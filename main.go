package main

import (
	"fmt"
	v3 "go-nexus3/v3"
	"log"
)

func main() {
	nexus := v3.NewNexus("http://192.168.229.139:8081", "admin", "admin")
	ListRepository, err := nexus.ListRepositories()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("%+v", ListRepository)
	for _, repository := range ListRepository {
		fmt.Printf("%+v", repository)
	}
}
