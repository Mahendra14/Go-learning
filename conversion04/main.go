package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza in between 1-5 :")
	input, _ := reader.ReadString('\n')
	// fmt.Printf(input)

	//to add one to rating
	//we get an error if we implicitly take of the trailing characters as it includes new line too.
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Added 1 to rating to make it ", numRating+1)
	}

}
