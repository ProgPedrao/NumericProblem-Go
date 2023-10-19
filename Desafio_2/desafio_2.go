package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {

		multiplo3 := i%3 == 0
		multiplo5 := i%5 == 0

		if multiplo3 && multiplo5 {
			fmt.Println("Pin Pan")
		} else if multiplo3 {
			fmt.Println("Pin")
		} else if multiplo5 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}

	}
}
