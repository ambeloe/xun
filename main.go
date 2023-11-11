package main

import (
	"flag"
	"os"
)

func main() {
	os.Exit(rMain())
}

func rMain() int {
	var err error

	var dump = flag.Bool("dump", false, "dump from unity")
	var dumpSelection =
}
