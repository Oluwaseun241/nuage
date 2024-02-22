package repositories

import (
	"errors"
	"nuage/internal/entity"
)

type InMemoryFileRepository struct {
	files []*entity.File
}

type FileRepository interface {
	UploadFile(user *entity.User, name string, content []byte) (*entity.File, error)
	DownloadFile(user *entity.User, filename string) ([]byte, error)
}

func (repo *InMemoryFileRepository) UploadFile(user *entity.User, name string, content []byte) (*entity.File, error) {
	// Check user quota and file size
	if len(user.Email) == 0 || len(content) == 0 {
		return nil, errors.New("invalid user or empty content")
	}
	if len(content) > 200*1024*1024 {
		return nil, errors.New("file size exceeds limit")
	}

	// Create new file
	newFile := &entity.File{
		Name:     name,
		Size:     int64(len(content)),
		Owner:    user,
		Contents: content,
	}
	repo.files = append(repo.files, newFile)
	return newFile, nil
}
