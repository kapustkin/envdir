package main

import (
	"log"
	"os"

	"github.com/kapustkin/envdir/internal"
)

func main() {
	err := internal.Run(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
