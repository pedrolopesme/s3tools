// MIT License
//
// Copyright (c) 2018 Pedro Mendes
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	log "github.com/sirupsen/logrus"
	"github.com/aws/aws-sdk-go/service/s3"
	"strings"
)

// Grep identifies occurrences of a given string or pattern
// on files stored in a S3 bucket
func Grep(bucket string, pattern string) {
	if bucket == "" {
		fmt.Println("Bucket parameter cannot be blank")
		return
	}

	if pattern == "" {
		fmt.Println("Pattern parameter cannot be blank")
		return
	}

	files := ListObjects(bucket)
	fmt.Println(files)

	sess := session.Must(session.NewSession())
	downloader := s3manager.NewDownloader(sess)
	buff := &aws.WriteAtBuffer{}

	for _, file := range files {
		log.Info("Searching in " + file)

		_, err := downloader.Download(buff,
			&s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(file),
			})

		if err != nil {
			log.Error(err)
		}

		scanner := bufio.NewScanner(bytes.NewReader(buff.Bytes()))
		for scanner.Scan() {
			line := scanner.Text()

			if strings.Contains(line, pattern) {
				log.Info("Match : " + scanner.Text())
			}
		}
	}


	log.Info("Finished")
}
