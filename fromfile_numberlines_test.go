package cat_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/cat"
)

func ExampleCat_fromFile_numberLines() {
	// cat -n testdata/simple.txt
	gloo.MustRun(
		Cat(NumberLines, gloo.File("testdata/simple.txt")),
	)
	// Output:
	//      1	Line one
	//      2	Line two
	//      3	Line three
}
