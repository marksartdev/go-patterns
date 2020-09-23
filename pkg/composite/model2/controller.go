// Package model2 Составной паттерн "Модель 2".
package model2

import (
	"html/template"

	"github.com/gin-gonic/gin"

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
