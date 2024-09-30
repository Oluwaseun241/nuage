package entities

type Folder struct {
	ID    int
	Name  string
	Files []*File
	Owner *User
}
