package utils

import (
	"bytes"
	"io"
	"os"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Captures the output by replacing the
// standard stdout with a buffer and then restore it again.
// Returns the buffer contents as String
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// mockedS3BufferedFile allows to mock an S3BufferedFile,
// storing a given value to be returned during tests.
type mockedS3BufferedFile struct {
	returnValue []byte
	s3BufferedFile
}

// GetBufferedContent returns a given value of type
// []byte from mockedS3BufferedFile
func (f mockedS3BufferedFile) GetBufferedContent() ([]byte, error) {
	return f.returnValue, nil
}

// newMockedFile builds a mockedS3File, injecting a
// value of type []byte to be used during tests.
func newMockedFile(returnedValue []byte) mockedS3BufferedFile {
	return mockedS3BufferedFile{
		returnValue: returnedValue,
		s3BufferedFile: s3BufferedFile{
			Bucket: "test",
			Path:   "test",
		},
	}
}

// mockS3Client allows us to replace S3 default implementation
type mockS3Client struct {
	s3iface.S3API
}

// mockS3Client allows us to replace S3 default implementation
func (m *mockS3Client) ListObjectsPages(*s3.ListObjectsInput, func(*s3.ListObjectsOutput, bool) bool) error {
	return nil
}

