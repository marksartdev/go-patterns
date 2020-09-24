package model2

import (
	"strconv"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	"github.com/marksartdev/go-patterns/pkg/composite/mvc"
)

// Наблюдатель.
// nolint:structcheck // Используется в дочерних структурах.
type observer struct {
	ws             *websocket.Conn
	model          mvc.BeatModelInterface
	removeObserver func()
}

// Отправить сообщение.
func (o *observer) send(msg string) {
	w, err := o.ws.NextWriter(websocket.TextMessage)
	if err != nil {
		o.removeObserver()

		err = o.ws.Close()
		if err != nil {
			log.Error(err)
		}

		return
	}

	_, err = w.Write([]byte(msg))
	if err != nil {
		log.Error(err)
	}

	err = w.Close()
	if err != nil {
		log.Error(err)
	}
}

// Наблюдатель за ударами.
type beatObserver struct {
	observer
}

// UpdateBeat Отправить удар.
func (b *beatObserver) UpdateBeat() {
	b.send("")
}

// Создать наблюдателя за ударами.
func getBeatObserver(ws *websocket.Conn, model mvc.BeatModelInterface) mvc.BeatObserver {
	b := &beatObserver{}
	b.ws = ws
	b.model = model
	b.removeObserver = func() {
		b.model.RemoveBeatObserver(b)
	}

	return b
}

// Наблюдатель за изменением BPM.
type bpmObserver struct {
	observer
}

// UpdateBPM Отправить новый BPM.
func (b *bpmObserver) UpdateBPM() {
	bpm := b.model.GetBPM()
	if bpm == 0 {
		b.send("offline")
	} else {
		b.send(strconv.Itoa(bpm))
	}
}

// Создать наблюдателя за изменением BPM.
func getBPMObserver(ws *websocket.Conn, model mvc.BeatModelInterface) mvc.BPMObserver {
	b := &bpmObserver{}
	b.ws = ws
	b.model = model
	b.removeObserver = func() {
		b.model.RemoveBPMObserver(b)
	}

	return b
}
