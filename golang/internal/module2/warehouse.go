package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Warehouse() {
	reader := bufio.NewReader(os.Stdin)
	arrlen1, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	intarrlen1, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimSuffix(arrlen1, "\n"), "\r"))
	if intarrlen1 > 100 || intarrlen1 < 1 {
		return
	}
	line1, err := reader.ReadString('\n')
	strarr1 := strings.Split(strings.TrimSuffix(strings.TrimSuffix(line1, "\n"), "\r"), " ")
	var goods []int
	for i := 0; i < len(strarr1); i++ {
		inti, _ := strconv.Atoi(strarr1[i])
		goods = append(goods, inti)
	}
	if (len(goods) != intarrlen1) || intarrlen1 > 10000 || intarrlen1 < 1 {
		return
	}
	arrlen2, err := reader.ReadString('\n')
	intarrlen2, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimSuffix(arrlen2, "\n"), "\r"))
	line2, err := reader.ReadString('\n')
	strarr2 := strings.Split(strings.TrimSuffix(strings.TrimSuffix(line2, "\n"), "\r"), " ")
	var orders []int
	for i := 0; i < len(strarr2); i++ {
		inti, _ := strconv.Atoi(strarr2[i])
		orders = append(orders, inti)
	}
	if (len(orders) != intarrlen2) || intarrlen2 > 10000 || intarrlen2 < 1 {
		return
	}
	for j := 0; j < len(orders)-1; j++ {
		for k := 0; k < len(orders)-j-1; k++ {
			if orders[k] > orders[k+1] {
				orders[k], orders[k+1] = orders[k+1], orders[k]
			}
		}
	}
	if orders[0] < 1 || orders[len(orders)-1] > intarrlen1 {
		return
	}
	orders = orderSort(orders, orders[len(orders)-1])
	for index, value := range goods {
		if value < orders[index] {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

func orderSort(orders []int, maxo int) []int {
	c := make([]int, maxo+1)
	for i := 0; i < len(orders); i++ {
		c[orders[i]]++
	}
	return c[1:]
}
