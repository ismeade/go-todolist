//
// main.go
//

package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	h bool
)

func main() {

	if h {
		flag.Usage()
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	if len(args) >= 1 {
		switch args[0] {
		case "l", "list":
			fmt.Printf("%s\n", "list")
		case "a", "add":
			if len(args) == 3 {

			} else {
				fmt.Println("args error")
			}
		case "u", "update":
			fmt.Printf("%s\n", "update")
		case "d", "download":
			fmt.Printf("%s\n", "download")
		default:
			strType := args[0]
			if len(strType) > 10 {
				strType = strType[:10] + "..."
			}

			fmt.Println("No support:", strType)
		}
	} else {
		fmt.Printf("args error: %d\n", len(args))
	}
}

func init() {
	flag.BoolVar(&h, "h", false, "help")

	flag.Parse()
	flag.Usage = usage
}

func usage() {
	fmt.Println(`help: todo-list`)
	flag.PrintDefaults()
}
