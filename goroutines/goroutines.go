package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutinesTest() {
	go say("world")
	say("hello")
}

func main() {
	goroutinesTest()
}
