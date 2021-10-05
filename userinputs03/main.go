package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//working with wollbus operator
	welcome := "Welcome to Pzzza Place"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the rating to the pizza in between 1 to 5: ")

	//in the bracker we show until which the string has to be read
	input, _ := reader.ReadString('\n')

	fmt.Println("Thank you for rating the pizza : ", input)
}
