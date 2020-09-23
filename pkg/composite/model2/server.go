package model2

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// StartServer Запустить сервер.
func StartServer(addr string) {
	router := gin.Default()
	// todo Set main handler

	api := router.Group("api")
	// todo Set API handlers

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go startServer(srv)
	go stopServer(wg, srv)

	wg.Wait()
}

// Запустить сервер.
func startServer(srv *http.Server) {
	log.Infof("Starting server on %s...", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// Остановить сервер.
func stopServer(wg *sync.WaitGroup, srv *http.Server) {
	defer wg.Done()
	defer log.Info("Server is stopped")

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	fmt.Println()
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Warning("Server forced to shutdown")
	}
}
