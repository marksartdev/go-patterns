// Package model2 Составной паттерн "Модель 2".
package model2

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/marksartdev/go-patterns/pkg/composite/mvc"
)

// Получить обработчик главной страницы.
func getMainHandler(model mvc.BeatModelInterface) gin.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("pkg/composite/model2/view.gohtml"))

	return func(ctx *gin.Context) {
		var err error

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
