package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBufferedContentWithEmptyBucket(test *testing.T) {
	file := newBufferedFile("", "some-path")
	buf, err := file.GetBufferedContent()
	assert.Empty(test, err)
	assert.Empty(test, buf)
}
