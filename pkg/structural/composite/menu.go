package composite

import "fmt"

// Меню.
type menu struct {
	menuComponent
	menuComponents []MenuComponent
	name           string
	description    string
}

// Add Добавляет дочерний компонент.
func (m *menu) Add(component MenuComponent) error {
	m.menuComponents = append(m.menuComponents, component)

	return nil
}

// Remove Удаляет дочерний компонент.
func (m *menu) Remove(component MenuComponent) error {
	for i, item := range m.menuComponents {
		if item == component {
			copy(m.menuComponents[i:], m.menuComponents[i+1:])
			m.menuComponents = m.menuComponents[:len(m.menuComponents)-1]
		}
	}

	return nil
}

// GetChild Возвращает дочерние компоненты.
func (m *menu) GetChild(i int) (MenuComponent, error) {
	return m.menuComponents[i], nil
}

// GetName Возвращает название меню.
func (m *menu) GetName() (string, error) {
	return m.name, nil
}

// GetDescription Возвращает описание меню.
func (m *menu) GetDescription() (string, error) {
	return m.description, nil
}

// Print Печатает меню.
func (m *menu) Print() error {
	msg := fmt.Sprintf("\n%s,  %s\n--------------------", m.name, m.description)

	err := m.write(msg)
	if err != nil {
		return err
	}

	for _, component := range m.menuComponents {
		err = component.Print()
		if err != nil {
			return err
		}
	}

	return nil
}

// NewMenu Создает меню.
func NewMenu(name, description string) MenuComponent {
	return &menu{newMenuComponent(), nil, name, description}
}
