package cat_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/cat"
)

func ExampleCat_multipleFlags() {
	// echo "First line\n\n\nSecond line" | cat -n -E -s
	gloo.MustRun(
		Cat(NumberLines, ShowEnds, SqueezeBlank, strings.NewReader("First line\n\n\nSecond line")),
	)
	// Output:
	//      1	First line$
	//      2	$
	//      3	Second line$
}
