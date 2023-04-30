package main

import (
    "fmt"
    "os"
    "io"
    
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

func main() {
    // Create a new AWS session
    sess := session.Must(session.NewSession(&aws.Config{
    Region: aws.String("us-west-2"),
}))


    // Create a new S3 client
    svc := s3.New(sess)

    // Set the S3 bucket and object key
    bucket := "keploy-enterprise"
    key := "keploy/io/0.1.0/enterprise-server_windows_amd64.tar.gz"

    // Create a new GetObjectInput
    input := &s3.GetObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
    }

    // Use the S3 client to download the object
    result, err := svc.GetObject(input)
    if err != nil {
        fmt.Println("Error downloading object:", err)
        return
    }
    defer result.Body.Close()

    // Create a new file to save the object content to
    file, err := os.Create("enterprise-server_windows_amd64.tar.gz")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    // Save the object content to the file
    _, err = io.Copy(file, result.Body)
    if err != nil {
        fmt.Println("Error saving object content to file:", err)
        return
    }

    fmt.Println("File downloaded successfully!")
}

// package main

// import (
// 	"fmt"
// 	"regexp"
// 	"strings"
// )

// func main() {
// 	email := "Cample@xample.com"
// 	email = strings.ToLower(email)
// 	// regular expression pattern to match email addresses
// 	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	
// 	// compile the pattern into a regular expression object
// 	regex := regexp.MustCompile(pattern)
	
// 	// check if the email matches the pattern
// 	if regex.MatchString(email) {
// 		fmt.Println("Email is valid")
// 	} else {
// 		fmt.Println("Email is invalid")
// 	}
// }