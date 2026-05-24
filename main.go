package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	fmt.Println("приветик")
	for {
		fmt.Println("хотите ли рассчитать ваш индекс массы тела? 1 - да, 0 - нет")
		x := 1
		fmt.Scan(&x)

		if x == 1 {
			fmt.Println("Введите ваш рост(в сантимертрах)")
			var Height float64
			fmt.Scan(&Height)
			fmt.Println("Введите ваш вес")
			var Kg float64
			fmt.Scan(&Kg)
			result := Kg / math.Pow(Height/100, IMTPower)
			fmt.Printf("ваш индекс массы тела %.0f", result)
		} else if x == 0 {
			fmt.Println("пока")
			break
		}
	}
}
