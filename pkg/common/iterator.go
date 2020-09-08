package common

// Iterator Интерфейс итератора.
type Iterator interface {
	HasNext() bool
	Next() interface{}
	Remove() error
}

// Пустой итератор.
type nullIterator struct{}

// HasNext Проверяет, есть ли еще элемент в коллекции.
func (n nullIterator) HasNext() bool {
	return false
}

// Next Возвращает следующий элемент.
func (n nullIterator) Next() interface{} {
	return nil
}

// Remove Удаляет текущий элемент.
func (n nullIterator) Remove() error {
	return UnsupportedOperationError{}
}

// NewNullIterator Создает пустой итератор.
func NewNullIterator() Iterator {
	return nullIterator{}
}

// UnsupportedOperationError Ошибка при неподдерживаемой операции.
type UnsupportedOperationError struct{}

func (UnsupportedOperationError) Error() string {
	return "This operation is not supported"
}

// IllegalStateError Ошибка при попытки удалить элемент до выполнения next().
type IllegalStateError struct{}

func (IllegalStateError) Error() string {
	return "You can't remove an item until you've done at least one next()"
}

// IllegalAccessError Ошибка при недоступной операции.
type IllegalAccessError struct{}

func (IllegalAccessError) Error() string {
	return "You don't have access to this operation"
}
