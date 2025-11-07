package cat_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/cat"
)

func ExampleCat_showEnds() {
	// echo "First line\nSecond line" | cat -E
	yup.MustRun(
		Cat(ShowEnds, strings.NewReader("First line\nSecond line")),
	)
	// Output:
	// First line$
	// Second line$
}

