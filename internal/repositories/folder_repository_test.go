package repositories

import (
	"nuage/internal/entities"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Test Create Folder
func TestCreateFolder(t *testing.T) {

	user := &entities.User{ID: uuid.New(), Email: "test@example.com"}

	repo := &InMemoryFolderRepository{}

	folder, err := repo.CreateFolder(user, "Docs")

	assert.NoError(t, err)
	assert.NotNil(t, folder)
	assert.Equal(t, folder.Name, "Docs")
	assert.Equal(t, folder.Owner, user)
}

func TestAddFileToFolder(t *testing.T) {}
