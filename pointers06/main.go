package main

import "fmt"

func main() {
	fmt.Println("Here we go")

	//var one int = 2
	var ptr *int //int pointer
	fmt.Println("the pointer default is as follows,", ptr)
	//default value of ptr is  <nil>

	myNum := 23

	ptr = &myNum

	fmt.Println("the value of ptr now will be , ", ptr)
	fmt.Println("the value of the value in ptr will be ,", *ptr)

	//results are as follows:
	//the value of ptr now will be ,  0xc000012088
	// the value of the value in ptr will be , 23
	//as in all programming languages can modify values in ptr using ptr directly.
	*ptr = *ptr * 2
	fmt.Println("the value of the value in ptr will be ,", *ptr)
	//value is as follows:
	//the value of the value in ptr will be , 46

}
