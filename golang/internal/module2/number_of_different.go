package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Number_of_different() {
	reader := bufio.NewReader(os.Stdin)
	arrlen, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	intarrlen, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimSuffix(arrlen, "\n"), "\r"))
	line, err := reader.ReadString('\n')
	strarr := strings.Split(strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r"), " ")
	if (len(strarr) != intarrlen) || intarrlen > 100000 || intarrlen < 1 {
		return
	}
	var arr []int
	for i := 0; i < len(strarr); i++ {
		inti, _ := strconv.Atoi(strarr[i])
		arr = append(arr, inti)
	}
	fmt.Println(len(arr))
	fmt.Println(quicksort(arr, 0, len(arr), 0))
}

func quicksort(arr []int, left, right, diff int) ([]int, int) {
	if len(arr) <= 2 {
		return arr, diff
	}
	arr, q, diff := partition(arr, left, right, diff)
	arr, diff = quicksort(arr, left, q, diff)
	arr, diff = quicksort(arr, q, right, diff)
	return arr, diff
}

func partition(arr []int, left, right, diff int) ([]int, int, int) {
	v, i, j := arr[(left+right)/2], left, right-1
	for i <= j {
		for arr[i] < v {
			i++
		}
		for arr[j] > v {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		fmt.Println(arr)
		diff++
		i++
		j--
	}
	return arr, j + 1, diff
}
