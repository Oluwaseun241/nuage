package repositories

import (
	"errors"
	"nuage/internal/entities"
)

type InMemoryFolderRepository struct {
	folders []*entities.Folder
}

type FolderRepositories interface {
	CreateFolder(user *entities.User, name string) (*entities.Folder, error)
	AddFileToFolder(user *entities.User, file *entities.File, fileId int) error
	RemoveFileFromFolder(user *entities.User, fileId int) error
	DeleteFolder(user *entities.User, folderId int) error
}

func (repo *InMemoryFolderRepository) CreateFolder(user *entities.User, name string) (*entities.Folder, error) {
	if len(user.Email) == 0 || name == "" {
		return nil, errors.New("invalid user or empty folder name")
	}
	newFolder := &entities.Folder{
		ID:   1,
		Name: name,
		//Files: [], // Empty folder
		Owner: user,
	}
	repo.folders = append(repo.folders, newFolder)
	return newFolder, nil
}
