package user

type (
	Username   string
	Password   string
	UserEntity struct {
		Username   Username
		Password   Password
		FirstName  string
		MiddleName string
		LastName   string
	}
)
