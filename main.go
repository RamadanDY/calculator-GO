package main

import "fmt"

func main() {
	var num1, num2 int
	var operation string

	fmt.Println("calculator 1.0")
	fmt.Println("===============")
	fmt.Println("which operation do u want to perform (add,sub,multiply)")

	/////
	//with this it means we access the pointer not the copy of the variable

	// this make it posible to change the value of the variable itself

	//so now lets get the numbers form the users

	fmt.Scanf("%s", &operation)

	fmt.Println("enter the first number")
	fmt.Scanf("%v", &num1)
	fmt.Println("enter the second digit")
	fmt.Scanf("%v", &num2)

	switch operation {
	case "add":
		fmt.Println("the addition of the two numbewrs are:", num1+num2)
	}

}

/*
fmt.Scanf("%s", &operation)
	fmt.Println("enter the first value")
	fmt.Scanf("%v", &num1)
	fmt.Println("enter the second value")
	fmt.Scanf("%v", &num2)

	switch operation {
	case "add":
		fmt.Println("the number that we added is:", num1+num2)

	case "subtract":
		fmt.Println(num1 - num2)

	case "multiply":
		fmt.Println(num1 * num2)

	case "divide":
		fmt.Println(num1 / num2)
	}
*/
