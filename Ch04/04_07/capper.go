package main

import (
	"fmt"
	"io"
	"os"
)

// Capper implements io.Writer and turns everything to uppercase
type Capper struct {
	writer io.Writer
}

func (c *Capper) Write(wordByte []byte) (int, error) {
	// out := make([]byte, len(wordByte))
	diff := byte('a' - 'A')
	fmt.Println(diff)
	c.writer.Write(wordByte)
	return 0, nil
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}
