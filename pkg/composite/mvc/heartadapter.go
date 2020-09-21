package mvc

// Адаптер.
type heartAdapter struct {
	heart HeartModelInterface
}

// Инициализировать модель.
func (h heartAdapter) init() {
	h.heart.init()
}

// Включить.
func (h heartAdapter) on() {}

// Выключить.
func (h heartAdapter) off() {}

// Установить BPM.
func (h heartAdapter) setBPM(_ int) {}

// Получить текущий BPM.
func (h heartAdapter) getBPM() int {
	return h.heart.getHeartRate()
}

// Зарегистрировать наблюдателя за ударами.
func (h heartAdapter) registerBeatObserver(o beatObserver) {
	h.heart.registerBeatObserver(o)
}

// Удалить наблюдателя за ударами.
func (h heartAdapter) removeBeatObserver(o beatObserver) {
	h.heart.removeBeatObserver(o)
}

// Установить наблюдателя за изменением BPM.
func (h heartAdapter) registerBPMObserver(o bpmObserver) {
	h.heart.registerBPMObserver(o)
}

// Удалить наблюдателя за изменением BPM.
func (h heartAdapter) removeBPMObserver(o bpmObserver) {
	h.heart.removeBPMObserver(o)
}

//  Создать адаптер.
func newHeartAdapter(heart HeartModelInterface) BeatModelInterface {
	return heartAdapter{heart}
}
