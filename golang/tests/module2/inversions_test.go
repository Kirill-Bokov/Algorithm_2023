package module2_test

import (
	"io"
	"os"
	"testing"

	"github.com/Kirill-Bokov/Algorithm_2023/golang/internal/module2"
	"github.com/Kirill-Bokov/Algorithm_2023/golang/tests/helpers"
	"github.com/stretchr/testify/assert"
)

func TestInversions(t *testing.T) {
	assert := assert.New(t)

	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	t.Run("Case: 1; 1", func(t *testing.T) {
		r, w := helpers.Replacer(`1
		1
`, t)
		os.Stdin = r
		os.Stdout = w

		module2.Inversions()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`0
`, string(out))
	})
	t.Run("Case: 2; 3 1", func(t *testing.T) {
		r, w := helpers.Replacer(`2
		3 1
`, t)
		os.Stdin = r
		os.Stdout = w

		module2.Inversions()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`1
`, string(out))
	})
	t.Run("Case: 5; 5 4 3 2 1", func(t *testing.T) {
		r, w := helpers.Replacer(`5
		5 4 3 2 1
`, t)
		os.Stdin = r
		os.Stdout = w

		module2.Inversions()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(`10 
`, string(out))
	})
}
