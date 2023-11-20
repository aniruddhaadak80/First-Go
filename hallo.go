package main

import "fmt"

func main() {
	fmt.Println("Hallo Go")

	//Single line comment in Go
	/*Multi-line comment in Go */

	//variable in Go
	var name = "Aniruddha @ Go"
	roll := 30
	fmt.Println("My name is :", name)
	fmt.Println("My roll is :", roll)

	//Array in Go
	var arrNum = [3]int{3, 4, 5}
	fmt.Println("Array of number : ", arrNum)

	var cars = [...]string{"tesla", "lamborgrini", "maruti", "tesla", "tesla", "lamborgrini", "maruti", "tesla"}
	fmt.Println("Cars I will have :", cars)

	//Arithmatic operations
	var result = 3 + 4*0 - 4/2*3/3 //output : 1
	fmt.Println("Result of  3 + 4*0 - 4/2*3/3  is:", result)

	//conditions
	if 5 > 6 {
		fmt.Println("5 is greter than 6")
	} else if 6 > 5 {
		fmt.Println("6 is greter than 5. Condition satisfied...")

	} else {
		fmt.Println("5 is equal to  6 ")
	}

	fmt.Printf("I 💜 Go \n")
	fmt.Print("I have to Go ,But I will come back next time \n\n\n")

	/*End of programme of Go . I 💜 Go */
}
