package main

import (
	"fmt"

	"shopping-cart/internal/config"
	"shopping-cart/internal/core/services/cartservice"
	"shopping-cart/internal/handlers/carthandler"
	"shopping-cart/internal/infra/server"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		fmt.Printf("failed to load env configs, running default values. Error:%v", err)
	}

	cartService := cartservice.New(nil)

	cartHandler := carthandler.New(cartService)

	srv := server.NewServer(cfg.Server)
	srv.NewRouter(cartHandler)

	err = srv.Run()
	if err != nil {
		panic(err)
	}
}
