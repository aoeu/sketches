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

const barSize int = 512

type placeholderReader struct {
	offset int
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

func (p *placeholder) readWithPipeAndBufio(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	r, w := io.Pipe()
	bw = bufio.NewWriterSize(w, len(b))
	go p.WriteTo(bw)
	return r.Read(b)
}

func print(size, read int, err error) {
	fmt.Printf("print: buffer size after write is %v\n", size)
	fmt.Printf("print: The number of bytes read was %v with error '%v'\n", read, err)
}


func main() {
	foo := "foo"
	placeholder := &placeholder{foo: foo, bar: [barSize]byte{'b', 'a', 'r'}}
	
	// Works for a destination byte slice that is a smaller size than all source data bytes.
	b := make([]byte, 6)
	n, err := placeholder.readWithPipeAndBufio(b)
	print(bw.Buffered(), n, err)

	// Deadlocks for a destination byte slice that is equal to or greater than all source data bytes.
	/*
	b = make([]byte, barSize + len(foo))
	n, err = placeholder.readWithPipeAndBufio(b)
	print(bw.Buffered(), n, err)
	*/

	fmt.Printf("main: Contents of destination byte slice is: %v\n", b)
}
