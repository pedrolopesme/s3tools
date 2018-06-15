package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

// ListObjects retrieves a list of files from a given
// bucket and return a s3File list.
func ListObjects(bucket string, path string) (files []s3File) {
	if len(bucket) < 2 {
		fmt.Println("you must specify a bucket")
		return nil
	}

	sess := session.Must(session.NewSession())
	svc := s3.New(sess)

	listObjectsInput := &s3.ListObjectsInput{ Bucket: &bucket }
	if path != "" {
		listObjectsInput.Prefix = aws.String(path)
	}

	fnCallback := func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		for _, s3Object := range p.Contents {
			file := NewBufferedFile(bucket, string(*s3Object.Key))
			files = append(files, file)
		}
		return true
	}

	err := svc.ListObjectsPages(listObjectsInput, fnCallback)
	if err != nil {
		fmt.Println("failed to list objects", err)
		return
	}
	return
}
