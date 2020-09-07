package vproxy

// Заместитель табло курса валют.
type proxyExchanger struct {
	exchanger Exchanger
	state     state
	initState state
	mainState state
}

// Show Отобразить текущий курс валют.
func (e *proxyExchanger) Show() {
	e.state.show()
}

// GetRates Получить курс валют.
func (e *proxyExchanger) GetRates() map[string]float64 {
	return e.state.getRates()
}

// Загрузить настоящее табло курса валют.
func (e *proxyExchanger) initExchanger() {
	e.exchanger = newExchanger()
	e.state = e.mainState
	e.Show()
}

// NewExchanger Создать табло курса валют из заместителя.
func NewExchanger() Exchanger {
	proxy := &proxyExchanger{}
	proxy.initState = initStage{proxy}
	proxy.mainState = mainState{proxy}
	proxy.state = proxy.initState

	return proxy
}
