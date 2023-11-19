package main

import (
	"log"
	"net/http"

	"github.com/Samandarxon/market_app/config"
	"github.com/Samandarxon/market_app/controller"
	"github.com/Samandarxon/market_app/storage/postgres"
)

func main() {
	var cfg = config.Load()

	pgStorage, err := postgres.NewConnectionPostgres(&cfg)
	if err != nil {
		panic(err)
	}
	handler := controller.NewController(&cfg, pgStorage)

	http.HandleFunc("/category", handler.Category)
	http.HandleFunc("/client", handler.Client)
	http.HandleFunc("/product", handler.Product)
	http.HandleFunc("/branch", handler.Branch)

	log.Println("Listening:", cfg.ServiceHost+cfg.ServiceHTTPPort, "...")
	if err := http.ListenAndServe(cfg.ServiceHost+cfg.ServiceHTTPPort, nil); err != nil {
		panic("Listent and service panic:" + err.Error())
	}
}
