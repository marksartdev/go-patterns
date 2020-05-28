package factory

// Изготовить тонкое тесто.
func newThinCrustDough() dough {
	return dough{
		abstractIngredient{
			name: "Thin Crust Dough",
		},
	}
}

// Изготовить толстое тесто.
func newThickCrustDough() dough {
	return dough{
		abstractIngredient{
			name: "Thick Crust Dough",
		},
	}
}

// Изготовить соус "Маринара".
func newMarinaraSauce() sauce {
	return sauce{
		abstractIngredient{
			name: "Marinara Sauce",
		},
	}
}

// Изготовить томатный соус.
func newPlumTomatoSauce() sauce {
	return sauce{
		abstractIngredient{
			name: "Plum Tomato Sauce",
		},
	}
}

// Изготовить сыр "Реджиано".
func newReggianoCheese() cheese {
	return cheese{
		abstractIngredient{
			name: "Grated Reggiano Cheese",
		},
	}
}

// Изготовить сыр "Моцарелла".
func newMozzarellaCheese() cheese {
	return cheese{
		abstractIngredient{
			name: "Shredded Mozzarella Cheese",
		},
	}
}

// Изготовить сыр "Пармезан".
func newParmesanCheese() cheese {
	return cheese{
		abstractIngredient{
			name: "Parmesan Cheese",
		},
	}
}

// Изготовить чеснок.
func newGarlic() veggie {
	return veggie{
		abstractIngredient{
			name: "garlic",
		},
	}
}

// Изготовить лук.
func newOnion() veggie {
	return veggie{
		abstractIngredient{
			name: "onion",
		},
	}
}

// Изготовить грибы.
func newMushroom() veggie {
	return veggie{
		abstractIngredient{
			name: "mushroom",
		},
	}
}

// Изготовить красный перец.
func newRedPepper() veggie {
	return veggie{
		abstractIngredient{
			name: "red pepper",
		},
	}
}

// Изготовить орегано.
func newOregano() veggie {
	return veggie{
		abstractIngredient{
			name: "oregano",
		},
	}
}

// Изготовить баклажан.
func newEggplant() veggie {
	return veggie{
		abstractIngredient{
			name: "eggplant",
		},
	}
}

// Изготовить шпинат.
func newSpinach() veggie {
	return veggie{
		abstractIngredient{
			name: "spinach",
		},
	}
}

// Изготовить оливки.
func newOlives() veggie {
	return veggie{
		abstractIngredient{
			name: "olives",
		},
	}
}

// Изготовить нарезанное пепперони.
func newSlicedPepperoni() pepperoni {
	return pepperoni{
		abstractIngredient{
			name: "sliced pepperoni",
		},
	}
}

// Изготовить свежие мидии.
func newFreshClams() clams {
	return clams{
		abstractIngredient{
			name: "fresh calms",
		},
	}
}

// Изготовить замороженные мидии.
func newFrozenClams() clams {
	return clams{
		abstractIngredient{
			name: "frozen calms",
		},
	}
}
