package main

import (
	"log"

	"github.com/minoritea/sns/rpc/servers"
)

func main() {
	err := servers.Up()
	if err != nil {
		log.Println(err)
	}
}
