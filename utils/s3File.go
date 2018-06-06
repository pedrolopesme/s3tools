package utils

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3File struct {
	Bucket string
	Path   string
}

func (f *S3File) GetBufferedContent() (buf []byte, err error) {

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
