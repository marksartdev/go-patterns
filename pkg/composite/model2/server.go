package model2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	"github.com/marksartdev/go-patterns/pkg/composite/mvc"
)

// StartServer Запустить сервер.
func StartServer(model mvc.BeatModelInterface, addr string) {
	upgrader := &websocket.Upgrader{}

	router := gin.Default()
	router.GET("", getMainHandler(model))

	ws := router.Group("ws")
	ws.GET("beat", getBeatObserverHandler(model, upgrader))
	ws.GET("bpm", getBPMObserverHandler(model, upgrader))

	api := router.Group("api")
	api.GET("set", getSetBPMHandler(model))
	api.GET("decrease", getDecreaseHandler(model))
	api.GET("increase", getIncreaseHandler(model))
	api.GET("start", getStartHandler(model))
	api.GET("stop", getStopHandler(model))

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
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
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
