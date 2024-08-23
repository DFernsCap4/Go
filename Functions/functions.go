package main

import (
	"fmt"
	"errors"
	"math"
)

func add(a int, b int) int {
	return a+b
}

func divide(a float64, b float64) (result float64, err error) {
	if b == 0 {
		err = fmt.Errorf("division by zero")
		return
	}
	result = a/b
	return
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums{
		total += num
	}
	return total	
}

// func guessNumber(num int) int{
// 	if num < 0 || num > 100 {
// 		return 0
// 	}

// 	randomInt := rand.Intn(101)
// 	fmt.Println(randomInt)
// 	if randomInt == num {
// 		return randomInt
// 	}else {
// 		return guessNumber(num)
// 	}
// }

func sqrt(number float64) (float64, error){
	if number < 0 {
		return 0, errors.New("cannot calculate the square root of a negative number")
	}
	return math.Sqrt(number), nil
}

func main() {
	fmt.Println(add(2,3))

	result, err := divide(4.5, 2)
	if err != nil {
		fmt.Println("Error:", err)
	}else {
		fmt.Println("Result:", result)
	}
	fmt.Println("Sum:", sum(1,2,3,4,5,6,7,8,9,10,11))

	//fmt.Printf("Your number is %d\n", guessNumber(50))

	result, err = sqrt(-9)
	if err != nil {
		fmt.Println("Error:", err)
	}else{
		fmt.Println("Result:", result)
	}
}