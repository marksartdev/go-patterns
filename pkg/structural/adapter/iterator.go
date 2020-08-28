package adapter

// Enumeration Интерфейс перечисления.
type Enumeration interface {
	HasMoreElements() bool
	NextElement() interface{}
}

// SimpleEnumeration Простое перечисление.
type SimpleEnumeration struct {
	Elements []interface{}
	index    int
}

// HasMoreElements Проверяет наличие следующего элемента.
func (s *SimpleEnumeration) HasMoreElements() bool {
	return s.index < len(s.Elements)
}

// NextElement Возвращает следующий элемент.
func (s *SimpleEnumeration) NextElement() interface{} {
	current := s.Elements[s.index]
	s.index++

	return current
}

// SimpleIterator Простой итератор.
type SimpleIterator struct {
	Elements []interface{}
	index    int
}

// HasNext Проверяет наличие следующего элемента.
func (s *SimpleIterator) HasNext() bool {
	return s.index < len(s.Elements)
}

// Next Возвращает следующий элемент.
func (s *SimpleIterator) Next() interface{} {
	current := s.Elements[s.index]
	s.index++

	return current
}

// Remove Удаляет текущий элемент.
// nolint:unparam // Implement interface.
func (s *SimpleIterator) Remove() error {
	if s.index < len(s.Elements)-1 {
		copy(s.Elements[s.index:], s.Elements[s.index+1:])
		s.Elements[len(s.Elements)-1] = struct{}{}
	}

	s.Elements = s.Elements[:len(s.Elements)-1]

	return nil
}
