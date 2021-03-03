package main

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
)

func main() {
	var decode bool
	flag.BoolVar(&decode, "d", false, "decode")
	flag.Parse()

	if !isPipe() {
		fmt.Println("only piped input for now")
		os.Exit(1)
	}

	buf, err := io.ReadAll(os.Stdin)
	check(err)
	inp := string(buf)

	var res string

	if decode {
		b, err := url.QueryUnescape(inp)
		check(err)
		res = string(b)
	} else {
		res = url.QueryEscape(inp) + "\n"
	}
	fmt.Printf("\n%s\n", res)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func isPipe() bool {
	info, err := os.Stdin.Stat()
	check(err)
	return (info.Mode() & os.ModeCharDevice) == 0
}
