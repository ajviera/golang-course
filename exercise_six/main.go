package main

// 6º Ejercicio
// Evolucionar el primer ejercicio del curso de forma tal que ahora se ejecute
// en un webserver que escuche en el puerto 8080 y sea invocado al llamar al servicio
// GET /calculate?value=p. El servicio debe devolver la respuesta en formato JSON

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/calculate", func(c *gin.Context) {
		value, err := strconv.Atoi(c.Query("value"))
		if err != nil {
			fail(c)
		} else {
			success(c, value)
		}
	})
	r.Run()
}

func fail(c *gin.Context) {
	c.JSON(400, gin.H{
		"error": "invalid value, must be integer",
	})
}

func success(c *gin.Context, value int) {
	slice, total := calculate(value)
	c.JSON(200, gin.H{
		"slice": slice,
		"total": total,
	})
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
