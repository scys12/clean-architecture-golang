package aws

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

const (
	maxSize = int64(5120000)
	bucket  = ""
)

type S3Store interface {
	uploadFileToS3(*FileForm, string, string) (string, error)
}

func uploadFileToS3(fileF *FileForm, ID uint64, folder string) (string, error) {
	size := fileF.fileHeader.Size
	buffer := make([]byte, size)
	fileF.file.Read(buffer)
	tempFileName := fmt.Sprintf("%v/%v-%v%v", folder, strconv.FormatUint(ID, 10), uuid.New().String(), filepath.Ext(fileF.fileHeader.Filename))
	_, err := s3.New(fileF.session).PutObject(&s3.PutObjectInput{
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
