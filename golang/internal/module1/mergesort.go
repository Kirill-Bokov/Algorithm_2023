package module1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Mergesort() {
	reader := bufio.NewReader(os.Stdin)
	arrlen, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	intarrlen, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimSuffix(arrlen, "\n"), "\r"))
	line, err := reader.ReadString('\n')
	strarr := strings.Split(strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r"), " ")
	if (len(strarr) != intarrlen) || intarrlen > 10^5 || intarrlen < 1 {
		return
	}
	var arr []int
	for l := 0; l < len(strarr); l++ {
		inti, _ := strconv.Atoi(strarr[l])
		arr = append(arr, inti)
	}
	sorted := merge_sort(arr)
	fmt.Println(sorted)
}

func merge(a, b []int) []int {
	i, j, k := 0, 0, 0
	c := make([]int, len(a)+len(b))
	fmt.Println(len(c), a, b)
	for k < len(c) {
		if (j == len(b)) || (i < len(a) && a[i] <= b[j]) {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
		k++
		fmt.Println(c)
	}
	return c
}

func merge_sort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	mid := len(a) / 2
	return merge(merge_sort(a[:mid]), merge_sort(a[mid:]))
}
