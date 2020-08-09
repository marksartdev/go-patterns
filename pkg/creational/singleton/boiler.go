package singleton

// ChocolateBoiler Интерфейс бойлера для нагревания шоколадно-молочной смеси.
type ChocolateBoiler interface {
	Fill() string
	Drain() string
	Boil() string
}

// Бойлер для нагревания шоколадно-молочной смеси.
type chocolateBoiler struct {
	empty  bool
	boiled bool
}

// Fill Наполнить бойлер.
func (c *chocolateBoiler) Fill() string {
	if c.empty {
		c.empty = false
		c.boiled = false

		return "Filling ..."
	}

	return ""
}

// Drain Слить смесь из бойлера.
func (c *chocolateBoiler) Drain() string {
	if !c.empty && c.boiled {
		c.empty = true

		return "Draining ..."
	}

	return ""
}

// Boil Нагреть смесь.
func (c *chocolateBoiler) Boil() string {
	if !c.empty && !c.boiled {
		c.boiled = true

		return "Boiling ..."
	}

	return ""
}
