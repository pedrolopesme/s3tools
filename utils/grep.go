//MIT License
//
//Copyright (c) 2018 Pedro Mendes
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Grep identifies occurrences of a given string or pattern
// on files stored in a S3 bucket
func Grep(bucket string) {
	files := listObjects(bucket)
	fmt.Println(files)
}

func listObjects(bucket string) (files []string) {
	if len(bucket) < 2 {
		fmt.Println("you must specify a bucket")
		return
	}

	sess := session.Must(session.NewSession())

	svc := s3.New(sess)

	i := 0
	err := svc.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: &bucket,
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		fmt.Println("Page,", i)
		i++

		for _, obj := range p.Contents {
			files = append(files, string(*obj.Key))
		}
		return true
	})

	if err != nil {
		fmt.Println("failed to list objects", err)
		return
	}
	return
}