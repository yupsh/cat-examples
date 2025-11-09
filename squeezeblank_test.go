package cat_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/cat"
)

func ExampleCat_squeezeBlank() {
	// echo "Line 1\n\n\nLine 2\n\n\n\nLine 3" | cat -s
	gloo.MustRun(
		Cat(SqueezeBlank, strings.NewReader("Line 1\n\n\nLine 2\n\n\n\nLine 3")),
	)
	// Output:
	// Line 1
	//
	// Line 2
	//
	// Line 3
}
