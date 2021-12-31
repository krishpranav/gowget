package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/cheggaaa/pb"
)

var (
	usage      = "usage: ./gowget -url=http://somewebsite/filename"
	version    = "version: 1.0"
	about      = "wget built using golang"
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

		checkStatus := http.Client{
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				r.URL.Opaque = r.URL.Path
				return nil
			},
		}

		response, err := checkStatus.Get(*cliUrl)

		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		defer response.Body.Close()
		fmt.Printf("Request Status: %s\n\n", response.Status)

		filesize := response.ContentLength

		go func() {
			n, err := io.Copy(file, response.Body)
			if n != filesize {
				fmt.Println("Truncated")
			}
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		}()

		countSize := int(filesize / 1000)
		bar := pb.StartNew(countSize)
		var fi os.FileInfo
		for fi == nil || fi.Size() < filesize {
			fi, _ = file.Stat()
			bar.Set(int(fi.Size() / 1000))
			time.Sleep(time.Millisecond)
		}
		finishMessage := fmt.Sprintf("\n%s with %v bytes downloaded",
			fileName, filesize)
		bar.FinishPrint(finishMessage)

		if err != nil {
			panic(err)
		}

	} else if *cliVersion {
		fmt.Println(flag.Lookup("version").Usage)
	} else if *cliHelp {
		fmt.Println(flag.Lookup("help").Usage)
	} else if *cliAbout {
		fmt.Println("\n\n" + flag.Lookup("about").Usage)
	} else {
		fmt.Println(flag.Lookup("help").Usage)
	}
}
