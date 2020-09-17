package mvc

import "log"

// Интерфейс контроллера.
type controllerInterface interface {
	start()
	stop()
	increaseBPM()
	decreaseBPM()
	setBPM(bpm int)
}

// Контроллер.
type beatController struct {
}

// Запустить.
func (b *beatController) start() {
	log.Println("Start!")
}

// Остановить.
func (b *beatController) stop() {
	log.Println("Stop")
}

// Увеличить BPM на 1.
func (b *beatController) increaseBPM() {
	log.Println("Increase")
}

// Уменьшить BPM на 1.
func (b *beatController) decreaseBPM() {
	log.Println("Decrease")
}

// Установить BPM.
func (b *beatController) setBPM(bpm int) {
	log.Printf("Set BPM on %d\n", bpm)
}

// Создать новый контроллер.
func newBeatController() controllerInterface {
	return &beatController{}
}
