// Package dproxy "Защитный заместитель" - разновидность паттерна "Заместитель".
package dproxy

// PersonBean Интерфейс анкеты.
type PersonBean interface {
	GetName() string
	GetGender() string
	GetInterests() string
	GetHotOrNotRating() int
	SetName(name string)
	SetGender(gender string)
	SetInterests(interests string)
	SetHotOrNotRating(rating int)
}

// Анкета.
type personBean struct {
	name        string
	gender      string
	interests   string
	rating      int
	ratingCount int
}

// GetName Получить имя.
func (p *personBean) GetName() string {
	return p.name
}

// GetGender Получить пол.
func (p *personBean) GetGender() string {
	return p.gender
}

// GetInterests Получить интересы.
func (p *personBean) GetInterests() string {
	return p.interests
}

// GetHotOrNotRating Получить рейтинг.
func (p *personBean) GetHotOrNotRating() int {
	if p.ratingCount == 0 {
		return 0
	}

	return p.rating / p.ratingCount
}

// SetName Задать имя.
func (p *personBean) SetName(name string) {
	p.name = name
}

// SetGender Задать пол.
func (p *personBean) SetGender(gender string) {
	p.gender = gender
}

// SetInterests Задать интересы.
func (p *personBean) SetInterests(interests string) {
	p.interests = interests
}

// SetHotOrNotRating Оценить кандидата.
func (p *personBean) SetHotOrNotRating(rating int) {
	p.rating += rating
	p.ratingCount++
}

// NewPersonBean Создать новую анкету.
func NewPersonBean() PersonBean {
	return &personBean{}
}
