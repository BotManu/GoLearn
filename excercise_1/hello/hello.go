package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func add(a, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var variable, b, m, n bool //package level variables
// := short assignment not working outside a function

func main() {

	var number1 float32 = 6.898
	var number2 float64 = 9.0432
	var name string = "Emanuel"

	number1 = float32(number2)
	fmt.Printf("%v of type %T\n", number1, number1)
	fmt.Printf("%v of type %T\n", number2, number2)
	fmt.Printf("Name is %s\n", name)

	number := rand.Intn(100)
	square_num := math.Sqrt(float64(number))

	fmt.Printf("My favourite number is %d and its square root is %.2f\n", number, square_num)

	//Exported names begin with capital letter in Go
	fmt.Printf("Pi is %f\n", math.Pi)

	fmt.Print("Sum of 2 and 3 is: ", add(2, 3), "\n")

	myString1, myString2 := swap("Hello", "Woirld")
	fmt.Println("Swapped strings are: ", myString1, myString2)
	fmt.Println(swap("I am", "Oh my"))
	fmt.Println(split(7))

	fmt.Println(variable, b, m, n)

	fmt.Println("Example of a complex number is", cmplx.Sqrt(-4-4i))
}
