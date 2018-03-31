package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockWriter struct {
	nock.Mock
}

func (w *MockWriter) Write(data []byte) (n int, err error) {
	w.
}

func TestStickyWriter(t *testing.T) {

}
