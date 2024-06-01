package entities

type File struct {
	Name     string
	Size     int64
	Owner    *User
	Contents []byte
}
