package module2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Warehouse() {
	input, _ := os.OpenFile("input.txt", os.O_RDONLY, 0666)
	_, _ = input.WriteString("\n")
	defer input.Close()
	reader := bufio.NewReader(input)
	var line []string
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = append(line, str)
	}
	strgoods := strings.Split(strings.TrimSuffix(strings.TrimSuffix(line[1], "\n"), "\r"), " ")
	var goods []int
	var good int
	for i := 0; i < len(strgoods); i++ {
		good, _ = strconv.Atoi(strgoods[i])
		goods = append(goods, good)
	}
	strorders := strings.Split(strings.TrimSuffix(strings.TrimSuffix(line[3], "\n"), "\r"), " ")
	orders := make([]int, len(strorders))
	for i, order := range strorders {
		num, _ := strconv.Atoi(order)
		orders[i] = num
	}
	for j := 0; j < len(orders); j++ {
		for k := 0; k < len(orders)-j-1; k++ {
			if orders[k] > orders[k+1] {
				orders[k], orders[k+1] = orders[k+1], orders[k]
			}
		}
	}
	orders = orderSort(orders, orders[len(orders)-1])
	output, _ := os.OpenFile("output.txt", os.O_APPEND|os.O_TRUNC|os.O_CREATE, 0644)
	for index, value := range goods {
		if index == len(goods)-1 {
			if value < orders[index] {
				output.WriteString("yes")
			} else {
				output.WriteString("no")
			}
		} else {
			if value < orders[index] {
				output.WriteString("yes\n")
			} else {
				output.WriteString("no\n")
			}
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
