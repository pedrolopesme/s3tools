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
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGrepFilesOnEmptyBucket(test *testing.T) {
	output := captureOutput(func(){
		GrepFiles("", "some-pattern")
	})

	assert.Equal(test,"Bucket parameter cannot be blank\n", output)
}

func TestGrepFilesForEmptyPattern(test *testing.T) {
	output := captureOutput(func(){
		GrepFiles("some-bucket", "")
	})

	assert.Equal(test,"Pattern parameter cannot be blank\n", output)
}

func TestGrepFileForNilFile(test *testing.T) {
	output := captureOutput(func(){
		GrepFile(nil, "some-pattern")
	})

	assert.Equal(test,"File parameter cannot be blank\n", output)
}

func TestGrepFileForEmptyPattern(test *testing.T) {
	output := captureOutput(func(){
		GrepFile(S3BufferedFile{}, "")
	})

	assert.Equal(test,"Pattern parameter cannot be blank\n", output)
}

func TestGrepFileWithNilBuffer(test *testing.T) {
	output := captureOutput(func(){
		GrepFile(NewMockedFile(nil), "some-pattern")
	})

	assert.Equal(test,"", output)
}

func TestGrepWithoutMatching(test *testing.T) {
	content := []byte("dummy value")
	output := captureOutput(func(){
		GrepFile(NewMockedFile(content), "some-pattern")
	})
	assert.Equal(test,"", output)
}

func TestGrepWithMatching(test *testing.T) {
	files := NewMockedFile([]byte("dummy value 1"))

	output := captureOutput(func(){
		GrepFile(files, "dummy")
	})
	assert.Equal(test,"test : dummy value 1\n", output)
}

func TestGrepWithMultipleMatchingAtTheSameLine(test *testing.T) {
	files := NewMockedFile([]byte("dummy value 1, dummy value 2, dummy value 3"))

	output := captureOutput(func(){
		GrepFile(files, "dummy")
	})
	assert.Equal(test,"test : dummy value 1, dummy value 2, dummy value 3\n", output)
}
