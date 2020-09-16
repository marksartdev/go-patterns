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
	duration  time.Duration
	beatEvent func()
	cancel    context.CancelFunc
	started   bool
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
	b.duration = time.Minute / time.Duration(bpm)
}

// Рабочий процесс.
func (b *beatGenerator) do(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			b.beatEvent()

			time.Sleep(b.duration)
		}
	}
}

// Создать генератор.
func newGenerator(beatEvent func()) generator {
	g := &beatGenerator{}
	g.beatEvent = beatEvent

	return g
}
