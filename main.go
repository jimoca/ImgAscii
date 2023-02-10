package main

import (
	"flag"
	"fmt"
	"imgAscii/util"
)

var imageFile string

func init() {
	flag.StringVar(&imageFile, "f", "", "Filename")
}

func main() {

	flag.Parse()
	fmt.Println(util.Convert(imageFile))
}
