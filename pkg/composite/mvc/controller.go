package mvc

// ControllerInterface Интерфейс контроллера.
type ControllerInterface interface {
	start()
	stop()
	increaseBPM()
	decreaseBPM()
	setBPM(bpm int)
	Run()
}

// Контроллер.
type beatController struct {
	model BeatModelInterface
	view  djViewInterface
}

// Запустить.
func (b *beatController) start() {
	b.model.On()
	b.view.disableStartButton()
	b.view.enableStopButton()
}

// Остановить.
func (b *beatController) stop() {
	b.model.Off()
	b.view.disableStopButton()
	b.view.enableStartButton()
}

// Увеличить BPM на 1.
func (b *beatController) increaseBPM() {
	bpm := b.model.GetBPM()
	b.model.SetBPM(bpm + 1)
}

// Уменьшить BPM на 1.
func (b *beatController) decreaseBPM() {
	bpm := b.model.GetBPM()
	b.model.SetBPM(bpm - 1)
}

// Установить BPM.
func (b *beatController) setBPM(bpm int) {
	b.model.SetBPM(bpm)
}

// Run Запустить приложение.
func (b *beatController) Run() {
	b.view.Run()
}

// NewBeatController Создать новый контроллер.
func NewBeatController(model BeatModelInterface) ControllerInterface {
	c := &beatController{}
	c.model = model
	c.view = newDJView(c, c.model)
	c.view.init()
	c.view.disableStopButton()
	c.view.enableStartButton()
	c.model.Init()

	return c
}
