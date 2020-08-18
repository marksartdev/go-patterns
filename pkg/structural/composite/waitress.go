package composite

// Waitress Интерфейс официантки.
type Waitress interface {
	PrintMenu() error
}

// Официантка.
type waitress struct {
	allMenu MenuComponent
}

// PrintMenu Печатает меню.
func (w waitress) PrintMenu() error {
	return w.allMenu.Print()
}

// NewWaitress Создает официантку.
func NewWaitress(allMenu MenuComponent) Waitress {
	return waitress{allMenu}
}
