package main

import (
	"fmt"

	"github.com/marksartdev/go-patterns/pkg/structural/dproxy"
)

func main() {
	joe := createPerson("Joe Bean", "male", "bowling")

	ownerProxy := dproxy.NewOwnerProxy(joe)

	fmt.Println("\n---OWNER PROXY---")
	fmt.Printf("Name is %s\n", ownerProxy.GetName())
	fmt.Printf("Old interests: %s\n", ownerProxy.GetInterests())
	fmt.Println("Trying to set interests...")
	ownerProxy.SetInterests("bowling, Go")
	fmt.Printf("New interests: %s\n", ownerProxy.GetInterests())
	fmt.Printf("Old raiting: %d\n", ownerProxy.GetHotOrNotRating())
	fmt.Println("Trying to set rating...")
	// nolint:gomnd // Only for test
	ownerProxy.SetHotOrNotRating(5)
	fmt.Printf("New raiting: %d\n", ownerProxy.GetHotOrNotRating())

	nonOwnerProxy := dproxy.NewNonOwnerProxy(joe)

	fmt.Println("\n---NON-OWNER PROXY---")
	fmt.Printf("Name is %s\n", nonOwnerProxy.GetName())
	fmt.Printf("Old interests: %s\n", nonOwnerProxy.GetInterests())
	fmt.Println("Trying to set interests...")
	nonOwnerProxy.SetInterests("bowling, Go")
	fmt.Printf("New interests: %s\n", nonOwnerProxy.GetInterests())
	fmt.Printf("Old raiting: %d\n", nonOwnerProxy.GetHotOrNotRating())
	fmt.Println("Trying to set rating...")
	// nolint:gomnd // Only for test
	nonOwnerProxy.SetHotOrNotRating(5)
	fmt.Printf("New raiting: %d\n", nonOwnerProxy.GetHotOrNotRating())

	fmt.Println()
}

func createPerson(name, gender, interests string) dproxy.PersonBean {
	person := dproxy.NewPersonBean()
	person.SetName(name)
	person.SetGender(gender)
	person.SetInterests(interests)

	return person
}
