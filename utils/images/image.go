package images

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"strings"
)

type Image struct {
	ContentType string
	Extension   string
	Data        []byte
	Size        int

	image image.Image
}

func (i *Image) DataURI() string {
	return fmt.Sprintf("data:%s;base64,%s", i.ContentType, base64.StdEncoding.EncodeToString(i.Data))
}

func (i *Image) Upload(u Uploader, path string) (*UploadResult, error) {
	return u.Upload(path, i)
}

func (i *Image) Resize(width, height int) (*Image, error) {
	dstImageFit := imaging.Fit(i.image, width, height, imaging.Lanczos)
	b := new(bytes.Buffer)
	switch i.ContentType {
	case "image/png":
		png.Encode(b, dstImageFit)
	case "image/jpeg":
		jpeg.Encode(b, dstImageFit, nil)
	case "image/gif":
		gif.Encode(b, dstImageFit, nil)
	default:
		return nil, errors.New("지원하지 않는 image content 입니다.")
	}

	return &Image{
		ContentType: i.ContentType,
		Extension:   i.Extension,
		Data:        b.Bytes(),
		Size:        0,
		image:       dstImageFit,
	}, nil
}

func Base64EncodingToImage(base64image string) (*Image, error) {
	coI := strings.Index(base64image, ",")

	if coI == -1 {
		return nil, fmt.Errorf("not image file [data=%v]", base64image)
	}

	content := strings.TrimSuffix(base64image[5:coI], ";base64")
	raw := base64image[coI+1:]
	if !isImageContent(content) {
		return nil, errors.New("not image file")
	}

	unbased, _ := base64.StdEncoding.DecodeString(raw)
	res := bytes.NewReader(unbased)

	image := &Image{
		Data:        unbased,
		Size:        len(unbased),
		ContentType: content,
	}
	switch content {
	case "image/png":
		image.image, _ = png.Decode(res)
		image.Extension = "png"
	case "image/jpeg":
		image.image, _ = jpeg.Decode(res)
		image.Extension = "jpeg"
	case "image/gif":
		image.image, _ = gif.Decode(res)
		image.Extension = "gif"
	default:
		return nil, errors.New("지원하지 않는 image content 입니다.")
	}
	return image, nil
}

func isImageContent(contentType string) bool {
	return contentType == "image/png" || contentType == "image/jpeg" || contentType == "image/gif"
}
