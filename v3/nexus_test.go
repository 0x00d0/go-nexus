package v3

import (
	"fmt"
	"log"
	"testing"
)

func TestListRepository(t *testing.T) {
	nexus := NewNexus("http://192.168.229.139:8081", "admin", "password")
	ListRepository, err := nexus.ListRepositories()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(ListRepository)
}
