package mvc

// Интерфейс наблюдателя beat.
type beatObserver interface {
	updateBeat()
}

// Интерфейс наблюдателя BPM.
type bpmObserver interface {
	updateBPM()
}
