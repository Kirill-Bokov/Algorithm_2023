package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Insertionsort() {
	reader := bufio.NewReader(os.Stdin)
	arrlen, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	intarrlen, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimSuffix(arrlen, "\n"), "\r"))
	line, err := reader.ReadString('\n')
	strarr := strings.Split(strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r"), " ")
	if (len(strarr) != intarrlen) || intarrlen > 1000 || intarrlen < 1 {
		return
	}
	var arr []int
	for k := 0; k < len(strarr); k++ {
		inti, _ := strconv.Atoi(strarr[k])
		arr = append(arr, inti)
	}
	fmt.Println(arr)
	for i := 1; i < len(arr); i++ {
		j, x := i-1, arr[i]
		for j >= 0 && arr[j] > x {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = x
	}
	fmt.Println(arr)
}
