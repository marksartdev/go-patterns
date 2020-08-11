package common

// ArrayList Интерфейс списка элементов.
type ArrayList interface {
	Size() int
	Get(i int) interface{}
	Add(item interface{})
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

// NewArrayList Создает список элементов.
func NewArrayList() ArrayList {
	return &arrayList{}
}
