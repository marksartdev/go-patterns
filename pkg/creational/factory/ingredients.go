package factory

// Изготовить тонкое тесто.
func newThinCrustDough() *dough {
	ingredient := new(dough)
	ingredient.name = "Thin Crust Dough"

	return ingredient
}

// Изготовить толстое тесто.
func newThickCrustDough() *dough {
	ingredient := new(dough)
	ingredient.name = "Thick Crust Dough"

	return ingredient
}

// Изготовить соус "Маринара".
func newMarinaraSauce() *sauce {
	ingredient := new(sauce)
	ingredient.name = "Marinara Sauce"

	return ingredient
}

// Изготовить томатный соус.
func newPlumTomatoSauce() *sauce {
	ingredient := new(sauce)
	ingredient.name = "Plum Tomato Sauce"

	return ingredient
}

// Изготовить сыр "Реджиано".
func newReggianoCheese() *cheese {
	ingredient := new(cheese)
	ingredient.name = "Grated Reggiano Cheese"

	return ingredient
}

// Изготовить сыр "Моцарелла".
func newMozzarellaCheese() *cheese {
	ingredient := new(cheese)
	ingredient.name = "Shredded Mozzarella Cheese"

	return ingredient
}

// Изготовить сыр "Пармезан".
func newParmesanCheese() *cheese {
	ingredient := new(cheese)
	ingredient.name = "Parmesan Cheese"

	return ingredient
}

// Изготовить чеснок.
func newGarlic() *veggie {
	ingredient := new(veggie)
	ingredient.name = "garlic"

	return ingredient
}

// Изготовить лук.
func newOnion() *veggie {
	ingredient := new(veggie)
	ingredient.name = "onion"

	return ingredient
}

// Изготовить грибы.
func newMushroom() *veggie {
	ingredient := new(veggie)
	ingredient.name = "mushroom"

	return ingredient
}

// Изготовить красный перец.
func newRedPepper() *veggie {
	ingredient := new(veggie)
	ingredient.name = "red pepper"

	return ingredient
}

// Изготовить орегано.
func newOregano() *veggie {
	ingredient := new(veggie)
	ingredient.name = "oregano"

	return ingredient
}

// Изготовить баклажан.
func newEggplant() *veggie {
	ingredient := new(veggie)
	ingredient.name = "eggplant"

	return ingredient
}

// Изготовить шпинат.
func newSpinach() *veggie {
	ingredient := new(veggie)
	ingredient.name = "spinach"

	return ingredient
}

// Изготовить оливки.
func newOlives() *veggie {
	ingredient := new(veggie)
	ingredient.name = "olives"

	return ingredient
}

// Изготовить нарезанное пепперони.
func newSlicedPepperoni() *pepperoni {
	ingredient := new(pepperoni)
	ingredient.name = "sliced pepperoni"

	return ingredient
}

// Изготовить свежие мидии.
func newFreshClams() *clams {
	ingredient := new(clams)
	ingredient.name = "fresh calms"

	return ingredient
}

// Изготовить замороженные мидии.
func newFrozenClams() *clams {
	ingredient := new(clams)
	ingredient.name = "frozen calms"

	return ingredient
}
