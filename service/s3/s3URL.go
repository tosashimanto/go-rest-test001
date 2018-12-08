package s3

import (
	"fmt"
	"github.com/tosashimanto/go-rest-test001/service"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3署名付きURLを提供する
type AWSCredential struct {
	AccessKey       string
	SecretAccessKey string
	Token           string
}

func createAwsCredential() *credentials.Credentials {
	creds := credentials.NewStaticCredentials(
		service.BUCKETEER_AWS_ACCESS_KEY_ID,
		service.BUCKETEER_AWS_SECRET_ACCESS_KEY,
		"",
	)
	return creds
}

func createAwsConfig(creds *credentials.Credentials) *aws.Config {
	s3Config := &aws.Config{
		Credentials: creds,
		// Endpoint:         aws.String("http://localhost:8083"),
		Region: aws.String(service.BUCKETEER_AWS_REGION),
		// DisableSSL:       aws.Bool(true),
		// S3DisableContentMD5Validation :  aws.Bool(false),
		// S3ForcePathStyle: aws.Bool(true),
	}
	return s3Config
}

func NewPutPreSignedS3URL(objectKey string) (string, error) {

	creds := createAwsCredential()
	newSession := session.New()
	s3Config := createAwsConfig(creds)

	// Create an S3 service object in the default region.
	s3Client := s3.New(newSession, s3Config)
	params := &s3.PutObjectInput{
		Bucket: aws.String(service.BUCKETEER_BUCKET_NAME),
		Key:    aws.String(objectKey),
	}
	req, _ := s3Client.PutObjectRequest(params)
	url, err := req.Presign(10 * time.Minute)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(url)
	return url, err
}

func NewGetPreSignedS3URL(objectKey string) (string, error) {

	creds := createAwsCredential()
	newSession := session.New()
	s3Config := createAwsConfig(creds)
	// Create an S3 service object in the default region.
	s3Client := s3.New(newSession, s3Config)
	params := &s3.GetObjectInput{
		Bucket: aws.String(service.BUCKETEER_BUCKET_NAME),
		Key:    aws.String(objectKey),
	}
	req, _ := s3Client.GetObjectRequest(params)
	url, err := req.Presign(10 * time.Minute)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(url)
	return url, err
}
