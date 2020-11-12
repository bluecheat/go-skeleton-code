package images

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type Uploader interface {
	Upload(path string, image *Image) (*UploadResult, error)
}

type UploadResult struct {
	Name       string
	FullPath   string
	OriginPath string
}

func NewUploader(config *config.Config) Uploader {
	return &FileUploader{
		config.FilePath,
	}
}

type FileUploader struct {
	DefaultPath string
}

func (fu *FileUploader) Upload(path string, image *Image) (*UploadResult, error) {
	fileName := utils.RandToken(20) + "." + image.Extension
	originPath := filepath.Join(fu.DefaultPath, path)
	createDirIfNotExist(originPath)
	fullPath := originPath + "/" + fileName
	err := ioutil.WriteFile(fullPath, image.Data, 0600)
	if err != nil {
		return nil, err
	}
	return &UploadResult{
		Name:       fileName,
		OriginPath: originPath,
		FullPath:   fullPath,
	}, nil
}

type AWSUploader struct {
	DefaultPath string
}

func (fu *AWSUploader) Upload(path string, image *Image) (*UploadResult, error) {
	fileName := utils.RandToken(20) + "." + image.Extension
	originPath := filepath.Join(fu.DefaultPath, path)
	createDirIfNotExist(originPath)
	fullPath := originPath + "/" + fileName
	err := ioutil.WriteFile(fullPath, image.Data, 0600)
	if err != nil {
		return nil, err
	}
	return &UploadResult{
		Name:       fileName,
		OriginPath: originPath,
		FullPath:   fullPath,
	}, nil
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
