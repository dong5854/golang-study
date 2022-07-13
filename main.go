package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	people := [2]string{"dong", "chae"}
	for _, person := range people {
		go smartCount(person, channel)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-channel)
	}
}

func smartCount(person string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(person, "is smart", i)
		time.Sleep(time.Second)
	}
	channel <- person + " is smart!"
}
