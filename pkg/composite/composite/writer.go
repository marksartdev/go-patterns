package composite

import (
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"
)

type customWriter struct {
	writer io.Writer
}

// SetWriter Set writer.
func (c *customWriter) SetWriter(writer io.Writer) {
	c.writer = writer
}

func (c *customWriter) write(msg string) {
	_, err := fmt.Fprintln(c.writer, msg)
	if err != nil {
		log.Error(err)
	}
}
