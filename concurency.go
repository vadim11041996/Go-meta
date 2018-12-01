package main

import (
	"fmt"
	"time"
)

func Concurency() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 secconds"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}

}
