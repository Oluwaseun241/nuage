package entities

type Folder struct {
	Name  string
	Files []*File
	Owner *User
}
