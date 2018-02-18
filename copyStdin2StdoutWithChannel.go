package main

import (
	"io"
	"os"
	"time"
)

func main() {
	go echo(os.Stdin, os.Stdout)
	// Blocking code. Makes main routine to wait for channel to be filled
	<-time.After(5 * time.Second)
}

func echo(out io.Writer, in io.Reader) {
	io.Copy(out, in)
}
