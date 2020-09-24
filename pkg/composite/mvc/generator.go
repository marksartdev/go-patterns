package mvc

import (
	"context"
	"time"
)

// Интерфейс генератора.
type generator interface {
	start()
	stop()
	setTempoInBPM(bpm int)
}

// Генератор ритма.
type beatGenerator struct {
	duration    time.Duration
	beatEvent   func()
	cancel      context.CancelFunc
	started     bool
	resetTicker bool
}

// Запустить генератор.
func (b *beatGenerator) start() {
	if b.started {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	b.cancel = cancel
	b.started = true

	go b.do(ctx)
}

// Остановить генератор.
func (b *beatGenerator) stop() {
	b.cancel()
	b.started = false
}

// Установить bpm.
func (b *beatGenerator) setTempoInBPM(bpm int) {
	if bpm == 0 {
		b.duration = 0
		b.stop()
	} else {
		b.duration = time.Minute / time.Duration(bpm)
		b.resetTicker = true
	}
}

// Рабочий процесс.
func (b *beatGenerator) do(ctx context.Context) {
	ticker := time.NewTicker(b.duration)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			b.beatEvent()

			if b.resetTicker {
				ticker.Reset(b.duration)
				b.resetTicker = false
			}
		}
	}
}

// Создать генератор.
func newGenerator(beatEvent func()) generator {
	g := &beatGenerator{}
	g.beatEvent = beatEvent

	return g
}
