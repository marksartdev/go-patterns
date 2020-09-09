package dproxy_test

import (
	"bytes"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/marksartdev/go-patterns/pkg/structural/dproxy"
)

func TestNewOwnerProxy(t *testing.T) {
	buffer := bytes.NewBufferString("")
	log.SetOutput(buffer)

	person := createPerson()
	ownerProxy := dproxy.NewOwnerProxy(person)
	ownerProxy.SetName("New Name")
	ownerProxy.SetGender("female")
	ownerProxy.SetInterests("testing, Go")
	assert.Empty(t, buffer.String())
	ownerProxy.SetHotOrNotRating(5)
	assert.NotEmpty(t, buffer.String())

	buffer.Reset()

	assert.Equal(t, "New Name", ownerProxy.GetName())
	assert.Equal(t, "female", ownerProxy.GetGender())
	assert.Equal(t, "testing, Go", ownerProxy.GetInterests())
	assert.Equal(t, 0, ownerProxy.GetHotOrNotRating())
	assert.Empty(t, buffer.String())
}

func TestNewNonOwnerProxy(t *testing.T) {
	buffer := bytes.NewBufferString("")
	log.SetOutput(buffer)

	person := createPerson()
	nonOwnerProxy := dproxy.NewNonOwnerProxy(person)
	nonOwnerProxy.SetName("New Name")
	assert.NotEmpty(t, buffer.String())
	buffer.Reset()
	nonOwnerProxy.SetGender("female")
	assert.NotEmpty(t, buffer.String())
	buffer.Reset()
	nonOwnerProxy.SetInterests("testing, Go")
	assert.NotEmpty(t, buffer.String())
	buffer.Reset()
	nonOwnerProxy.SetHotOrNotRating(5)
	assert.Empty(t, buffer.String())

	assert.Equal(t, "Test Mock", nonOwnerProxy.GetName())
	assert.Equal(t, "male", nonOwnerProxy.GetGender())
	assert.Equal(t, "testing", nonOwnerProxy.GetInterests())
	assert.Equal(t, 5, nonOwnerProxy.GetHotOrNotRating())
	assert.Empty(t, buffer.String())
}

func createPerson() dproxy.PersonBean {
	person := dproxy.NewPersonBean()
	person.SetName("Test Mock")
	person.SetGender("male")
	person.SetInterests("testing")

	return person
}
