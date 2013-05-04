// list_media.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Please specify the directory to scan.")
		os.Exit(1)
	}
	list, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	} else {
		for _, fi := range list {
			fmt.Println(fi.Name())
		}
	}
}
