package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

type person struct {
	name string
	age  int
}

func main() {

	//goroutine
	go count("first task")
	go count("seccond task")

	fmt.Scanln() // will execute the code until we press enter

	//sum
	result := sum(2, 3)
	fmt.Println(result)

	//loop
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println(index, value)
	}

	//map
	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["circle"] = 3
	vertices["square"] = 4

	delete(vertices, "square")
	fmt.Println(vertices)

	//pointer increment value
	i := 7
	increment(&i)
	fmt.Println(i)

	//struct
	p := person{name: "Jake", age: 23}
	fmt.Println(p)

	//sqrt get
	res, err := sqrt(-9)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func count(task string) {
	for i := 1; true; i++ {
		fmt.Println(i, task)
		time.Sleep(time.Millisecond * 500)
	}
}

func increment(x *int) {
	*x++
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined , negative number")
	}
	return math.Sqrt(x), nil
}
