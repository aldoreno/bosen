package auth

type (
	Username   string
	Password   string
	LoginInput struct {
		Username Username `json:"username" form:"username"`
		Password Password `json:"password" form:"password"`
	}
)
