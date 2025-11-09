package cat_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/cat"
)

func ExampleCat_basic() {
	// echo "Hello World\nThis is a test" | cat
	gloo.MustRun(
		Cat(strings.NewReader("Hello World\nThis is a test")),
	)
	// Output:
	// Hello World
	// This is a test
}
