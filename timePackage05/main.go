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

	//to create an exe file from this we need to do go build in cmd
	//we can create custom ones to diff os like mac and linux using following commands
	//GOOS:"darwin" go build  :----  in case of mac
	//GOOS:"linux" go build :---- in case of linux
	//to see the value of goos on curretn system : go env and
	//In my case it is as follows in the values: set GOOS=windows
}
