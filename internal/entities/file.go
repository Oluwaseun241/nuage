package entities

type File struct {
	ID       int
	Name     string
	Size     int64
	Owner    *User
	Contents []byte
}
