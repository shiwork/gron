package main

import (
	"flag"
	"fmt"
	"log"
	"time"
	"github.com/shiwork/gron/config"
)

var confFilePath string = "/etc/gron/config.json"

func main() {
	flag.StringVar(&confFilePath, "conf", confFilePath, "set config file path")
	flag.Parse()

	_, err := config.Parse(confFilePath)
	if err != nil {
		log.Fatal(err)
	}
	// check config data

	for {
		fmt.Println("loop now!")
		time.Sleep(5 * time.Second)
	}
}
