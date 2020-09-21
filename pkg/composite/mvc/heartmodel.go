package mvc

import (
	"math/rand"
	"time"
)

const (
	minHeartRate     = 50
	maxHeartRate     = 120
	maxTimeForUpdate = 5
)

// HeartModelInterface Интерфейс модели данных сердцебиения.
type HeartModelInterface interface {
	init()
	getHeartRate() int
	registerBeatObserver(o beatObserver)
	removeBeatObserver(o beatObserver)
	registerBPMObserver(o bpmObserver)
	removeBPMObserver(o bpmObserver)
}

// Модель данных сердцебиения.
type heartModel struct {
	generator     generator
	beatObservers []beatObserver
	bpmObservers  []bpmObserver
	rate          int
}

// Инициализировать генератор.
func (h *heartModel) init() {
	if h.generator == nil {
		h.generator = newGenerator(h.beatEvent)
		h.generator.setTempoInBPM(minHeartRate)
		h.generator.start()

		go h.generateRate()
	}
}

// Получить частоту сердцебиения.
func (h *heartModel) getHeartRate() int {
	return h.rate
}

// Зарегистрировать наблюдателя за ударами.
func (h *heartModel) registerBeatObserver(o beatObserver) {
	h.beatObservers = append(h.beatObservers, o)
}

// Удалить наблюдателя за ударами.
func (h *heartModel) removeBeatObserver(o beatObserver) {
	for i := range h.beatObservers {
		if h.beatObservers[i] == o {
			h.beatObservers = append(h.beatObservers[:i], h.beatObservers[i+1:]...)

			break
		}
	}
}

// Оповестить наблюдателей за ударами.
func (h *heartModel) notifyBeatObservers() {
	for i := range h.beatObservers {
		h.beatObservers[i].updateBeat()
	}
}

// Установить наблюдателя за изменением BPM.
func (h *heartModel) registerBPMObserver(o bpmObserver) {
	h.bpmObservers = append(h.bpmObservers, o)
}

// Удалить наблюдателя за изменением BPM.
func (h *heartModel) removeBPMObserver(o bpmObserver) {
	for i := range h.bpmObservers {
		if h.bpmObservers[i] == o {
			h.bpmObservers = append(h.bpmObservers[:i], h.bpmObservers[i+1:]...)

			break
		}
	}
}

// Оповестить наблюдателей за изменением BPM.
func (h *heartModel) notifyBPMObservers() {
	for i := range h.bpmObservers {
		h.bpmObservers[i].updateBPM()
	}
}

// Callback для генератора.
func (h *heartModel) beatEvent() {
	h.notifyBeatObservers()
}

// Генерация псевдослучайного сердечного ритма.
func (h *heartModel) generateRate() {
	source := rand.NewSource(time.Now().Unix())
	// nolint:gosec // It's not necessary.
	r := rand.New(source)

	timer := time.NewTimer(1 * time.Second)

	for range timer.C {
		h.rate = r.Intn(maxHeartRate-minHeartRate) + minHeartRate
		h.generator.setTempoInBPM(h.rate)
		h.notifyBPMObservers()

		timer.Reset(time.Duration(r.Intn(maxTimeForUpdate)) * time.Second)
	}
}

// NewHeartModel Создать модель.
func NewHeartModel() HeartModelInterface {
	return &heartModel{}
}
