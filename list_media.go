// list_media.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"flag"
	"os"
)

var dir *string = flag.String("directory", "", "Directory to scan")

func main() {
	flag.Parse()
	if *dir == "" {
		fmt.Println("Please specify the directory to scan")
		os.Exit(1)
	}
	list, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Fatal(err)
	} else {
		for _, fi := range list {
			fmt.Println(fi.Name())
		}
	}
}
