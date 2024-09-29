package repositories

import (
	"errors"
	"nuage/internal/entities"
)

type InMemoryFileRepository struct {
	files []*entities.File
}

type FileRepository interface {
	UploadFile(user *entities.User, name string, content []byte) (*entities.File, error)
	DownloadFile(user *entities.User, filename string) ([]byte, error)
}

func (repo *InMemoryFileRepository) UploadFile(user *entities.User, name string, content []byte) (*entities.File, error) {
	// Check user quota and file size
	if len(user.Email) == 0 || len(content) == 0 {
		return nil, errors.New("invalid user or empty content")
	}
	if len(content) > 200*1024*1024 {
		return nil, errors.New("file size exceeds limit")
	}

	// Create new file
	newFile := &entities.File{
		Name:     name,
		Size:     int64(len(content)),
		Owner:    user,
		Contents: content,
	}
	repo.files = append(repo.files, newFile)
	return newFile, nil
}

func (repo *InMemoryFileRepository) DownloadFile(user *entities.User, filename string) []byte {
	return nil
}
