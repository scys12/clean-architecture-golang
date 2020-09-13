package aws

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type FileForm struct {
	session    *session.Session
	file       multipart.File
	fileHeader *multipart.FileHeader
}

func InitializeSessionAWS() (*session.Session, error) {
	region := ""
	secretAccessKey := ""
	sessionToken := ""
	accessKeyID := ""
	return session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKeyID,     // id
			secretAccessKey, // secret
			sessionToken),
	})
}
