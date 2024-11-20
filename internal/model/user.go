package model

type UserWithId struct {
	Id string `json:"id"`
}

/* type UserEntity struct {
	id         string
	userName   string
	isVerified bool
	avatar     string
}

func NewUser(id string, userName string, isVerified bool, avatar string) *UserEntity {
	return &UserEntity{id: id, userName: userName, isVerified: isVerified, avatar: avatar}
}

func (u *UserEntity) CanLogIn() bool {
	return u.isVerified
} */
