package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ListObjects retrieves a list of files from a given
// bucket and return a S3File list.
func ListObjects(bucket string) (files []S3File) {
	if len(bucket) < 2 {
		fmt.Println("you must specify a bucket")
		return nil
	}

	sess := session.Must(session.NewSession())
	svc := s3.New(sess)
	err := svc.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: &bucket,

	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		for _, s3Object := range p.Contents {
			file := S3BufferedFile{Bucket: bucket, Path: string(*s3Object.Key)}
			files = append(files, file)
		}
		return true
	})

	if err != nil {
		fmt.Println("failed to list objects", err)
		return
	}
	return
}
