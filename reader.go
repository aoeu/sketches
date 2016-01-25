package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

type placeholder struct {
	foo string
	bar [512]byte
}

func (p *placeholder) WriteTo(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, []byte(p.foo)); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, p.bar); err != nil {
		return err
	}
	return nil
}

var bw *bufio.Writer

func (p *placeholder) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	r, w := io.Pipe()
	bw = bufio.NewWriterSize(w, len(b))
	go p.WriteTo(bw)
	return r.Read(b)
}

func main() {
	placeholder := &placeholder{foo: "foo", bar: [512]byte{'b', 'a', 'r'}}
	b := make([]byte, 6)
	n, err := placeholder.Read(b)
	fmt.Printf("main: buffer size after write is %v\n", bw.Buffered())
	fmt.Printf("main: The number of bytes read was %v with error '%v'\n", n, err)
	fmt.Printf("main: Contents of destination byte slice is: %v\n", b)
}
