package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetBufferedContentWithEmptyBucket(test *testing.T) {
	file := NewBufferedFile("", "some-path")
	buf, err := file.GetBufferedContent()
	assert.Empty(test, err)
	assert.Empty(test, buf)
}