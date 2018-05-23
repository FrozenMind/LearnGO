package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	val, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Timer started with %v seconds\n", val)
	for val > 0 {
		fmt.Printf("%v seconds left\n", val)
		time.Sleep(1 * time.Second)
		val = val - 1
	}
	fmt.Println("TIMER DONE!")

}
