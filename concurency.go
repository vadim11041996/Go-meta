package main

import (
	"fmt"
	"time"
)

func Concurency() {
	c := make(chan string)
	go count("task", c)

	for msg := range c {
		fmt.Println(msg)
	}

}

func count(task string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- task
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
