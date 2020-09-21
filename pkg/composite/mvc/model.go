// Package mvc Составной паттерн "MVC".
package mvc

const startBPM = 90

// BeatModelInterface Интерфейс модели.
type BeatModelInterface interface {
	init()
	on()
	off()
	setBPM(bpm int)
	getBPM() int
	registerBeatObserver(o beatObserver)
	removeBeatObserver(o beatObserver)
	registerBPMObserver(o bpmObserver)
	removeBPMObserver(o bpmObserver)
}

// Модель.
type beatModel struct {
	generator     generator
	beatObservers []beatObserver
	bpmObservers  []bpmObserver
	bpm           int
}

// Инициализировать генератор.
func (b *beatModel) init() {
	if b.generator == nil {
		b.generator = newGenerator(b.beatEvent)
	}
}

// Включить.
func (b *beatModel) on() {
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
func (b *beatModel) registerBeatObserver(o beatObserver) {
	b.beatObservers = append(b.beatObservers, o)
}

// Удалить наблюдателя за ударами.
func (b *beatModel) removeBeatObserver(o beatObserver) {
	for i := range b.beatObservers {
		if b.beatObservers[i] == o {
			b.beatObservers = append(b.beatObservers[:i], b.beatObservers[i+1:]...)

			break
		}
	}
}

// Оповестить наблюдателей за ударами.
func (b *beatModel) notifyBeatObservers() {
	for i := range b.beatObservers {
		b.beatObservers[i].updateBeat()
	}
}

// Установить наблюдателя за изменением BPM.
func (b *beatModel) registerBPMObserver(o bpmObserver) {
	b.bpmObservers = append(b.bpmObservers, o)
}

// Удалить наблюдателя за изменением BPM.
func (b *beatModel) removeBPMObserver(o bpmObserver) {
	for i := range b.bpmObservers {
		if b.bpmObservers[i] == o {
			b.bpmObservers = append(b.bpmObservers[:i], b.bpmObservers[i+1:]...)

			break
		}
	}
}

// Оповестить наблюдателей за изменением BPM.
func (b *beatModel) notifyBPMObservers() {
	for i := range b.bpmObservers {
		b.bpmObservers[i].updateBPM()
	}
}

// Callback для генератора.
func (b *beatModel) beatEvent() {
	b.notifyBeatObservers()
}

// NewBeatModel Создать модель.
func NewBeatModel() BeatModelInterface {
	return &beatModel{}
}
