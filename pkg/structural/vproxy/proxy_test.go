package vproxy_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/marksartdev/go-patterns/pkg/structural/vproxy"
)

func TestNewExchanger(t *testing.T) {
	buffer := bytes.NewBufferString("")

	exchanger := vproxy.NewExchanger("2020-09-10")
	exchanger.SetWriter(buffer)
	exchanger.Show()

	assert.Equal(t, "Loading...\n", buffer.String())
	assert.Empty(t, exchanger.GetRates())

	buffer.Reset()

	time.Sleep(7 * time.Second)

	expectedMap := map[string]float64{
		"USD": 75.0189889442,
		"EUR": 88.89,
	}

	assert.Equal(t, expectedMap, exchanger.GetRates())
	assert.NotEmpty(t, buffer.String())

	buffer = bytes.NewBufferString("")
	exchanger.SetWriter(buffer)

	exchanger.Show()
	assert.NotEmpty(t, buffer.String())
}
