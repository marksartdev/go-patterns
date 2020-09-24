// Package model2 Составной паттерн "Модель 2".
package model2

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/marksartdev/go-patterns/pkg/composite/mvc"
)

// Получить обработчик главной страницы.
func getMainHandler(model mvc.BeatModelInterface) gin.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("pkg/composite/model2/view.gohtml"))

	return func(ctx *gin.Context) {
		var err error

		model.Init()

		bpm := model.GetBPM()
		if bpm == 0 {
			err = tmpl.Execute(ctx.Writer, "offline")
		} else {
			err = tmpl.Execute(ctx.Writer, bpm)
		}

		if err != nil {
			_ = ctx.Error(err)
		}
	}
}

// Получить обработчик наблюдателя за ударами.
func getBeatObserverHandler(model mvc.BeatModelInterface, upgrader *websocket.Upgrader) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			_ = ctx.Error(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)

			return
		}

		observer := getBeatObserver(ws, model)
		model.RegisterBeatObserver(observer)
	}
}

// Получить обработчик наблюдателя за изменением BPM.
func getBPMObserverHandler(model mvc.BeatModelInterface, upgrader *websocket.Upgrader) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			_ = ctx.Error(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)

			return
		}

		observer := getBPMObserver(ws, model)
		model.RegisterBPMObserver(observer)
	}
}

// Команда для представления.
type command struct {
	Elem   string `json:"elem"`
	Enable bool   `json:"enable"`
}

// Получить обработчик запроса на установку BPM.
func getSetBPMHandler(model mvc.BeatModelInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bpm := ctx.Query("bpm")

		newBPM, err := strconv.Atoi(bpm)
		if err != nil {
			_ = ctx.Error(err)
			ctx.AbortWithStatus(http.StatusInternalServerError)

			return
		}

		model.SetBPM(newBPM)
	}
}

// Получить обработчик запроса на уменьшение BPM.
func getDecreaseHandler(model mvc.BeatModelInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bpm := model.GetBPM()
		if bpm > 1 {
			bpm--
		}

		model.SetBPM(bpm)
	}
}

// Получить обработчик запроса на увеличение BPM.
func getIncreaseHandler(model mvc.BeatModelInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bpm := model.GetBPM()
		model.SetBPM(bpm + 1)
	}
}

// Получить обработчик запроса на запуск.
func getStartHandler(model mvc.BeatModelInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		model.On()

		response := []command{
			{"start", false},
			{"stop", true},
		}

		ctx.JSON(http.StatusOK, response)
	}
}

// Получить обработчик запроса на остановку.
func getStopHandler(model mvc.BeatModelInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		model.Off()

		response := []command{
			{"start", true},
			{"stop", false},
		}

		ctx.JSON(http.StatusOK, response)
	}
}
