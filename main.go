package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	fmt.Println("приветик")
	for {
		fmt.Println("хотите ли рассчитать вашу массу тела? 1 - да, 0 - нет")
		x := 1
		fmt.Scan(&x)

		if x == 1 {
			fmt.Println("Введите ваш рост(В МЕТРАХ)")
			var Height float64
			fmt.Scan(&Height)
			fmt.Println("Введите ваш вес")
			var Kg float64
			fmt.Scan(&Kg)
			result := Kg / math.Pow(Height, IMTPower)
			fmt.Println(result)
		} else if x == 0 {
			fmt.Println("пока")
			break
		}
	}

}
