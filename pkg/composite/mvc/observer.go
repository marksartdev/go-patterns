package mvc

// BeatObserver Интерфейс наблюдателя beat.
type BeatObserver interface {
	UpdateBeat()
}

// BPMObserver Интерфейс наблюдателя BPM.
type BPMObserver interface {
	UpdateBPM()
}
