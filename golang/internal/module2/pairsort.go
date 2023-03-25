package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Pairsort() {
	reader := bufio.NewReader(os.Stdin)
	arrnum, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	intarrnum, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimSuffix(arrnum, "\n"), "\r"))
	var darr [][]int
	for h := 0; h < intarrnum; h++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		strarrns := strings.Split(strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r"), " ")
		arg1, _ := strconv.Atoi(strarrns[0])
		arg2, _ := strconv.Atoi(strarrns[1])
		darr = append(darr, []int{arg1, arg2})
	}
	for j := 0; j < len(darr)-1; j++ {
		for k := 0; k < len(darr)-j-1; k++ {
			if darr[k][1] < darr[k+1][1] {
				darr[k], darr[k+1] = darr[k+1], darr[k]
			}
		}
	}
	for j := 0; j < len(darr)-1; j++ {
		for k := 0; k < len(darr)-j-1; k++ {
			if (darr[k][0] > darr[k+1][0]) && (darr[k][1] == darr[k+1][1]) {
				darr[k], darr[k+1] = darr[k+1], darr[k]
			}
		}
	}
	for l := 0; l < len(darr); l++ {
		fmt.Println(strconv.Itoa(darr[l][0]) + " " + strconv.Itoa(darr[l][1]))
	}
}
