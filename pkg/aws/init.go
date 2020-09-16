package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type awsClient struct {
	session *session.Session
}

func InitializeSessionAWS() (S3Store, error) {
	region := ""
	secretAccessKey := ""
	sessionToken := ""
	accessKeyID := ""
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKeyID,     // id
			secretAccessKey, // secret
			sessionToken),
	})
	return &awsClient{
		session: sess,
	}, err
}
