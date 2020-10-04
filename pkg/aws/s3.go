package aws

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	bucket = "sea-marketplace"
)

type FileParam struct {
	FileHeader *multipart.FileHeader
	UserID     primitive.ObjectID
	FolderName string
	FileURL    string
}

type S3Store interface {
	UploadFileToS3(FileParam) (string, error)
	setUploadFileName(FileParam) string
}

func (awsS3 *awsClient) setUploadFileName(fileParam FileParam) string {
	var tempFileName string
	if fileParam.FileURL != "" {
		tempFileName = strings.Split(fileParam.FileURL, "s3.amazonaws.com/")[1]
	} else {
		tempFileName = fmt.Sprintf("%v/%v-%v%v", fileParam.FolderName, fileParam.UserID.Hex(), uuid.New().String(), filepath.Ext(fileParam.FileHeader.Filename))
	}
	return tempFileName
}

func (awsS3 *awsClient) UploadFileToS3(fileParam FileParam) (string, error) {
	src, err := fileParam.FileHeader.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	size := fileParam.FileHeader.Size
	buffer := make([]byte, size)
	_, err = src.Read(buffer)
	if err != nil {
		return "", err
	}
	tempFileName := awsS3.setUploadFileName(fileParam)
	_, err = s3.New(awsS3.session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://%v.s3.amazonaws.com/%v", bucket, tempFileName)
	return url, err
}
