package mvc

// Адаптер.
type heartAdapter struct {
	heart HeartModelInterface
}

// Init Инициализировать модель.
func (h heartAdapter) Init() {
	h.heart.init()
}

// On Включить.
func (h heartAdapter) On() {}

// Off Выключить.
func (h heartAdapter) Off() {}

// SetBPM Установить BPM.
func (h heartAdapter) SetBPM(_ int) {}

// GetBPM Получить текущий BPM.
func (h heartAdapter) GetBPM() int {
	return h.heart.getHeartRate()
}

// RegisterBeatObserver Зарегистрировать наблюдателя за ударами.
func (h heartAdapter) RegisterBeatObserver(o beatObserver) {
	h.heart.registerBeatObserver(o)
}

// RemoveBeatObserver Удалить наблюдателя за ударами.
func (h heartAdapter) RemoveBeatObserver(o beatObserver) {
	h.heart.removeBeatObserver(o)
}

// RegisterBPMObserver Установить наблюдателя за изменением BPM.
func (h heartAdapter) RegisterBPMObserver(o bpmObserver) {
	h.heart.registerBPMObserver(o)
}

// RemoveBPMObserver Удалить наблюдателя за изменением BPM.
func (h heartAdapter) RemoveBPMObserver(o bpmObserver) {
	h.heart.removeBPMObserver(o)
}

//  Создать адаптер.
func newHeartAdapter(heart HeartModelInterface) BeatModelInterface {
	return heartAdapter{heart}
}
