// this example shows how to create a sticky err writer which preserves the error
// very simply programming pattern

package main

import (
	"io"
)

type stickyErrWriter struct {
	err error
	w   io.Writer
}

func (w stickyErrWriter) Write(data []byte) (n int, err error) {
	if w.err != nil {
		return 0, err
	}
	return w.Write(data)
}

func main() {

}
