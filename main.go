package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
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

	cliVersion = flag.Bool("version", false, version)

	cliHelp = flag.Bool("help", false, help)

	cliAbout = flag.Bool("about", false, about)
}

func main() {
	flag.Parse()

	if *cliUrl != "" {
		fmt.Println("\nDownloading file...\n")

		fileUrl, err := url.Parse(*cliUrl)

		if err != nil {
			panic(err)
		}

		filePath := fileUrl.Path

		segments := strings.Split(filePath, "/")

		fileName := segments[len(segments)-1]

		file, err := os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		defer file.Close()
	}
}
