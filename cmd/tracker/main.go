package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/zomgra/tracker/internal/db/postgres"
	"github.com/zomgra/tracker/internal/shipment"
	"github.com/zomgra/tracker/internal/web"
	"github.com/zomgra/tracker/pkg/config"
)

func main() {

	// Initial logger
	log.SetFormatter(&log.JSONFormatter{})

	// Initial config
	appConfig := config.ApplicationConfig{}
	config.MustSetConfig(&appConfig)
	dbClient, err := postgres.NewClient(*appConfig.DbConfig)
	if err != nil {
		// Quick exit app, idk what doing
		log.Fatalf("error with connecting to db: %s", err.Error())
	}
	defer dbClient.Close()

	shipmentRepository := shipment.NewRepository(dbClient)
	shipmentService := shipment.NewService(shipmentRepository)
	handler := shipment.NewHandler(shipmentService)

	r := web.NewRoutes(handler)

	injectionErrorChan := make(chan error)

	go func() {
		shipmentService.InjectFromDB(injectionErrorChan)
	}()

	go func() {
		log.Printf("Server is running on %s", appConfig.Port)
		if err := http.ListenAndServe(":"+appConfig.Port, r); err != nil {
			if err != http.ErrServerClosed {
				log.Fatalf("Error starting server: %v", err)
			}
		}
	}()

	sigchan := createSignal()
	select {
	case err := <-injectionErrorChan:
		if err != nil {
			// Quick exit app, idk what doing
			dbClient.Close()
			os.Exit(1)
		}
	case <-sigchan:
		dbClient.Close()
		os.Exit(1)
	}
}

func createSignal() chan os.Signal {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	return sigchan
}
