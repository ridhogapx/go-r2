package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type R2Config struct {
	Bucket    string
	AccountID string
	Key       string
	Secret    string
}

type bucketBasic struct {
	Client *s3.Client
}

func (b *bucketBasic) UploadFile(bucketName string, objectKey string, filename string, filetype string) error {
	file, err := os.Open(filename)

	if err != nil {
		log.Println("Couldn't open file: ", err)
		return err
	}

	defer file.Close()
	_, err = b.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(objectKey),
		Body:        file,
		ContentType: aws.String(filetype),
	})

	if err != nil {
		log.Println("Couldn't upload file to S3: ", err)
		return err
	}

	return nil
}

func main() {
	// Initialize credentials
	r2 := R2Config{
		Bucket:    "dev-01",
		AccountID: "1b29a6e5d3b7451ee454d27dc2350700",
		Key:       "933b0ca2a33bb969d4885def8d2f95c7",
		Secret:    "ec9ffb554ca242384d58cf9a72105ca872a4b6aa3f67d26d1412770b32b7aa1c",
	}

	// Logging for testing purpose
	fmt.Println("THIS IS R2 Cloudflare")

	r2revolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", r2.AccountID),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2revolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(r2.Key, r2.Secret, "")),
		config.WithRegion("auto"),
	)

	if err != nil {
		log.Fatal("Failed to initialize config:", err)
	}

	client := s3.NewFromConfig(cfg)

	// Initialize BucketBasic
	basic := bucketBasic{
		Client: client,
	}

	// Upload file example-file.txt
	basic.UploadFile(r2.Bucket, "loli", "klee.png", "image/png")

	listObjects, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &r2.Bucket,
	})

	// Error handling when fetching list of objects
	if err != nil {
		log.Fatal("Failed to retrieve objects:", err)
	}

	// Logging for checking program is still working
	fmt.Println("R2 Service Cloudflare")

	// Iterate over object data
	for _, object := range listObjects.Contents {
		obj, _ := json.MarshalIndent(object, "", "\t")
		fmt.Println(string(obj))
	}
}
