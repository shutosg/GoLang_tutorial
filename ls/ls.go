package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	directory := "./"
	infos, error := ioutil.ReadDir(directory)
	if error != nil {
		fmt.Println(error)
		return
	}
	for i := 0; i < len(infos); i++ {
		fmt.Println(infos[i].Name())
	}
}
