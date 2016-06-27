// The MIT License (MIT)

// Copyright (c) 2015-2016 Christian Bergum Bergersen

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// CAUTION: This file in used by a unit test in file basicblock_test.go,
// changing the structure or logic in the program will with high possibility
// break the test and give false positive errors. Please DO NOT change this
// file unless you know what you are doing!
package main

import "fmt"

func main() { // BB #0 ending.
	number := 3

	switch number { // BB #1 ending.

	case 0: // BB #2 ending.
		fmt.Println("0")
	case 1: // BB #3 ending.
		fmt.Println("1")
		fmt.Println("1.a")
	case 2: // BB #4 ending.
		fmt.Println("2")
	case 3: // BB #5 ending.
		fmt.Println("3")
	case 4: // BB #6 ending.
		fmt.Println("4")
		return // BB #7 ending.
	default: // BB #8 ending.
		fmt.Printf("No match, number is %d!\n", number)
	}
} // BB #9 ending.
