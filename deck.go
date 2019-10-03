package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck [] string

func NewDeck() string {
	new := "ace"
	return new
}
func writeIntToFile(filename string){
	fmt.Println("writing")
	ioutil.WriteFile(filename, []byte("ace"), 0666)
}

func ReadFile (filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err!= nil {
		fmt.Println("File Not found", err)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func add(a ... int) int {
	var sum int
	for _, v := range a{
		sum = sum + v
	}
	return sum
}

