package utilits

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/skip2/go-qrcode"
)

func GenerateQRCode() (string, error) {
	data, err := GenerateQRCodeData()
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%s.jpg", data)
	filePath := filepath.Join("uploads", filename)
	err = qrcode.WriteFile(data, qrcode.Medium, 256, filePath)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func GenerateQRCodeData() (string, error) {
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", os.ModePerm)
	}

	directory := "uploads"
	files, err := os.ReadDir(directory)
	if err != nil {
		return "", err
	}

	fileCount := len(files)

	month := fmt.Sprintf("%02d", time.Now().Month())
	year := strconv.Itoa(time.Now().Year())

	fileNumber := fmt.Sprintf("%04d", fileCount+1)

	data := fmt.Sprintf("HL%s%s%s", month, year[len(year)-2:], fileNumber)

	return data, nil

}
