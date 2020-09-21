package mvc

// Контроллер ритма сердцебиения.
type heartController struct {
	model HeartModelInterface
	view  djViewInterface
}

// Запустить.
func (h *heartController) start() {}

// Остановить.
func (h *heartController) stop() {}

// Увеличить BPM на 1.
func (h *heartController) increaseBPM() {}

// Уменьшить BPM на 1.
func (h *heartController) decreaseBPM() {}

// Установить BPM.
func (h *heartController) setBPM(_ int) {}

// Run Запустить приложение.
func (h *heartController) Run() {
	h.view.Run()
}

// NewHeartController Создать новый контроллер.
func NewHeartController(model HeartModelInterface) ControllerInterface {
	c := &heartController{}
	c.model = model
	c.view = newDJView(c, newHeartAdapter(c.model))
	c.view.init()
	c.view.disableStopButton()
	c.view.disableStartButton()
	c.model.init()

	return c
}
