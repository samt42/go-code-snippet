package main

import (
	"fmt"
	"log"
	"flag"

	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)


func main() {
	var bucket_name string
	var object_name string
	var acl string
	flag.StringVar(&bucket_name, "BN", "", "-BN bucket_name")
	flag.StringVar(&object_name, "ON", "", "-ON object_name")
	flag.StringVar(&acl, "A", "", "-A acl")
	flag.Parse()
	if bucket_name == "" || object_name == "" || acl == "" {
		fmt.Println("Please input your information ,example go run xx.go -A acl -BN bucketName -ON objectName")
	}else{
		svc := s3.New(session.New(&aws.Config{Region: aws.String("cn-north-1")}))
		_, err := svc.PutObjectAcl(&s3.PutObjectAclInput{
				Bucket: &bucket_name,
				Key: &object_name,
				ACL: &acl,
		})
		if err != nil {
			log.Fatal("put objectAcl error", err)
		}
	}
	
}


