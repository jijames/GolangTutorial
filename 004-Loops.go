package main

import "fmt"

// Video link: https://youtu.be/h2uYmtyS85w

func main() {
	TEST := 20
	END := 30

	/* Multi-line comment
	for i := 1; i <= TEST; i++ {
		fmt.Println(i)
	}*/

	for TEST <= END {
		fmt.Println(TEST)
		TEST++
	}

}
