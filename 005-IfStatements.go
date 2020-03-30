package main

import "fmt"

// Video link: https://youtu.be/XtqtkjPxWcE

func main() {
	INT1 := 0
	INT2 := 1
	STR1 := "string"
	STR2 := "string"

	if INT1 <= INT2 {
		fmt.Println("INT1 is less than INT2")
	} else if STR1 == STR2 {
		fmt.Println("STR1 is the same as STR2")
	} else {
		fmt.Println("None are true!")
	}

	if STR1 != STR2 {
		fmt.Println("STR1 STR2 check")
	}

	if INT1 > INT2 && STR1 == STR2 {
		fmt.Println("Double test")
	}
}
