package main

import (
	"exercise_three/tree"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if slice := createSlice(); slice != nil {
		tree := createTree(slice)
		fmt.Println("the values were sorted in ascending order: ")
		fmt.Println(tree.PrintInOrder())
	} else {
		fmt.Println("The entered values are incorrect")
	}
}

func createSlice() []int {
	var slice []int
	for i := 1; i < len(os.Args); i++ {
		index, err := strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println("Append -> " + os.Args[i])
		slice = append(slice, index)
	}
	return slice
}

func createTree(slice []int) *tree.Tree {
	var t = new(tree.Tree)
	for i := 0; i < len(slice); i++ {
		t = t.Add(slice[i])
	}
	return t
}
