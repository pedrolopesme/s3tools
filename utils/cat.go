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
	"fmt"
	"regexp"
)

// printFileContent tries to print the content of a
// files stored at a s3 bucket
// TODO implement and test
func printFileContent(file s3File) {
	fmt.Println("Printing file content", file)
}

// match tries to match a fileName using a regex pattern
func match(pattern string, fileName string) bool {
	if pattern == ""{
		fmt.Println("Pattern parameter cannot be blank")
		return false
	}
	if fileName == ""{
		fmt.Println("Filename parameter cannot be blank")
		return false
	}
	result, err := regexp.Match(pattern, []byte(fileName))
	if err != nil {
		fmt.Println("It was impossible to evaluate", fileName, "using pattern", pattern, "due to", err)
		return false
	}
	return result
}

// filterFiles gets all files from the bucket and filters all files
// whose name matches any of the given filename patterns
func filterFiles(filesPattern []string, bucketFiles []s3File) (files []s3File, err error) {
	if len(filesPattern) == 0 {
		fmt.Println("You must provide at least one file pattern")
	}

	if len(bucketFiles) == 0 {
		fmt.Println("You must provide at least one file from the bucket")
	}

	for _, bucketFile := range bucketFiles {
		filename := bucketFile.GetPath()
		for _, pattern := range filesPattern {
			if match(pattern, filename) {
				files = append(files, bucketFile)
			}
		}
	}
	return
}

// CatFiles will read all files matching their name to the user input
// and print the content to the standard output
func CatFiles(bucket string, filesPattern []string) {
	if len(bucket) == 0 {
		fmt.Println("Bucket parameter cannot be blank")
		return
	}
	if len(filesPattern) == 0 {
		fmt.Println("You must provide at least one file name")
		return
	}

	bucketFiles, err := ListObjects(bucket, "")
	if err != nil {
		fmt.Println("Something happened while retrieving filesPattern from bucket:", err)
		return
	}

	files, err := filterFiles(filesPattern, bucketFiles)
	if err != nil {
		fmt.Println("Something happened while filtering files:", err)
		return
	}

	for _, file := range files {
		printFileContent(file)
	}
	fmt.Println("CatFiles finished.")
}
