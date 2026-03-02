package main

import (
	"log"

	"github.com/subrata-bash/social/internal/env"
	"github.com/subrata-bash/social/internal/store"
)

func main() {
	cfg := config{
		Addr: env.GetString("ADDR", ":8080"),
	}
	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
