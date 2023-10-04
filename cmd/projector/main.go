package main

import (
	"fmt"
	"github.com/AndreiMoraru123/projector.go/pkg/projector"
	"log"
)

func main() {
	opts, err := projector.GetOpts()

	if err != nil {
		log.Fatalf("unable to get options %v", err)
	}

	fmt.Printf("opts: %+v", opts)
}
