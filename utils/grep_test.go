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

func TestGrepFilesOnEmptyBucket(test *testing.T) {
	output := captureOutput(func() {
		GrepFiles("", "some-pattern", "")
	})

	assert.Equal(test, "Bucket parameter cannot be blank\n", output)
}

func TestGrepFilesForEmptyPattern(test *testing.T) {
	output := captureOutput(func() {
		GrepFiles("some-bucket", "", "")
	})

	assert.Equal(test, "Pattern parameter cannot be blank\n", output)
}

func TestGrepFileForNilFile(test *testing.T) {
	output := captureOutput(func() {
		GrepFile(nil, "some-pattern")
	})

	assert.Equal(test, "File parameter cannot be blank\n", output)
}

func TestGrepFileForEmptyPattern(test *testing.T) {
	output := captureOutput(func() {
		GrepFile(s3BufferedFile{}, "")
	})

	assert.Equal(test, "Pattern parameter cannot be blank\n", output)
}

func TestGrepFileWithNilBuffer(test *testing.T) {
	output := captureOutput(func() {
		GrepFile(newMockedFile(nil), "some-pattern")
	})

	assert.Equal(test, "", output)
}

func TestGrepWithoutMatching(test *testing.T) {
	content := []byte("dummy value")
	output := captureOutput(func() {
		GrepFile(newMockedFile(content), "some-pattern")
	})
	assert.Equal(test, "", output)
}

func TestGrepWithMatching(test *testing.T) {
	files := newMockedFile([]byte("dummy value 1"))

	output := captureOutput(func() {
		GrepFile(files, "dummy")
	})
	assert.Equal(test, "\x1b[1;36mtest\x1b[0m : \x1b[1;31mdummy\x1b[0m value 1\n", output)
}

func TestGrepWithMultipleMatchingAtTheSameLine(test *testing.T) {
	files := newMockedFile([]byte("dummy value 1, dummy value 2, dummy value 3"))

	output := captureOutput(func() {
		GrepFile(files, "dummy")
	})
	assert.Equal(test, "\x1b[1;36mtest\x1b[0m : \x1b[1;31mdummy\x1b[0m value 1, \x1b[1;31mdummy\x1b[0m value 2, \x1b[1;31mdummy\x1b[0m value 3\n", output)
}

func TestGrepWithWildcardAtTheEndOfPatternWithAMatch(test *testing.T) {
	files := newMockedFile([]byte("dummy value 1"))

	output := captureOutput(func() {
		GrepFile(files, "dummy*")
	})
	assert.Equal(test, "\x1b[1;36mtest\x1b[0m : \x1b[1;31mdummy\x1b[0m value 1\n", output)
}

func TestGrepWithWildcardAtTheEndOfPatternWithoutAMatch(test *testing.T) {
	files := newMockedFile([]byte("dummy value 1"))

	output := captureOutput(func() {
		GrepFile(files, "whatever*")
	})
	assert.Equal(test, "", output)
}

func TestGrepWithWildcardAtTheBeginningOfPatternWithAMatch(test *testing.T) {
	files := newMockedFile([]byte("dummy value 1"))

	output := captureOutput(func() {
		GrepFile(files, "*value")
	})
	assert.Equal(test, "\x1b[1;36mtest\x1b[0m : dummy \x1b[1;31mvalu\x1b[0me 1\n", output)
}

func TestGrepWithWildcardAtTheBeginningOfPatternWithoutAMatch(test *testing.T) {
	files := newMockedFile([]byte("dummy value 1"))

	output := captureOutput(func() {
		GrepFile(files, "*whatever")
	})
	assert.Equal(test, "", output)
}

func TestColorizedGrepWithOneMatching(test *testing.T) {
	files := newMockedFile([]byte("dummy value 1"))

	output := captureOutput(func() {
		GrepFile(files, "dummy")
	})
	assert.Equal(test, "\x1b[1;36mtest\x1b[0m : \x1b[1;31mdummy\x1b[0m value 1\n", output)
}

func TestColorizedGrepWithMultipleMatching(test *testing.T) {
	files := newMockedFile([]byte("dummy value 1, dummy value 2, dummy value 3"))

	output := captureOutput(func() {
		GrepFile(files, "dummy")
	})
	assert.Equal(test, "\x1b[1;36mtest\x1b[0m : \x1b[1;31mdummy\x1b[0m value 1, \x1b[1;31mdummy\x1b[0m value 2, \x1b[1;31mdummy\x1b[0m value 3\n", output)
}
