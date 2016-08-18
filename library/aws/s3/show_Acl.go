package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	var bucket_name string
	var object_name string
	var types string
	flag.StringVar(&bucket_name, "BN", "", "-BN bucket_name")
	flag.StringVar(&object_name, "ON", "", "-ON object_name")
	flag.StringVar(&types, "T", "", "-T type")
	flag.Parse()

	if types == "" {
		fmt.Println("please input your type you want, example go run xx.go -T B or O\n")
	}
	svc := s3.New(session.New(&aws.Config{Region: aws.String("cn-north-1")}))
	switch types {
	case "B":
		getBucketAcl(svc, bucket_name, types)
	case "O":
		getObjectAcl(svc, bucket_name, object_name, types)
	default:
		fmt.Println("have no this type")
	}
}

func getBucketAcl(svc *s3.S3, bucket string, types string) {
	if bucket == "" || types != "B" {
		fmt.Println("you perhaps have a wrong type or have no bucketname")
		fmt.Println("get BucketAcl example: go run xx.go -T B -BN bucketName")
	}
	result, err := svc.GetBucketAcl(&s3.GetBucketAclInput{
		Bucket: &bucket,
	})
	if err != nil {
		log.Fatal("show Acl failed\n", err)
	}
	for _, acl := range result.Grants {
		fmt.Println("ACL:", *acl.Permission)
	}
}

func getObjectAcl(svc *s3.S3, bucket string, key string, types string) {
	if bucket == "" || key == "" || types != "O" {
		fmt.Println("you perhaps have a wrong type or have no bucketname or objectname")
		fmt.Println("get objectAcl example: go run xx.go -T O -BN bucketName -ON objectName")
	}
	result, err := svc.GetObjectAcl(&s3.GetObjectAclInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		log.Fatal("show Acl failed\n", err)
	}
	for _, acl := range result.Grants {
		fmt.Println("Acl: ", *acl.Permission)
	}
}
