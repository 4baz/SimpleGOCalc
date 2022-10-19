package main

import (
	"fmt"

	//	"rsc.io/quote"
)

	var num1 int;
	var num2 int;



func addition(numerouno int, numerodos int) (result int, nicetext string){
		result = numerouno + numerodos;

		s:= fmt.Sprint(numerouno, "+ ",numerodos, "=")
		nicetext = s;

return;
	}
func subtraction(numerouno int, numerodos int) int{
	return (numerouno + numerodos);
}
func multiply(numerouno int, numerodos int) int {
	return (numerouno * numerodos);
}
func divide(numerouno int, numerodos int) (result int){
	result = (numerouno / numerodos);
	return
}
var numberswitch int;
var running bool = true;
func main() {

	//numberswitch := 0;


	//running := true;

	if (running){
		fmt.Printf("Simple Go Calculator\n");
		running = false;
	}


	fmt.Printf("Please input first number\n");

	fmt.Scan(&num1)
fmt.Printf("Please enter second number\n")
fmt.Scan(&num2)

	fmt.Println("The first number is",num1,"second is",num2);
fmt.Println("What would you like to do now?")
fmt.Println("1: addition\n2: subtration\n3: multiply\n4: divide\n5: restart")

fmt.Scan(&numberswitch)

	switch(numberswitch){

	case 1:

		result, text :=addition(num1,num2);

		fmt.Println(text, result)
	case 2:
		fmt.Println(subtraction(num1,num2))
	case 3:
		fmt.Println(multiply(num1,num2))

	case 4:
		fmt.Println(divide(num1,num2))

	case 5:
		main()

	default:
		if (numberswitch == 0){
			fmt.Println("Not a valid input")
		}
	}


}
