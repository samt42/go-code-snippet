package main

import (
	"fmt"
	"flag"
	"os"
	"log"

	"github.com/aws/aws-sdk-go/aws"
    	"github.com/aws/aws-sdk-go/aws/session"
    	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)


func main() {
	var bucket_name string
	var object_name string
	flag.StringVar(&bucket_name, "BN", "default", "-BN bucketName")
	flag.StringVar(&object_name, "ON", "default", "-ON objectName")
	flag.Parse()

	file, err := os.Open("upload_file")
	if err != nil{
		log.Fatal("Failed to open file", err)
	}

	uploader := s3manager.NewUploader(session.New(&aws.Config{Region: aws.String("cn-north-1")}))
	result ,err := uploader.Upload(&s3manager.UploadInput{
		Body: file,
		Bucket: &bucket_name,
		Key: &object_name,
	})

	if err != nil{
		log.Fatal("Failed to upload", err)
	}

	fmt.Println("Success to upload", result.Location)
}
