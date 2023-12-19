package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/zomgra/tracker/configs"
	"github.com/zomgra/tracker/internal/handlers"
	"github.com/zomgra/tracker/internal/service"
	web "github.com/zomgra/tracker/internal/web"
	"github.com/zomgra/tracker/pkg/db/postgres"
)

func main() {
	dir, err := os.Getwd()
	err = godotenv.Load(dir + "/configs/app.env")
	if err != nil {
		log.Fatal(err)
	}
	dbconfig := postgres.Config{ConnectionString: os.Getenv("CONNECTION_STRING")}
	appConfig := configs.ApplicationConfig{Port: os.Getenv("APPLICATION_PORT")}

	dbClient, err := postgres.NewClient(dbconfig)
	defer dbClient.Close()
	if err != nil {
		// Quick exit app, idk what doing
		dbClient.Close()
		os.Exit(1)
	}

	shipmentRepository := service.NewShipmentRepository(dbClient)
	handler := handlers.NewHandler(shipmentRepository)

	r := web.NewRoutes(handler)

	injectionErrorChan := make(chan error)

	go func() {
		shipmentRepository.InjectFromDB(injectionErrorChan)
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
