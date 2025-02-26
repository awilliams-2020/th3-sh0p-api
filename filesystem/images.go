package filesystem

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofrs/uuid"
)

func saveImageFromURL(url, savePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http request failed: %v", resp.Status)
	}

	if err := os.MkdirAll(filepath.Dir(savePath), 0755); err != nil {
		return err
	}

	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func SaveImage(url string) (string, error) {
	u2, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	fileName := fmt.Sprintf("%s.png", u2.String())
	savePath := filepath.Join("th3-sh0p-images", fileName)
	if err := saveImageFromURL(url, savePath); err != nil {
		return "", err
	}
	return fileName, nil
}
