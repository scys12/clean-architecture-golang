package delivery

import (
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/scys12/clean-architecture-golang/pkg/payload/request"
)

func BindingFormValue(e echo.Context) (map[string]interface{}, error) {
	form := make(map[string]interface{})
	values, err := e.FormParams()
	if err != nil {
		return nil, err
	}
	for key, vals := range values {
		for _, val := range vals {
			form[key] = val
		}
	}
	return form, nil
}

func BindingFormFile(e echo.Context, formName string) (*multipart.FileHeader, error) {
	file, err := e.FormFile(formName)
	if err != nil {
		return nil, nil
	}
	err = checkFileMimeType(file)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func checkFileMimeType(fileHeader *multipart.FileHeader) error {
	buffer := make([]byte, 512)
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Read(buffer)
	if err != nil {
		return nil
	}
	contentType := http.DetectContentType(buffer)
	if contentType != "image/jpeg" && contentType != "image/png" {
		return errors.New("Not an image file")
	}
	return nil
}

func DecodeForm(req *request.ProfileRequest, form map[string]interface{}, file *multipart.FileHeader) error {
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &req,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	err = decoder.Decode(form)
	if err != nil {
		return err
	}
	req.Image = file
	return nil
}
