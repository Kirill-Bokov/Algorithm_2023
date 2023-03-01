package main

import (
	"fmt"

	"github.com/Kirill-Bokov/Algorithm_2023/golang/internal/module1"
)

func main() {
	alb, bla := false, false
	for n := 0; n < 40; n++ {
		if alb == true && bla == true {
			break
		}
		for m := 0; m < 40; m++ {
			if alb == true && bla == true {
				break
			}
			if (alb == false) && (module1.Alg1time(n, m) > module1.Alg2time(n, m)) {
				fmt.Println("Алг. 1 работает быстрее алг.2:", module1.Alg1time(n, m), module1.Alg2time(n, m), "при N M =", n, m)
				alb = true
			}
			if (bla == false) && (module1.Alg1time(n, m) < module1.Alg2time(n, m)) && (module1.Alg1time(n, m) > 0) {
				fmt.Println("Алг. 2 работает быстрее алг.1:", module1.Alg1time(n, m), module1.Alg2time(n, m), "при N M =", n, m)
				bla = true
			}
		}
	}
	module1.Bubblesort()
}
