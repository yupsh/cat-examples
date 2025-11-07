package cat_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/cat"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleCat_fromFile_basic() {
	// cat testdata/simple.txt
	yup.MustRun(
		Cat(yup.File("testdata/simple.txt")),
	)
	// Output:
	// Line one
	// Line two
	// Line three
}

