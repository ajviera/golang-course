package main

import (
	// "bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter a number: ")
	// num, _ := reader.ReadString('\n')
	// fmt.Println("You enter: " + num)
	index, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	} else {
		print(calculate(index))
	}
}

func calculate(index int) ([]int, int) {
	var slice []int
	var total int
	for i := 1; i < index; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			total += i
			slice = append(slice, i)
		}
	}
	return slice, total
}

func print(slice []int, total int) {
	fmt.Println("if we add ")
	fmt.Printf("%v", slice)
	fmt.Println(" ")
	fmt.Println("the result is: " + strconv.Itoa(total))
}
