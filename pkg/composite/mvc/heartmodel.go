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
	registerBeatObserver(o BeatObserver)
	removeBeatObserver(o BeatObserver)
	registerBPMObserver(o BPMObserver)
	removeBPMObserver(o BPMObserver)
}

// Модель данных сердцебиения.
type heartModel struct {
	generator     generator
	beatObservers []BeatObserver
	bpmObservers  []BPMObserver
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
func (h *heartModel) registerBeatObserver(o BeatObserver) {
	h.beatObservers = append(h.beatObservers, o)
}

// Удалить наблюдателя за ударами.
func (h *heartModel) removeBeatObserver(o BeatObserver) {
	for i := range h.beatObservers {
		if h.beatObservers[i] == o {
			h.beatObservers[i] = nil

			break
		}
	}
}

// Оповестить наблюдателей за ударами.
func (h *heartModel) notifyBeatObservers() {
	h.collectBeatObserversGarbage()

	for i := range h.beatObservers {
		h.beatObservers[i].UpdateBeat()
	}
}

// Собрать мусор из коллекции beatObservers.
func (h *heartModel) collectBeatObserversGarbage() {
	collection := make([]BeatObserver, 0, len(h.beatObservers))

	for _, beatObserver := range h.beatObservers {
		if beatObserver != nil {
			collection = append(collection, beatObserver)
		}
	}

	h.beatObservers = collection
}

// Установить наблюдателя за изменением BPM.
func (h *heartModel) registerBPMObserver(o BPMObserver) {
	h.bpmObservers = append(h.bpmObservers, o)
}

// Удалить наблюдателя за изменением BPM.
func (h *heartModel) removeBPMObserver(o BPMObserver) {
	for i := range h.bpmObservers {
		if h.bpmObservers[i] == o {
			h.bpmObservers[i] = nil

			break
		}
	}
}

// Оповестить наблюдателей за изменением BPM.
func (h *heartModel) notifyBPMObservers() {
	h.collectBPMObserversGarbage()

	for i := range h.bpmObservers {
		h.bpmObservers[i].UpdateBPM()
	}
}

// Собрать мусор из коллекции bpmObservers.
func (h *heartModel) collectBPMObserversGarbage() {
	collection := make([]BPMObserver, 0, len(h.bpmObservers))

	for _, bpmObserver := range h.bpmObservers {
		if bpmObserver != nil {
			collection = append(collection, bpmObserver)
		}
	}

	h.bpmObservers = collection
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
