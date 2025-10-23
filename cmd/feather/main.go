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
	if err := f.RunTLS("443",
		"/etc/letsencrypt/live/shrtn.it.com/fullchain.pem",
		"/etc/letsencrypt/live/shrtn.it.com/privkey.pem"); err != nil {
		log.Fatal(err)
	}
}
