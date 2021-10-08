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

	//custom date creation

	//we can make use of same format method to format this date too.
	createdDate := time.Date(2020, time.March, 18, 18, 18, 0, 0, time.UTC)
	fmt.Println("Printing the created Date : ", createdDate)
}
