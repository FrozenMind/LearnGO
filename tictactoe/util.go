package main

import (
	"fmt"
	"strconv"
)

func printField(arr [3][3]string) {
	fmt.Println("---------------")
	for i := 0; i < len(arr); i++ {
		fmt.Print(" | ")
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " | ")
		}
		fmt.Println()
		fmt.Println("---------------")
	}
}

func strToInt(s string) int {
	//i, _ := strconv.ParseInt(s, 10, 64)
	i, _ := strconv.Atoi(s)
	return i
}
