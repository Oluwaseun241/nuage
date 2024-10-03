package repositories

import (
	"nuage/internal/entities"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Test Create Folder
func TestCreateFolder(t *testing.T) {

	user := &entities.User{ID: uuid.New(), Email: "test@email.com"}

	repo := &InMemoryFolderRepository{}

	folder, err := repo.CreateFolder(user, "Docs")

	assert.NoError(t, err)
	assert.NotNil(t, folder)
	assert.Equal(t, folder.Name, "Docs")
	assert.Equal(t, folder.Owner, user)
}

func TestAddFileToFolder(t *testing.T) {
	user := &entities.User{ID: uuid.New(), Email: "test@email.com"}

	repo := &InMemoryFolderRepository{}

	file := &entities.File{ID: 1, Name: "file.txt"}

	folder, err := repo.CreateFolder(user, "Docs")

	addToFolder, err := repo.AddFileToFolder(user, file, folder.ID)

	assert.NoError(t, err)
	assert.Equal(t, addToFolder.Files[0].Name, file.Name)
	assert.Equal(t, len(addToFolder.Files), 1)
}

func TestRemoveFileFromFolder(t *testing.T) {
	user := &entities.User{ID: uuid.New(), Email: "test@email.com"}

	repo := &InMemoryFolderRepository{}

	file1 := &entities.File{ID: 1, Name: "file1.txt"}
	file2 := &entities.File{ID: 2, Name: "file2.txt"}

	folder, err := repo.CreateFolder(user, "Docs")

	_, err = repo.AddFileToFolder(user, file1, folder.ID)
	_, err = repo.AddFileToFolder(user, file2, folder.ID)

	removeFromFolder, err := repo.RemoveFileFromFolder(user, folder.ID, file1.ID)

	assert.NoError(t, err)
	assert.Equal(t, len(removeFromFolder.Files), 1)
	assert.Equal(t, removeFromFolder.Files[0].ID, file2.ID)
}

func TestDeleteFolder(t *testing.T) {
	user := &entities.User{ID: uuid.New(), Email: "test@email.com"}

	repo := &InMemoryFolderRepository{}

	folder, err := repo.CreateFolder(user, "Docs")
	assert.NoError(t, err)

	err = repo.DeletFolder(user, folder.ID)
	assert.NoError(t, err)

	assert.Equal(t, len(repo.folders), 0)
}
