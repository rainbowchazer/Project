package file

import (
	"bufio"
	"fmt"
	"idel/enums"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type FileService struct {
	BaseUploadPath string
}

func NewFileService(baseUploadPath string) *FileService {
	return &FileService{
		BaseUploadPath: baseUploadPath,
	}
}

func (fs *FileService) GetFileContentByPath(path string) (string, error) {
	fullPath := filepath.Join(fs.BaseUploadPath, path)
	file, err := os.Open(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var contentBuilder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contentBuilder.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return contentBuilder.String(), nil
}

func (fs *FileService) UploadContent(content, filename string, prefix enums.FilePathPrefix) (string, error) {
	uploadDir := filepath.Join(fs.BaseUploadPath, prefix.String())
	if err := ensureDirectoryExists(uploadDir); err != nil {
		return "", err
	}

	filePath := filepath.Join(uploadDir, filename)
	if err := ioutil.WriteFile(filePath, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	return filepath.Join(prefix.String(), filename), nil
}

func (fs *FileService) UploadFile(fileContent []byte, originalName string, prefix enums.FilePathPrefix) (string, error) {
	uploadDir := filepath.Join(fs.BaseUploadPath, prefix.String())
	if err := ensureDirectoryExists(uploadDir); err != nil {
		return "", err
	}

	fileUUID := uuid.New().String()
	ext := filepath.Ext(originalName)
	filename := fileUUID + ext
	filePath := filepath.Join(uploadDir, filename)

	if err := ioutil.WriteFile(filePath, fileContent, 0644); err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	return filepath.Join(prefix.String(), filename), nil
}

func (fs *FileService) Delete(path string) error {
	fullPath := filepath.Join(fs.BaseUploadPath, path)
	if err := os.Remove(fullPath); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}

func (fs *FileService) SafeDelete(path string) {
	if err := fs.Delete(path); err != nil {
		log.Printf("Safe delete error: %v\n", err)
	}
}

func (fs *FileService) SafeDeleteMultiple(paths []string) {
	for _, path := range paths {
		fs.SafeDelete(path)
	}
}

func ensureDirectoryExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	return nil
}
