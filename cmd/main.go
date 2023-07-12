package main

import (
	"fmt"

	"shopping-cart/internal/config"
	"shopping-cart/internal/core/services/cartservice"
	"shopping-cart/internal/handlers/carthandler"
	"shopping-cart/internal/repositories/cartrepository"
	"shopping-cart/internal/repositories/itemrepository"

	"shopping-cart/internal/infra/server"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		fmt.Printf("failed to load env configs, running default values. Error:%v", err)
	}

	cartRepository := cartrepository.New()
	itemrepository := itemrepository.New()

	cartService := cartservice.New(cartRepository, itemrepository)

	cartHandler := carthandler.New(cartService)

	srv := server.NewServer(cfg.Server)
	srv.NewRouter(cartHandler)

	err = srv.Run()
	if err != nil {
		panic(err)
	}
}
