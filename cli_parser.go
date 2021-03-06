package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"fmt"
)

var (
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	name    = kingpin.Arg("name", "Name of user.").Required().String()
)

func main() {
	kingpin.Parse()
	fmt.Printf("%v, %s\n", *verbose, *name)
}