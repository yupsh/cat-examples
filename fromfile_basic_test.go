package cat_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/cat"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleCat_fromFile_basic() {
	// cat testdata/simple.txt
	gloo.MustRun(
		Cat(gloo.File("testdata/simple.txt")),
	)
	// Output:
	// Line one
	// Line two
	// Line three
}
