package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"ports/app"
)

func main() {
	application, err := app.Init(app.NewConfig())
	if err != nil {
		log.Fatal("could not init application: " + err.Error())
	}

	go func() {
		err := application.Start()
		if err != nil {
			log.Fatalf("could not start application: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	application.Shutdown()
	os.Exit(0)

}
