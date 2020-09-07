package vproxy

import "fmt"

type state interface {
	show()
	getRates() map[string]float64
}

type initStage struct {
	proxy *proxyExchanger
}

func (i initStage) show() {
	fmt.Println("Loading...")

	go i.proxy.initExchanger()
}

func (i initStage) getRates() map[string]float64 {
	return nil
}

type mainState struct {
	proxy *proxyExchanger
}

func (m mainState) show() {
	m.proxy.exchanger.Show()
}

func (m mainState) getRates() map[string]float64 {
	return m.proxy.exchanger.GetRates()
}
