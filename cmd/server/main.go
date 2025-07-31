package main

import (
	apiHttp "api-template/internal/http"
	"api-template/internal/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	serverApp := apiHttp.New(ctx)
	defer serverApp.Shutdown()

	// Initialize the application components
	serverApp.Run()

	// Set up the server after router is initialized
	serverTimeout, err := strconv.Atoi(serverApp.Config.TIMEOUT)

	if err != nil {
		log.Panic("Invalid timeout value:", err)
	}

	timeout := time.Duration(serverTimeout) * time.Second

	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%s", serverApp.Config.SERVER_HOST, serverApp.Config.PORT),
		Handler:           serverApp.Router,
		ReadTimeout:       timeout,
		WriteTimeout:      timeout,
		ReadHeaderTimeout: timeout,
		IdleTimeout:       timeout,
	}

	utils.HandleSignals(ctx, cancel, func() {
		err := server.Shutdown(ctx)

		if err != nil {
			log.Println("Error during server shutdown:", err)
		}

		log.Println("Server gracefully stopped")
	})

	log.Printf("Server started on %s:%s", serverApp.Config.SERVER_HOST, serverApp.Config.PORT)

	finalError := server.ListenAndServe()

	if finalError != nil {
		log.Panic(finalError)
	}

	<-ctx.Done()
}
