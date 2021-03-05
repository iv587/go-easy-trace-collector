package main

import (
	"collector/bootstrap"
	"log"
)

func main() {
	err := bootstrap.Start()
	log.Fatalln(err)
}
