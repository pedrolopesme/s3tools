package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type Bucket struct {
	Client s3iface.S3API
	Name   string
	Path   string
}

func (bucket *Bucket) GetFiles() (files []s3File, err error) {
	listObjectsInput := &s3.ListObjectsInput{Bucket: &bucket.Name}
	if bucket.Path != "" {
		listObjectsInput.Prefix = aws.String(bucket.Path)
	}

	fnCallback := func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		for _, s3Object := range p.Contents {
			file := newBufferedFile(bucket.Name, string(*s3Object.Key))
			files = append(files, file)
		}
		return true
	}

	err = bucket.Client.ListObjectsPages(listObjectsInput, fnCallback)
	if err != nil {
		fmt.Println("failed to list objects", err)
		return
	}
	return
}

func NewBucket(client s3iface.S3API, bucketName string, path string) (*Bucket) {
	return &Bucket{
		Client: client,
		Name:   bucketName,
		Path:   path}
}

// ListObjects retrieves a list of files from a given
// bucket and return a s3File list.
func ListObjects(bucketName string, path string) (files []s3File, err error) {
	if bucketName == "" {
		fmt.Println("Bucket parameter cannot be blank")
		return
	}

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("It was impossible to create an AWS session. Aborting")
	}

	return NewBucket(s3.New(sess), bucketName, path).GetFiles()
}
