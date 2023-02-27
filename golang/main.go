package main

import (
	"fmt"

	"github.com/Kirill-Bokov/Algorithm_2023/golang/internal/module1"
)

func main() {
	fmt.Println("Hello world")
	module1.Summ()
	for n := 0; n < 40; n++ {
		for m := 0; m < 40; m++ {
			module1.Alg1time(n, m)
			module1.Alg2time(n, m)
		}
	}
	fmt.Println("What's wrong?")
}
