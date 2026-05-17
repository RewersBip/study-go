package main

import "fmt"

func main() {
	a := 0
	a = 2
	fmt.Scan(&a)
	fmt.Println(a)

	forint, forint2, forstring, forbool := doublea(a)
	fmt.Println(forint)
	fmt.Println(forint2)
	fmt.Println(forstring)
	fmt.Println(forbool)

}
func doublea(a int) (int, int, string, bool) {
	a = a * 2
	return a, 5, "asdasd", true
}
