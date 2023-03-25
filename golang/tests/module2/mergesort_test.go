package module2_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/Kirill-Bokov/Algorithm_2023/golang/internal/module2"
	"github.com/stretchr/testify/assert"
)

func TestMergesort(t *testing.T) {
	assert := assert.New(t)
	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	t.Run("Case: 1; 1", func(t *testing.T) {
		input, err := os.Create("input.txt")
		if err != nil {
			os.Exit(1)
		}
		defer input.Close()
		input.WriteString(`1
		1`)
		module2.Mergesort()
		output, err := ioutil.ReadFile("output.txt")
		assert.Equal(string(output), `1`)
	})
	t.Run("Case: 2; 3 1", func(t *testing.T) {
		input, err := os.Create("input.txt")
		if err != nil {
			os.Exit(1)
		}
		defer input.Close()
		input.WriteString(`2
		3 1`)
		module2.Mergesort()
		output, err := ioutil.ReadFile("output.txt")
		assert.Equal(string(output), `1 2 1 3
		1 3`)
	})
	t.Run("Case: 5; 5 4 3 2 1", func(t *testing.T) {
		input, err := os.Create("input.txt")
		if err != nil {
			os.Exit(1)
		}
		defer input.Close()
		input.WriteString("5\n5 4 3 2 1")
		module2.Mergesort()
		output, err := ioutil.ReadFile("output.txt")
		assert.Equal(string(output), "1 2 4 5\n4 5 1 2\n3 5 1 3\n1 5 1 5\n1 2 3 4 5")
	})
}
