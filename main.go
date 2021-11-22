package main

import (
	"log"
	"github.com/manuelnelson7/twittit/handlers"
	"github.com/manuelnelson7/twittit/bd"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("No connection")
		return
	}
	handlers.Manejadores()
}
