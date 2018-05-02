package main

// 3º Ejercicio (struct, punteros, slices)
// Hacer un programa que reciba por parametro una lista de enteros, los almacene en un slice,
// popule un arbol binario y despues imprima los valores en orden ascendente

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ajviera/golang-course/exercise_three/tree"
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
