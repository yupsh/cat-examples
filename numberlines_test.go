package cat_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/cat"
)

func ExampleCat_numberLines() {
	// echo "Line one\nLine two\nLine three" | cat -n
	gloo.MustRun(
		Cat(NumberLines, strings.NewReader("Line one\nLine two\nLine three")),
	)
	// Output:
	//      1	Line one
	//      2	Line two
	//      3	Line three
}
