package factory

// Абстрактная структура ингредиента.
type abstractIngredient struct {
	name string
}

func (a abstractIngredient) String() string {
	return a.name
}

// Тесто.
type dough struct {
	abstractIngredient
}

// Соус.
type sauce struct {
	abstractIngredient
}

// Сыр.
type cheese struct {
	abstractIngredient
}

// Овощи.
type veggie struct {
	abstractIngredient
}

// Пепперони.
type pepperoni struct {
	abstractIngredient
}

// Мидии.
type clams struct {
	abstractIngredient
}
