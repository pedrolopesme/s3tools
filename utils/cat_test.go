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

func TestCatFilesShouldNotRunWithoutFileNames(test *testing.T) {
	output := captureOutput(func() {
		CatFiles("some-bucket", []string{})
	})

	assert.Equal(test, "You must provide at least one file name\n", output)
}

func TestCatFilesWithoutBucket(test *testing.T) {
	output := captureOutput(func() {
		CatFiles("", []string{"file1"})
	})

	assert.Equal(test, "Bucket parameter cannot be blank\n", output)
}

func TestMatchFileWithoutPattern(test *testing.T) {
	output := captureOutput(func() {
		match("", "dummy-file")
	})

	assert.Equal(test, "Pattern parameter cannot be blank\n", output)
}

func TestMatchFileWithoutFileName(test *testing.T) {
	output := captureOutput(func() {
		match("dummy-pattern", "")
	})

	assert.Equal(test, "Filename parameter cannot be blank\n", output)
}

func TestMatchFileMatchingFileInBucketRootDir(test *testing.T) {
	matched := match("file1", "file1")
	assert.True(test, matched)
}

func TestMatchFileMatchingFileInBucketSubDir(test *testing.T) {
	matched := match("file1", "dir1/dir2/dir3/file1")
	assert.True(test, matched)
}

func TestMatchFileWithoutMatching(test *testing.T) {
	matched := match("file1", "not-file-1")
	assert.False(test, matched)
}

func TestFilterFilesWithNoFilesPattern(test *testing.T) {
	output := captureOutput(func() {
		filterFiles([]string{}, []s3File{newMockedFile([]byte("teste"))})
	})

	assert.Equal(test, "You must provide at least one file pattern\n", output)
}

func TestFilterFilesWithNoFiles(test *testing.T) {
	output := captureOutput(func() {
		filterFiles([]string{"dummy-patten"}, []s3File{})
	})

	assert.Equal(test, "You must provide at least one file from the bucket\n", output)
}

func TestFilterFilesWithNoMatches(test *testing.T) {
	patterns := []string{"dummy"}
	bucketFiles := []s3File{
		newMockedFile([]byte("test 1")),
		newMockedFile([]byte("test 2")),
		newMockedFile([]byte("test 3")),
	}

	filesFound, err := filterFiles(patterns, bucketFiles)
	assert.Nil(test, err)
	assert.Empty(test, filesFound)
}

func TestFilterFilesWithOneMatch(test *testing.T) {
	patterns := []string{"my-file"}
	bucketFiles := []s3File{
		newMockedFileWithPath([]byte("content 1"), "dummy 1"),
		newMockedFileWithPath([]byte("content 2"), "dummy 2"),
		newMockedFileWithPath([]byte("content 3"), "my-file"),
	}

	filesFound, err := filterFiles(patterns, bucketFiles)
	assert.Nil(test, err)
	assert.NotEmpty(test, filesFound)
	assert.Equal(test, len(filesFound), 1)
	assert.Equal(test, filesFound[0].GetPath(), bucketFiles[2].GetPath())
}

func TestFilterFilesWithMultipleMatches(test *testing.T) {
	patterns := []string{"my-file"}
	bucketFiles := []s3File{
		newMockedFileWithPath([]byte("content 1"), "dummy 1"),
		newMockedFileWithPath([]byte("content 2"), "dummy 2"),
		newMockedFileWithPath([]byte("content 3"), "my-file 1"),
		newMockedFileWithPath([]byte("content 4"), "my-file 2"),
		newMockedFileWithPath([]byte("content 5"), "my-file 3"),
	}

	filesFound, err := filterFiles(patterns, bucketFiles)
	assert.Nil(test, err)
	assert.NotEmpty(test, filesFound)
	assert.Equal(test, len(filesFound), 3)
	assert.Equal(test, filesFound[0].GetPath(), bucketFiles[2].GetPath())
	assert.Equal(test, filesFound[1].GetPath(), bucketFiles[3].GetPath())
	assert.Equal(test, filesFound[2].GetPath(), bucketFiles[4].GetPath())
}

func TestPrintFileContentWithNoFiles(test *testing.T) {
	output := captureOutput(func() {
		printFilesContent([]s3File{})
	})

	assert.Equal(test, "", output)
}

func TestPrintFileContentWithOneFile(test *testing.T) {
	output := captureOutput(func() {
		printFilesContent([]s3File{
			newMockedFile([]byte("content 1")),
		})
	})

	assert.Equal(test, "content 1\n", output)
}

func TestPrintFileContentWithMultipleFiles(test *testing.T) {
	output := captureOutput(func() {
		printFilesContent([]s3File{
			newMockedFile([]byte("content 1")),
			newMockedFile([]byte("content 2")),
			newMockedFile([]byte("content 3")),
		})
	})

	assert.Equal(test, "content 1\ncontent 2\ncontent 3\n", output)
}
