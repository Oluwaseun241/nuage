package entity

type Folder struct {
	Name  string
	Files []*File
	Owner *User
}
