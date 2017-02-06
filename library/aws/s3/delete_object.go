package main

import (
    "log"
    "fmt"
    "flag"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)


func main() {
	var bucket_name string
	var object_name string
	flag.StringVar(&bucket_name, "BN", "default", "-BN bucketName")
	flag.StringVar(&object_name, "ON", "default", "-ON objectName")
	flag.Parse()
	
	svc := s3.New(session.New(&aws.Config{Region: aws.String("cn-north-1")}))
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
			Bucket: &bucket_name,
			Key: &object_name,
	})

	if err != nil{
		log.Fatal("Delete object fail", err)
		return
	}
	fmt.Println("deletet object success")
	//fmt.Print("~~~~~%s\n", &result.RequestCharged)
	/*if result.DeleteMarker == true {
		fmt.Println("delete object success")
	}*/
	
}
