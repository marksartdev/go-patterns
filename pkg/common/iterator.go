package common

// Iterator Интерфейс итератора.
type Iterator interface {
	HasNext() bool
	Next() interface{}
	Remove() error
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
