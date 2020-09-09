package dproxy

import (
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

type proxyPersonBean struct {
	handler invocationHandler
}

// GetName Получить имя.
func (p proxyPersonBean) GetName() string {
	resp, err := p.handler.invoke(p.getFuncName())
	if errHandle(err) {
		return ""
	}

	return resp.String()
}

// GetGender Получить пол.
func (p proxyPersonBean) GetGender() string {
	resp, err := p.handler.invoke(p.getFuncName())
	if errHandle(err) {
		return ""
	}

	return resp.String()
}

// GetInterests Получить интересы.
func (p proxyPersonBean) GetInterests() string {
	resp, err := p.handler.invoke(p.getFuncName())
	if errHandle(err) {
		return ""
	}

	return resp.String()
}

// GetHotOrNotRating Получить рейтинг.
func (p proxyPersonBean) GetHotOrNotRating() int {
	resp, err := p.handler.invoke(p.getFuncName())
	if errHandle(err) {
		return 0
	}

	return int(resp.Int())
}

// SetName Задать имя.
func (p proxyPersonBean) SetName(name string) {
	_, err := p.handler.invoke(p.getFuncName(), name)
	errHandle(err)
}

// SetGender Задать пол.
func (p proxyPersonBean) SetGender(gender string) {
	_, err := p.handler.invoke(p.getFuncName(), gender)
	errHandle(err)
}

// SetInterests Задать интересы.
func (p proxyPersonBean) SetInterests(interests string) {
	_, err := p.handler.invoke(p.getFuncName(), interests)
	errHandle(err)
}

// SetHotOrNotRating Оценить кандидата.
func (p proxyPersonBean) SetHotOrNotRating(rating int) {
	_, err := p.handler.invoke(p.getFuncName(), rating)
	errHandle(err)
}

func (p proxyPersonBean) getFuncName() string {
	callers := make([]uintptr, 15)
	// nolint:gomnd // Пропускаем в стеке вызовы в данном методе
	n := runtime.Callers(2, callers)
	frames := runtime.CallersFrames(callers[:n])
	frame, _ := frames.Next()

	fullFuncName := strings.Split(frame.Function, ".")

	return fullFuncName[len(fullFuncName)-1]
}

// NewOwnerProxy Создать заместителя для владельца анкеты.
func NewOwnerProxy(person PersonBean) PersonBean {
	handler := newOwnerInvocationHandler(person)

	return proxyPersonBean{handler}
}

// NewNonOwnerProxy Создать заместителя для гостя анкеты.
func NewNonOwnerProxy(person PersonBean) PersonBean {
	handler := newNonOwnerInvocationHandler(person)

	return proxyPersonBean{handler}
}

func errHandle(err error) bool {
	if err != nil {
		log.Error(err)

		return true
	}

	return false
}
