package main

import (
	"log"

	"github.com/lebrancconvas/ShortenURL/config"
	"github.com/lebrancconvas/ShortenURL/db"
	"github.com/lebrancconvas/ShortenURL/models"
	"github.com/lebrancconvas/ShortenURL/controllers"
	"github.com/lebrancconvas/ShortenURL/server"
)

func main() {
	cfg := config.Load()

	port := cfg.ServerPort()
	dburl := cfg.LoadDBURL()

	db := db.NewDatabase()
	if err := db.InitDB(dburl); err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	model := models.NewModel(db.GetDB())
	controller := controllers.NewController(model)
	router := server.NewRouter(controller)
	srv := server.NewServer(router, port)

	srv.Run()

}
