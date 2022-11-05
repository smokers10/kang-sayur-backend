package native

import (
	"encoding/base64"
	"fmt"
	"kang-sayur-backend/infrastructure/filemanager"
	"os"
	"path/filepath"
)

type nativeFM struct{}

// Remove implements filemanager.FilemanagerContract
func (*nativeFM) Remove(path string) (failure error) {
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

// Upload implements filemanager.FilemanagerContract
func (*nativeFM) Upload(fd *filemanager.FileData) (stored_file_path string, failure error) {
	// make directory
	if err := os.MkdirAll(fd.Path, 0755); err != nil {
		return "", err
	}

	// decode base 64 to file
	decodedFile, err := base64.StdEncoding.DecodeString(fd.Base64)
	if err != nil {
		return "", err
	}

	// merge file name /w format
	image := fmt.Sprintf("%s.%s", fd.Filename, fd.Format)

	// join path with file name
	joindedPath := filepath.Join(fd.Path, filepath.Base(image))

	// create named file
	file, err := os.Create(joindedPath)
	if err != nil {
		return "", err
	}

	// close file creation proccess when method is done
	defer file.Close()

	// write
	if _, err := file.Write(decodedFile); err != nil {
		return "", err
	}

	// Sync commits the current contents of the file to stable storage
	if err := file.Sync(); err != nil {
		return "", err
	}

	// return full path
	return joindedPath, nil
}

func NativeFM() filemanager.FilemanagerContract {
	return &nativeFM{}
}
