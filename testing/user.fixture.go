package testing

type UserFixture struct {
	PhoneNumber string
}

func GenerateUser(phoneNumber string) *UserFixture {
	return &UserFixture{PhoneNumber: phoneNumber}
}
