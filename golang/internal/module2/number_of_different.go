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
	if (len(strarr) != intarrlen) || intarrlen > 100000 || intarrlen <= 1 {
		return
	}
	var arr []int
	for i := 0; i < len(strarr); i++ {
		inti, _ := strconv.Atoi(strarr[i])
		arr = append(arr, inti)
	}
	fmt.Println(countUnique(arr))
}
func countUnique(arr []int) int {
	unique := make(map[int]bool)
	for _, num := range arr {
		unique[num] = true
	}
	return len(unique)
}
