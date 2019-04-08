package logger

import (
	"github.com/arthurkiller/rollingWriter"
)

type RotateWriter struct {
	// output destination.
	Writer rollingwriter.RollingWriter
}

// NewRotateWriter creates and initializes a new RotateWriter.
func NewRotateWriter(options ...rollingwriter.Option) RotateWriter {
	// Create rolling writer from config options
	writer, _ := rollingwriter.NewWriter(options...)

	rotateWriter := RotateWriter{
		Writer: writer,
	}

	return rotateWriter
}

// Implement io.Writer interface
func (rw RotateWriter) Write(p []byte) (n int, err error) {
	return rw.Writer.Write(p)
}
