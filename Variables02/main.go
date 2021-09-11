package main

import "fmt"

//Here it is a constant variable which cannot be changed. And the var name starts with a Capital letter it signifies it is a Public variable.
const PublicVar string = "This is a public value"

func main() {
	//fmt.Println("Variables")
	//String variable
	var stringVar string = "Mahendra"
	fmt.Println(stringVar)
	fmt.Printf("The type of the variable stringVar is : %T \n", stringVar)

	//bool
	var isBoolVar bool = true
	fmt.Println(isBoolVar)
	fmt.Printf("The type of the variable stringVar is : %T \n", isBoolVar)

	//small int unsigned - uint8
	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("The type of the variable stringVar is : %T \n", smallInt)

	//float 32 - we get only five digits after decimal point.
	var floatVar float32 = 255.45466433
	fmt.Println(floatVar)
	fmt.Printf("The type of the variable stringVar is : %T \n", floatVar)

	//default values for int and string being tested.
	var defaultVar int
	fmt.Println(defaultVar)

	var isDefaultVar bool
	fmt.Println(isDefaultVar)

	var stringDefVar string
	fmt.Println(stringDefVar)

	//implicit type def by the lexer
	var impVar = "this is implicitly typed by the lexer"
	fmt.Println(impVar)
	fmt.Printf("The type of the variable impVar is : %T", impVar)

	//no var declaration
	defVar := "Without the definition"
	fmt.Println(defVar)
	fmt.Printf("The var with the defVar is this : %T\n", defVar)

	fmt.Printf("This is a const variable the value of the PublicVar is : ")
	fmt.Println(PublicVar)

}
