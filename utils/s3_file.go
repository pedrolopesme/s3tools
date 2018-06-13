package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3BufferedFile struct {
	Bucket string
	Path   string
}

// GetBucket returns file path
func (f S3BufferedFile) GetBucket() string {
	return f.Bucket
}

// GetPath returns file path
func (f S3BufferedFile) GetPath() string {
	return f.Path
}

// GetBufferedContent creates a S3 session and reads the content
// of a file writing it into a buffer.
func (f S3BufferedFile) GetBufferedContent() (buf []byte, err error) {
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

// NewBufferedFile returns a concrete type of S3BufferedFile
// properly initialized. Following recommendations found at
// https://github.com/golang/go/wiki/CodeReviewComments#interfaces
func NewBufferedFile(bucket string, path string) S3BufferedFile {
	return S3BufferedFile{
		Bucket: bucket,
		Path:   path,
	}
}
