package repositories

import (
	"nuage/internal/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadFile(t *testing.T) {
	repo := &InMemoryFileRepository{}
	user := &entities.User{Email: "test@example.com"}
	content := []byte("This is a test file content")
	name := "testfile.txt"

	file, err := repo.UploadFile(user, name, content)

	assert.NoError(t, err)
	assert.NotNil(t, file)
	assert.Equal(t, name, file.Name)
	assert.Equal(t, int64(len(content)), file.Size)
	assert.Equal(t, user, file.Owner)
	assert.Equal(t, content, file.Contents)
}

func TestDownloadFile()
