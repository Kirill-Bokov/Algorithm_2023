package module2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Radixsort() {
	input, _ := os.OpenFile("input.txt", os.O_RDONLY, 0666)
	defer input.Close()
	reader := bufio.NewReader(input)
	var line []string
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		str = strings.TrimSuffix(str, "\r")
		str = strings.TrimSuffix(str, "\n")
		line = append(line, str)
	}
	inputlength, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimSuffix(line[0], "\n"), "\r"))
	var inputarray []string
	for i := 1; i <= inputlength; i++ {
		inputarray = append(inputarray, line[i])
	}
	fmt.Println(radix(inputarray, 2, 9))

}

func radix(a []string, m, k int) []string {
	for i := m - 1; i >= 0; i-- {
		c := make([]int, (k + 1))
		b := make([]int, len(a))
		for j := 0; j < len(a); j++ {
			inta, _ := strconv.Atoi(string(a[j][i]))
			c[inta]++
		}
		count := 0
		for j := 0; j < len(c); j++ {
			t := c[j]
			c[j] = count
			count += t
		}
		for j := 0; j < len(a); j++ {
			inta, _ := strconv.Atoi(a[j])
			intaji, _ := strconv.Atoi(string(a[j][i]))
			b[c[intaji]] = inta
			c[intaji]++
		}
		a = convert(b)
	}
	return a
}

func convert(bi []int) []string {
	bs := make([]string, len(bi))
	for i, num := range bi {
		bs[i] = strconv.Itoa(num)
	}
	return bs
}
