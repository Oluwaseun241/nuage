package repositories

import (
	"errors"
	"nuage/internal/entities"
)

var (
	ErrFolderNotFound = errors.New("folder not found")
)

type InMemoryFolderRepository struct {
	folders []*entities.Folder
}

type FolderRepositories interface {
	CreateFolder(user *entities.User, name string) (*entities.Folder, error)
	AddFileToFolder(user *entities.User, file *entities.File, folderId int) (*entities.Folder, error)
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

func (repo *InMemoryFolderRepository) AddFileToFolder(user *entities.User, file *entities.Folder, folderId int) (*entities.Folder, error) {
	for _, folder := range repo.folders {
		if folderId == folder.ID {
			folder.Files = file.Files
			return folder, nil
		}
	}
	return nil, ErrFolderNotFound
}
