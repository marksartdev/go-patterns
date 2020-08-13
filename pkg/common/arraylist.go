package common

// ArrayList Интерфейс списка элементов.
type ArrayList interface {
	Size() int
	Get(i int) interface{}
	Add(item interface{})
	Remove(i int)
	Iterator() Iterator
}

// Список элементов.
type arrayList struct {
	items []interface{}
}

// Size Возвращает количество элементов списка.
func (a *arrayList) Size() int {
	return len(a.items)
}

// Get Возвращает элемент списка.
func (a *arrayList) Get(i int) interface{} {
	return a.items[i]
}

// Add Добавляет элемент ы список.
func (a *arrayList) Add(item interface{}) {
	a.items = append(a.items, item)
}

// Remove Удаляет элемент списка.
func (a *arrayList) Remove(i int) {
	copy(a.items[i:], a.items[i+1:])
	a.items = a.items[:a.Size()-1]
}

// Iterator Создает итератор списка элементов.
func (a *arrayList) Iterator() Iterator {
	cp := &arrayList{}
	cp.items = make([]interface{}, a.Size())
	copy(cp.items, a.items)

	return &arrayListIterator{cp, 0}
}

// NewArrayList Создает список элементов.
func NewArrayList() ArrayList {
	return &arrayList{}
}

// Итератор списка элементов.
type arrayListIterator struct {
	arrayList ArrayList
	position  int
}

// HasNext Проверяет, есть ли еще элемент в коллекции.
func (a *arrayListIterator) HasNext() bool {
	return a.position < a.arrayList.Size()
}

// Next Возвращает следующий элемент.
func (a *arrayListIterator) Next() interface{} {
	item := a.arrayList.Get(a.position)
	a.position++

	return item
}

// Remove Удаляет текущий элемент.
func (a *arrayListIterator) Remove() error {
	if a.position <= 0 {
		return IllegalStateError{}
	}

	if a.position-1 < a.arrayList.Size() {
		a.arrayList.Remove(a.position - 1)
	}

	return nil
}
