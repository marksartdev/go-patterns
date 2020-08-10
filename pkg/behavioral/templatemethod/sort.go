package templatemethod

import (
	"fmt"
	"strings"
)

// Duck Утка.
type Duck struct {
	Name   string
	Wright int
}

func (s Duck) String() string {
	return fmt.Sprintf("%s weighs %d", s.Name, s.Wright)
}

// Ducks Коллекция уток.
type Ducks struct {
	Ducks []Duck
}

func (d Ducks) String() string {
	log := make([]string, 0, d.Len())

	for _, duck := range d.Ducks {
		log = append(log, duck.String())
	}

	return strings.Join(log, "\n")
}

// Len Возвращает количество уток в коллекции.
func (d Ducks) Len() int {
	return len(d.Ducks)
}

// Less Возвращает должна ли утка i быть перед j.
func (d Ducks) Less(i, j int) bool {
	return d.Ducks[i].Wright < d.Ducks[j].Wright
}

// Swap Меняет утки местами.
func (d Ducks) Swap(i, j int) {
	tmp := d.Ducks[j]
	d.Ducks[j] = d.Ducks[i]
	d.Ducks[i] = tmp
}
