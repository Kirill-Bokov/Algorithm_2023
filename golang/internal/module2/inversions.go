package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Inversions() {
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
	var numinv int
	arr, numinv = merge_sort2(arr, 0)
	fmt.Println(strconv.Itoa(numinv))
}
func merge2(left, right []int, numinvl, numinvr int) ([]int, int) {
	i, j, k := 0, 0, 0
	c := make([]int, len(left)+len(right))
	numinv := numinvl + numinvr
	for k < (len(c)) {
		if (j == len(right)) || (i < len(left) && (left[i] <= right[j])) {
			c[k] = left[i]
			i++
		} else {
			c[k] = right[j]
			j++
			numinv = numinv + len(left) - i
		}
		k++
	}
	return c, numinv
}

func merge_sort2(arr []int, numinv int) ([]int, int) {
	if len(arr) == 1 {
		return arr, numinv
	}
	mid := len(arr) / 2
	l, numinvl := merge_sort2(arr[:mid], numinv)
	r, numinvr := merge_sort2(arr[mid:], numinv)
	return merge2(l, r, numinvl, numinvr)
}
