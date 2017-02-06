package main

import(
     "fmt"
     "os"

     "github.com/aws/aws-sdk-go/aws"
     "github.com/aws/aws-sdk-go/aws/session"
     "github.com/aws/aws-sdk-go/service/s3"
)
const MB = 1048576
const KB = 1024
var total_size int64
func main() {
    //list_object
   svc := s3.New(session.New(&aws.Config{Region: aws.String("cn-north-1")}))
   i := 0
   err := svc.ListObjectsPages(&s3.ListObjectsInput{
   		Bucket: &os.Args[1], 
   	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
   		//fmt.Println("page,", i)
		//count := len(p.Contents)
		//fmt.Println("count:", count)
   		for _, obj := range p.Contents {
   			fmt.Println("Objects: ", *obj.Key)
			fmt.Printf("Objects_size: %d KB \n", (*obj.Size)/1024)
			total_size = total_size + *obj.Size
   		}
		if last == true {
			last_count := len(p.Contents)
			total := 1000 * i + last_count
			for _, obj := range p.Contents {
                        	fmt.Println("Objects: ", *obj.Key)
                        	fmt.Printf("Objects_size: %d KB \n", (*obj.Size)/KB)
                        	total_size = total_size + *obj.Size
                }
			fmt.Println("last page count:", last_count)
			fmt.Println("Total: ", total)
			fmt.Printf("total_size: %d MB", (total_size)/MB)
		}
		i++
   		return true
   })
   if err != nil {
   	fmt.Println("Failed to list objects", err)
   	return
   }
}
