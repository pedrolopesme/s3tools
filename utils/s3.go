package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"fmt"
)

type s3BufferedFile struct {
	Bucket string
	Path   string
}

// GetBucket returns file path
func (f s3BufferedFile) GetBucket() string {
	return f.Bucket
}

// GetPath returns file path
func (f s3BufferedFile) GetPath() string {
	return f.Path
}

// GetBufferedContent creates a S3 session and reads the content
// of a file writing it into a buffer.
func (f s3BufferedFile) GetBufferedContent() (buf []byte, err error) {
	if f.Bucket == "" {
		fmt.Println("Aborting due to empty bucket")
		return
	}

	sess := session.Must(session.NewSession())
	downloader := s3manager.NewDownloader(sess)
	awsBuffer := &aws.WriteAtBuffer{}

	_, err = downloader.Download(awsBuffer,
		&s3.GetObjectInput{
			Bucket: aws.String(f.Bucket),
			Key:    aws.String(f.Path),
		})

	if err != nil {
		return
	}

	buf = awsBuffer.Bytes()
	return
}

// NewBufferedFile returns a concrete type of s3BufferedFile
// properly initialized. Following recommendations found at
// https://github.com/golang/go/wiki/CodeReviewComments#interfaces
func NewBufferedFile(bucket string, path string) s3BufferedFile {
	return s3BufferedFile{
		Bucket: bucket,
		Path:   path,
	}
}
