package main

import (
	"github.com/secmohammed/deck-poc/config"
	"github.com/secmohammed/deck-poc/container"
	"github.com/secmohammed/deck-poc/internal/router"
	log "github.com/siruspen/logrus"
	"os"
)

func main() {
	c := config.Factory(config.Type(os.Getenv("CONFIG_TYPE")))
	app := container.NewApplication(c)
	r := router.Factory(router.REST, app)
	if err := r.Expose(); err != nil {
		log.Fatalf("failed to start our http server: %s", err)
		os.Exit(1)
	}
}
