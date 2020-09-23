// Package mvc Составной паттерн "MVC".
package mvc

const startBPM = 90

// BeatModelInterface Интерфейс модели.
type BeatModelInterface interface {
	Init()
	On()
	Off()
	SetBPM(bpm int)
	GetBPM() int
	RegisterBeatObserver(o beatObserver)
	RemoveBeatObserver(o beatObserver)
	RegisterBPMObserver(o bpmObserver)
	RemoveBPMObserver(o bpmObserver)
}

// Модель.
type beatModel struct {
	generator     generator
	beatObservers []beatObserver
	bpmObservers  []bpmObserver
	bpm           int
}

// Init Инициализировать генератор.
func (b *beatModel) Init() {
	if b.generator == nil {
		b.generator = newGenerator(b.beatEvent)
	}
}

// On Включить.
func (b *beatModel) On() {
	b.SetBPM(startBPM)
	b.generator.start()
}

// Off Выключить.
func (b *beatModel) Off() {
	b.generator.stop()
	b.SetBPM(0)
}

// SetBPM Установить BPM.
func (b *beatModel) SetBPM(bpm int) {
	b.bpm = bpm
	b.generator.setTempoInBPM(b.GetBPM())
	b.notifyBPMObservers()
}

// GetBPM Получить текущий BPM.
func (b *beatModel) GetBPM() int {
	return b.bpm
}

// RegisterBeatObserver Зарегистрировать наблюдателя за ударами.
func (b *beatModel) RegisterBeatObserver(o beatObserver) {
	b.beatObservers = append(b.beatObservers, o)
}

// RemoveBeatObserver Удалить наблюдателя за ударами.
func (b *beatModel) RemoveBeatObserver(o beatObserver) {
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

// RegisterBPMObserver Установить наблюдателя за изменением BPM.
func (b *beatModel) RegisterBPMObserver(o bpmObserver) {
	b.bpmObservers = append(b.bpmObservers, o)
}

// RemoveBPMObserver Удалить наблюдателя за изменением BPM.
func (b *beatModel) RemoveBPMObserver(o bpmObserver) {
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
