package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		max int
		msg string
		x   bool
	)

	flag.IntVar(&max, "max", 64, "maxmaxmax")
	flag.StringVar(&msg, "msg", "", "msg msg msg")
	flag.BoolVar(&x, "x", false, "bool bool bool")

	flag.Parse()

	fmt.Println(msg)
	fmt.Println(max)
	fmt.Println(x)
}
