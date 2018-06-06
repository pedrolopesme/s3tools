// MIT License
//
// Copyright (c) 2018 Pedro Mendes
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

// Grep identifies occurrences of a given string or pattern
// on files stored in a S3 bucket
func Grep(bucket string, pattern string) {
	if bucket == "" {
		fmt.Println("Bucket parameter cannot be blank")
		return
	}

	if pattern == "" {
		fmt.Println("Pattern parameter cannot be blank")
		return
	}

	files := ListObjects(bucket)

	for _, file := range files {
		buff, err := file.GetBufferedContent()
		if err != nil {
			continue
		}

		// TODO Extract it to a new func
		scanner := bufio.NewScanner(bytes.NewReader(buff))
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, pattern) {
				// TODO create a printer func
				fmt.Println(file.Path + " : " + scanner.Text())
			}
		}
	}

	fmt.Println("Grep finished.")
}
