package main

import (
	"fmt"
	"log"

	"github.com/Ttibsi/d-roll/src"
)

func main() {
	fmt.Println("Listening on port 3000")
	log.Default()
	src.Serve()
}
