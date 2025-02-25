package main

import (
	"fmt"
)

func main() {
	var x int
	for {
		fmt.Print("Input [0 to exit]: ")
		fmt.Scan(&x)
		if x<=0 {
			return
		}
		for i := 1; i < x*2; i++ {
			if i<=x {
				for j := 1; j <= i; j++ {
					fmt.Print("*")
				}
			} else {
				for j := 1; j <= x-(i-x); j++ {
					fmt.Print("*")
				}
			}
			fmt.Println()
		}
	}
}
