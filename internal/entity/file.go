package entity

type File struct {
	Name     string
	Size     string
	Owner    *User
	Contents []byte
}
