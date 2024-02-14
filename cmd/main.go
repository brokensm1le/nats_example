package main

import (
	"context"
	"fmt"
	"log"
	"nats_example/config"
	"nats_example/server/httpServer"
)

func main() {

	ctx := context.Background()

	viperInstance, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load config. Error: {%s}", err.Error())
	}

	cfg, err := config.ParseConfig(viperInstance)
	if err != nil {
		log.Fatalf("Cannot parse config. Error: {%s}", err.Error())
	}

	s := httpServer.NewServer(cfg)
	if err = s.Run(ctx); err != nil {
		fmt.Sprint(err)
	}

}
