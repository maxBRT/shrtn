package main

import (
	"fmt"
	"github.com/maxBRT/feather"
	"log"
)

func main() {
	f, err := feather.New("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Feather is running")
	if err := f.Run("8000"); err != nil {
		log.Fatal(err)
	}
}
