package main

import (
	"gateway-grpc/config"
	"log"
	"os"

	"gateway-grpc/cmd"
)

func main() {

	_, err := config.Setup()
	if err != nil {
		log.Fatal("Cannot load config ", err.Error())
	}

	if err := cmd.Cli().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
