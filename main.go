package main

import (
	"fmt"
)

func main() {
	for i:=1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("Pim")
		}
		if i%5 == 0 {
			fmt.Printf("Pam")
		}
		if (i%3 != 0 && i%5 != 0){
			fmt.Printf("%v",i)
		}
		fmt.Printf("\n")

	}
}

