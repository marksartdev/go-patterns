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
	RegisterBeatObserver(o BeatObserver)
	RemoveBeatObserver(o BeatObserver)
	RegisterBPMObserver(o BPMObserver)
	RemoveBPMObserver(o BPMObserver)
}

// Модель.
type beatModel struct {
	generator     generator
	beatObservers []BeatObserver
	bpmObservers  []BPMObserver
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
func (b *beatModel) RegisterBeatObserver(o BeatObserver) {
	b.beatObservers = append(b.beatObservers, o)
}

// RemoveBeatObserver Удалить наблюдателя за ударами.
func (b *beatModel) RemoveBeatObserver(o BeatObserver) {
	for i := range b.beatObservers {
		if b.beatObservers[i] == o {
			b.beatObservers[i] = nil

			break
		}
	}
}

// Оповестить наблюдателей за ударами.
func (b *beatModel) notifyBeatObservers() {
	b.collectBeatObserversGarbage()

	for i := range b.beatObservers {
		b.beatObservers[i].UpdateBeat()
	}
}

// Собрать мусор из коллекции beatObservers.
func (b *beatModel) collectBeatObserversGarbage() {
	collection := make([]BeatObserver, 0, len(b.beatObservers))

	for _, beatObserver := range b.beatObservers {
		if beatObserver != nil {
			collection = append(collection, beatObserver)
		}
	}

	b.beatObservers = collection
}

// RegisterBPMObserver Установить наблюдателя за изменением BPM.
func (b *beatModel) RegisterBPMObserver(o BPMObserver) {
	b.bpmObservers = append(b.bpmObservers, o)
}

// RemoveBPMObserver Удалить наблюдателя за изменением BPM.
func (b *beatModel) RemoveBPMObserver(o BPMObserver) {
	for i := range b.bpmObservers {
		if b.bpmObservers[i] == o {
			b.bpmObservers[i] = nil

			break
		}
	}
}

// Оповестить наблюдателей за изменением BPM.
func (b *beatModel) notifyBPMObservers() {
	b.collectBPMObserversGarbage()

	for i := range b.bpmObservers {
		b.bpmObservers[i].UpdateBPM()
	}
}

// Собрать мусор из коллекции bpmObservers.
func (b *beatModel) collectBPMObserversGarbage() {
	collection := make([]BPMObserver, 0, len(b.bpmObservers))

	for _, bpmObserver := range b.bpmObservers {
		if bpmObserver != nil {
			collection = append(collection, bpmObserver)
		}
	}

	b.bpmObservers = collection
}

// Callback для генератора.
func (b *beatModel) beatEvent() {
	b.notifyBeatObservers()
}

// NewBeatModel Создать модель.
func NewBeatModel() BeatModelInterface {
	return &beatModel{}
}
