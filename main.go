//
// main.go
// Copyright (C) 2020 liyang <yang.li@51vcheck.cn>
//
//

package main

import (
	"flag"
	"fmt"
	"go-todolist/data"
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
		case "list":
			fmt.Printf("%s\n", "list")
		case "add":
			if len(args) == 3 {
				data.Add(args[1], args[2])
			} else {
				fmt.Println("args error")
			}
		case "u", "update":
			fmt.Printf("%s\n", "update")
		case "download":
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
