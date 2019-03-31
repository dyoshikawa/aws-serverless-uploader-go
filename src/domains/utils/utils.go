package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	"io/ioutil"
	"os"
)

func FileGet(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func FileBase64Encode(file []byte) string {
	encoded := base64.StdEncoding.EncodeToString(file)
	return encoded
}

func FileBase64Decode(file []byte) (image.Config, error) {
	r := bytes.NewReader(file)
	decoded, _, err := image.DecodeConfig(r)
	if err != nil {
		return image.Config{}, err
	}
	return decoded, nil
}

func FileBase64Get(path string) (string, error) {
	file, err := FileGet(path)
	if err != nil {
		return "", err
	}
	encoded := FileBase64Encode(file)
	return "data:image/gif;base64," + encoded, nil
}
