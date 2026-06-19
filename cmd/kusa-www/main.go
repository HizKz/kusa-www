package main

import (
	"fmt"
	"log"

	"github.com/HizKz/kusa-www/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

}
