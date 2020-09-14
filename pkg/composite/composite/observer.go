package composite

import (
	"fmt"

	"github.com/marksartdev/go-patterns/pkg/common"
)

// Интерфейс наблюдаемого объекта.
type quackObservable interface {
	RegisterObserver(observer observer)
	notifyObserver()
}

// Интерфейс наблюдателя.
type observer interface {
	update(duck quackObservable)
}

// Структура, реализующая всю необходимую логику для наблюдения.
type observable struct {
	observers common.ArrayList
	duck      quackObservable
}

// RegisterObserver Зарегистрировать наблюдателя.
func (o *observable) RegisterObserver(observer observer) {
	o.observers.Add(observer)
}

// Оповестить наблюдателей.
func (o *observable) notifyObserver() {
	iterator := o.observers.Iterator()
	for iterator.HasNext() {
		if current, ok := iterator.Next().(observer); ok {
			current.update(o.duck)
		}
	}
}

// Создать наблюдаемый объект.
func newObservable(duck quackObservable) quackObservable {
	o := &observable{}
	o.observers = common.NewArrayList()
	o.duck = duck

	return o
}

// Quackologist Наблюдатель.
type Quackologist struct {
	customWriter
}

// Обновить.
func (q *Quackologist) update(duck quackObservable) {
	q.write(fmt.Sprintf("Quackologist: %s just quacked", duck))
}
