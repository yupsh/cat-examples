package cat_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/cat"
)

func ExampleCat_basic() {
	// echo "Hello World\nThis is a test" | cat
	yup.MustRun(
		Cat(strings.NewReader("Hello World\nThis is a test")),
	)
	// Output:
	// Hello World
	// This is a test
}

