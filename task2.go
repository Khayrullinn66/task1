package main

import (
	"fmt"
	"sync"
)

func counter1() {
	fmt.Println("Решение 1")
	array := [...]int{2, 4, 6, 8, 10}
	var vt sync.WaitGroup
	vt.Add(len(array))
	for _, i := range array {
		go func(i int) {
			fmt.Print(i*i, "\n")
			vt.Done()
		}(i)
	}
	vt.Wait()
}

func counter2() {
	fmt.Println("Решение 2")
	mass := [...]int{2, 4, 6, 8, 10}
	do := make(chan interface{}, len(mass))
	defer close(do)
	for _, i := range mass {
		go func(i int) {
			fmt.Print(i*i, "\n")
			do <- struct{}{}
		}(i)
	}
	for i := 0; i < len(mass); i++ {
		<-do
	}
}
func main() {
	counter1()
	counter2()
}
