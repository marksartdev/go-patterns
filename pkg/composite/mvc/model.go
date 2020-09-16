// Package mvc Составной паттерн "MVC".
package mvc

const startBPM = 90

// Интерфейс модели.
type beatModelInterface interface {
	on()
	off()
	setBPM(bpm int)
	getBPM() int
	registerBeatObserver(o observer)
	removeBeatObserver(o observer)
	registerBPMObserver(o observer)
	removeBPMObserver(o observer)
}

// Модель.
type beatModel struct {
	generator     generator
	beatObservers []observer
	bpmObservers  []observer
	bpm           int
}

// Включить.
func (b *beatModel) on() {
	b.init()
	b.setBPM(startBPM)
	b.generator.start()
}

// Выключить.
func (b *beatModel) off() {
	b.generator.stop()
	b.setBPM(0)
}

// Установить BPM.
func (b *beatModel) setBPM(bpm int) {
	b.bpm = bpm
	b.generator.setTempoInBPM(b.getBPM())
	b.notifyBPMObservers()
}

// Получить текущий BPM.
func (b *beatModel) getBPM() int {
	return b.bpm
}

// Зарегистрировать наблюдателя за ударами.
func (b *beatModel) registerBeatObserver(o observer) {
	b.beatObservers = append(b.beatObservers, o)
}

// Удалить наблюдателя за ударами.
func (b *beatModel) removeBeatObserver(o observer) {
	for i := range b.beatObservers {
		if b.beatObservers[i] == o {
			b.beatObservers = append(b.beatObservers[:i], b.beatObservers[i+1:]...)

			break
		}
	}
}

// Оповестить наблюдателей за ударами.
func (b *beatModel) notifyBeatObservers() {
	// todo
}

// Установить наблюдателя за изменением BPM.
func (b *beatModel) registerBPMObserver(o observer) {
	b.bpmObservers = append(b.bpmObservers, o)
}

// Удалить наблюдателя за изменением BPM.
func (b *beatModel) removeBPMObserver(o observer) {
	for i := range b.bpmObservers {
		if b.bpmObservers[i] == o {
			b.bpmObservers = append(b.bpmObservers[:i], b.bpmObservers[i+1:]...)

			break
		}
	}
}

// Оповестить наблюдателей за изменением BPM.
func (b *beatModel) notifyBPMObservers() {
	// todo
}

// Инициализировать генератор.
func (b *beatModel) init() {
	if b.generator == nil {
		b.generator = newGenerator(b.beatEvent)
	}
}

// Callback для генератора.
func (b *beatModel) beatEvent() {
	b.notifyBeatObservers()
}
