package main

import(
	"os"
	"fmt"
	"flag"
	"log"

	"github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func  main() {
        var bucket_name string
	var object_name string
	flag.StringVar(&bucket_name, "BN", "delault", "-BN bucketName")
	flag.StringVar(&object_name, "ON", "default", "-ON objectName")
	flag.Parse() 
	file, err := os.Create("download_file")
	if err != nil {
    	     log.Fatal("Failed to create file", err)
	}
	defer file.Close()

	downloader := s3manager.NewDownloader(session.New(&aws.Config{Region: aws.String("cn-north-1")}))
	numBytes, err := downloader.Download(file,&s3.GetObjectInput{
             Bucket: &bucket_name,
             Key:    &object_name,
        })
	if err != nil {
    	     fmt.Println("Failed to download file", err)
    	     return
	}

	fmt.Println("Downloaded file", file.Name(), numBytes, "bytes")

	
}
