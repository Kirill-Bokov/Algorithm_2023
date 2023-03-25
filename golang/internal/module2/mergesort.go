package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var output *os.File

func Mergesort() {
	var err error
	output, err = os.OpenFile("output.txt", os.O_APPEND|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()
	reader := bufio.NewReader(input)
	str, err := reader.ReadString('\n')
	n := 0
	n, _ = strconv.Atoi(strings.TrimSuffix(strings.TrimSuffix(str, "\n"), "\r"))
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		str, _ := reader.ReadString(' ')
		arr[i], _ = strconv.Atoi(strings.TrimSuffix(strings.TrimSuffix(str, "\n"), " "))
	}
	input.Close()
	sorted := merge_sort(arr, 0, n)
	output.WriteString(strings.Trim(fmt.Sprint(sorted), "[]"))
	output.Close()
}

func merge(a []int, b []int, pos int) []int {
	i := 0
	j := 0
	k := 0
	arr := make([]int, len(a)+len(b))
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			arr[k] = a[i]
			i++
		} else {
			arr[k] = b[j]
			j++
		}
		k++
	}
	for ; i < len(a); i++ {
		arr[k] = a[i]
		k++
	}
	for ; j < len(b); j++ {
		arr[k] = b[j]
		k++
	}
	output.WriteString(fmt.Sprintf("%d %d %d %d\n", pos+1, pos+k, arr[0], arr[k-1]))
	return arr
}

func merge_sort(v []int, l int, r int) []int {
	if r-l < 2 {
		return []int{v[l]}
	}

	return merge(merge_sort(v, l, (r+l)/2), merge_sort(v, (r+l)/2, r), l)
}
