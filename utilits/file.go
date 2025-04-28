package utilits

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func CreateFile(files []*multipart.FileHeader) ([]string, error) {
	if _, err := os.Stat("uploads/images"); os.IsNotExist(err) {
		os.Mkdir("uploads/images", os.ModePerm)
	}

	var imageURLs []string

	for _, file := range files {
		ext := filepath.Ext(file.Filename)
		newFilename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
		filePath := filepath.Join("uploads/images", newFilename)

		outFile, err := os.Create(filePath)
		if err != nil {
			return nil, errors.New("faýly döretmek şowsuz boldy")
		}
		defer outFile.Close()

		src, err := file.Open()
		if err != nil {
			return nil, errors.New("faýly açmak şowsuz boldy")
		}
		defer src.Close()

		_, err = io.Copy(outFile, src)
		if err != nil {
			return nil, errors.New("faýly ýazmak şowsuz boldy")
		}

		imageURLs = append(imageURLs, "/uploads/images/"+newFilename)
	}

	return imageURLs, nil
}

func RemoveFile(path string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	filePath := filepath.Join(currentDir, path)
	if err = os.Remove(filePath); err != nil {
		return err
	}
	return nil
}
