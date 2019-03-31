package images

import (
	"bytes"
	"image/jpeg"
	"image/png"

	"github.com/nfnt/resize"
)

const size = 300

func resizePng(file []byte) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(file))
	if err != nil {
		return nil, err
	}
	resizedImg := resize.Resize(size, 0, img, resize.Lanczos3)
	buf := new(bytes.Buffer)
	err = png.Encode(buf, resizedImg)
	if err != nil {
		return nil, err
	}
	resizedBytes := buf.Bytes()
	return resizedBytes, nil
}

func resizeJpeg(file []byte) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewReader(file))
	if err != nil {
		return nil, err
	}
	resizedImg := resize.Resize(size, 0, img, resize.Lanczos3)
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, resizedImg, nil)
	if err != nil {
		return nil, err
	}
	resizedBytes := buf.Bytes()
	return resizedBytes, nil
}
