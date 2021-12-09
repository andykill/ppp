package entity

type TypeUser struct {
	Id   uint32
	Name string
}

type User struct {
	Id   uint32
	Name string
	Type *TypeUser
}

func NewUser(name string, t *TypeUser) *User {
	return &User{
		Name: name,
		Type: t,
	}
}
