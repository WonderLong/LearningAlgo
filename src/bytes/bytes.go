package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte("this is a test")
	a := new(bytes.Buffer)
	a.Write(b)
	fmt.Println(string(b))

}
