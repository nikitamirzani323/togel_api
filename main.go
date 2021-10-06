package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"bitbucket.org/isbtotogroup/api_go/db"
	"bitbucket.org/isbtotogroup/api_go/routes"
)

func main() {
	db.Init()
	app := routes.Init()
	// log.Fatal(app.Listen(":7071"))

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":7071"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	log.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	log.Println("Fiber was successful shutdown.")
}
