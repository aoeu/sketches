package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	// TODO(aoeu): Why doesn't `echo foobarbaz | /usr/bin/md5sum -` yield the same sum as the others?
	s := "6df23dc03f9b54cc38a0fc1483df6e21" // https://duckduckgo.com/?q=md5+%22foobarbaz%22&t=canonical&ia=answer
	b := md5.Sum([]byte("foobarbaz"))
	fmt.Printf(" string %s\n with printf %s\n as hex encoded to string %s\n", s, fmt.Sprintf("%x", b), hex.EncodeToString(b[:]))
}
