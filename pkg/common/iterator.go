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
