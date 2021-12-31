package main

import (
	"flag"
	"fmt"
)

var (
	usage      = "usage: ./gowget -url=http://somewebsite/somefile"
	version    = "version: 1.0"
	about      = "wget build using golang"
	help       = fmt.Sprintf("\n\n  %s\n\n\n  %s\n\n\n  %s", usage, version, about)
	cliUrl     *string
	cliVersion *bool
	cliHelp    *bool
	cliAbout   *bool
)

func init() {
	cliUrl = flag.String("url", "", usage)
}

func main() {
	fmt.Printf(version)
}
