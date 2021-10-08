package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to know the time!")

	presentTime := time.Now()

	fmt.Println(presentTime)

	//in a particulat format
	// for he date we have to use the ones given below for sure:
	// 01-02-2006
	//Monday
	//15:04:05
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))
}
