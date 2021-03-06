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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListObjectsWithoutBucket(test *testing.T) {
	output := captureOutput(func() {
		ListObjects("", "some-path")
	})

	assert.Equal(test, "Bucket parameter cannot be blank\n", output)
}

func TestListObjectsWithNoObjects(test *testing.T) {
	bucket := newMockedS3Client([]string{})
	files, err := bucket.GetFiles()
	assert.Nil(test, err)
	assert.Empty(test, files)
}

func TestListObjectsWithOneObject(test *testing.T) {
	file1 := "dummy-file"
	bucket := newMockedS3Client([]string{file1})

	files, err := bucket.GetFiles()
	assert.Nil(test, err)
	assert.NotEmpty(test, files)
	assert.Equal(test, 1, len(files))
	assert.Equal(test, file1, files[0].GetPath())
}

func TestListObjectsWithMultipleObjects(test *testing.T) {
	f1 := "dummy-file 1"
	f2 := "dummy-file 2"
	f3 := "dummy-file 3"
	f4 := "dummy-file 4"
	f5 := "dummy-file 5"
	bucket := newMockedS3Client([]string{f1, f2, f3, f4, f5})

	files, err := bucket.GetFiles()
	assert.Nil(test, err)
	assert.NotEmpty(test, files)
	assert.Equal(test, 5, len(files))
	assert.Equal(test, f1, files[0].GetPath())
	assert.Equal(test, f2, files[1].GetPath())
	assert.Equal(test, f3, files[2].GetPath())
	assert.Equal(test, f4, files[3].GetPath())
	assert.Equal(test, f5, files[4].GetPath())
}
