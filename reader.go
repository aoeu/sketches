package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type placeholder struct {
	foo string
	bar [512]byte
}

func (p *placeholder) writeTo(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, []byte(p.foo)); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, p.bar); err != nil {
		return err
	}
	return nil
}

func (p *placeholder) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	buf := new(bytes.Buffer)
	fmt.Println("placeholder.Read: Can I avoid copying all of the contents of the source data into a bytes.Buffer?")
	err = p.writeTo(buf)
	fmt.Printf("placeholder.Read: The bytes.Buffer correctly grew to a length of %v, "+
		"but only the first %v bytes are needed.\n", buf.Len(), len(b))
	n = copy(b, buf.Bytes()[0:len(b)])
	return n, err
}

func main() {
	placeholder := &placeholder{foo: "foo", bar: [512]byte{'b', 'a', 'r'}}
	b := make([]byte, 6)
	n, err := placeholder.Read(b)

	fmt.Printf("main: The number of bytes read was %v with error '%v'\n", n, err)
	fmt.Printf("main: Contents of destination byte slice is: %v\n", b)
}
