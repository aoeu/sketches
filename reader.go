package main

// http://play.golang.org/p/bd8oG8VDmm

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

type placeholder struct {
	foo string
	bar [barSize]byte

	// This buffered writer wouldn't be exist on the real type this placeholder is in lieu of.
	// It exists here so its size can be printed after calls to Read(b []byte) to show
	// that it does not automatically grow like a bytes.Buffer and is therefore a more
	// memory efficient solution when implementing io.Reader for types backed by a lot of data.
	bw *bufio.Writer
}

const barSize int = 512

func (p *placeholder) WriteTo(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, []byte(p.foo)); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, p.bar); err != nil {
		return err
	}
	return io.EOF
}

func (p *placeholder) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	r, w := io.Pipe()
	p.bw = bufio.NewWriterSize(w, len(b))
	go func() {
		p.WriteTo(p.bw)
		// Without this flush, there would be a deadlock
		// if len(b) > bytes written by p.WriteTo(w io.Writer).
		p.bw.Flush()
	}()
	return r.Read(b)
}

func printResults(dstSize, bufSize, read int, err error) {
	fmt.Printf("printResults: Destination byte slice size is %v\n", dstSize)
	fmt.Printf("printResults: Buffer size after write is %v\n", bufSize)
	fmt.Printf("printResults: The number of bytes read was %v with error '%v'\n", read, err)
	fmt.Println()
}

func (p *placeholder) testRead(size int) []byte {
	b := make([]byte, size)
	n, err := p.Read(b)
	printResults(len(b), p.bw.Buffered(), n, err)
	return b
}

func main() {
	foo := "foo"
	p := &placeholder{foo: foo, bar: [barSize]byte{'b', 'a', 'r'}}

	// Successfully reads to a destination byte slice that is a smaller size than all source data bytes.
	p.testRead(6)

	// Successfully reads to a destination byte slice that is the size of all the source data bytes.
	p.testRead(barSize + len(foo))

	// Successfully reads to a destination byte slice that is larger than all of the source data bytes.
	b := p.testRead(9000)

	fmt.Printf("main: Partial contents of destination byte slice is: %v\n", b[0:barSize+len(foo)])
}
