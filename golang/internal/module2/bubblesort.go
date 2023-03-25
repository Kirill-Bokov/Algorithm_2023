package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Bubblesort() {
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
	for i := 0; i < len(strarr); i++ {
		inti, _ := strconv.Atoi(strarr[i])
		arr = append(arr, inti)
	}
	nopermutation := true
	for j := 0; j < len(arr)-1; j++ {
		for k := 0; k < len(arr)-j-1; k++ {
			if arr[k] > arr[k+1] {
				arr[k], arr[k+1] = arr[k+1], arr[k]
				for l := 0; l < len(arr)-1; l++ {
					fmt.Print(strconv.Itoa(arr[l]) + " ")
				}
				fmt.Println(arr[len(arr)-1])
				nopermutation = false
			}
		}
	}
	if nopermutation {
		fmt.Println(0)
	}
}
