package main

import(  
     "log"

     "github.com/aws/aws-sdk-go/aws"
     "github.com/aws/aws-sdk-go/aws/session"
     "github.com/aws/aws-sdk-go/service/s3"
)


func main() {
    //list_bucket	
    svc := s3.New(session.New(&aws.Config{Region: aws.String("cn-north-1")}))
    result, err := svc.ListBuckets(&s3.ListBucketsInput{})
    if err != nil {
        log.Println("Failed to list buckets", err)
        return
    }

    log.Println("Buckets:")
    for _, bucket := range result.Buckets {
        log.Printf("%s : %s\n", aws.StringValue(bucket.Name), bucket.CreationDate)
    }
}
