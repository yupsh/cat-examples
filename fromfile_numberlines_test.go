package cat_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/cat"
)

func ExampleCat_fromFile_numberLines() {
	// cat -n testdata/simple.txt
	yup.MustRun(
		Cat(NumberLines, yup.File("testdata/simple.txt")),
	)
	// Output:
	//      1	Line one
	//      2	Line two
	//      3	Line three
}

