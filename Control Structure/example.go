package main

import "fmt"

func main(){
	num := 9

	if num%2 == 0 {
		fmt.Println(num, "is even")
	}else if num%3 == 0{
		fmt.Println(num, "is divisible by 3")
	} else {
		fmt.Println(num, "is odd")
	}

	var grade string
	fmt.Print("Enter your grade: ")
	fmt.Scanln(&grade)
	switch grade{
	case "A":
		fmt.Println("Excellent")
	case "B":
		fmt.Println("Good")
	case "C":
		fmt.Println("Average")
	default:
		fmt.Println("Grade not recognised")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	num = 1
	for num < 5 {
		fmt.Println(num)
		num++
	}

	// for {
	// 	fmt.Println("Infinite loop!")
	// }

	nums := []int{2,4,6,8}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d \n", index, value)
	}

	colors := map[string]string{"red": "#FF0000", "green": "#00FF00", "blue": "#0000FF"}
	for key, value := range colors {
		fmt.Printf("Key: %s, Value: %s \n", key, value)
	}
}