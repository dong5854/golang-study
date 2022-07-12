package main

import (
	"fmt"
	"github.com/dong5854/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "Greeting")
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}
	dictionary.Delete(baseWord)
	word, err2 := dictionary.Search(baseWord)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(word)
}
