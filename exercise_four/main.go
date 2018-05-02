package main

// 4º Ejercicio
// Modelar la funcionalidad para un sistema de cines que calcule los ingresos netos de una
// función en base a los asistentes y al precio base de la entrada. Existen tres tipos
// de asistentes y tienen las siguientes características:
// - Normales: Pagan el 100%
// - Jubilados: Tienen un 50% de descuento
// - Invitados: No pagan nada

import (
	"fmt"
	"strconv"

	"github.com/ajviera/golang-course/exercise_four/entities"
)

func main() {
	slice := fakeData()
	funsion := createFunsion(slice)
	incomes := calculateTotal(funsion)

	fmt.Println("Total de asistentes: " + strconv.Itoa(len(funsion.Visitors)))
	fmt.Println("Precio de la funsión: $" + floatToString(funsion.Price))
	fmt.Println("Total de ingresos: $" + floatToString(incomes))
}

func calculateTotal(funsion entities.Funsion) float64 {
	var total float64
	for i := 0; i < len(funsion.Visitors); i++ {
		total += funsion.Price * (funsion.Visitors[i].Category.Discount / 100)
	}
	return total
}

func floatToString(num float64) string {
	result := strconv.FormatFloat(num, 'f', 2, 64)
	return result
}

func createFunsion(slice []entities.Visitor) entities.Funsion {
	var funsion entities.Funsion
	funsion.Visitors = slice
	funsion.Price = 90.2
	return funsion
}

func fakeData() []entities.Visitor {
	var slice []entities.Visitor
	normal := entities.Category{Name: "Normal", Discount: 0}
	pensioner := entities.Category{Name: "Jubilado", Discount: 50}
	guest := entities.Category{Name: "Invitado", Discount: 100}

	slice = append(slice, entities.Visitor{"A", 20, normal})
	slice = append(slice, entities.Visitor{"B", 75, pensioner})
	slice = append(slice, entities.Visitor{"C", 60, normal})
	slice = append(slice, entities.Visitor{"D", 11, guest})

	return slice
}
