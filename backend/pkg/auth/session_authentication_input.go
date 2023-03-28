package auth

import uuid "github.com/satori/go.uuid"

type (
	Username                 string
	Password                 string
	AuthenticateSessionInput struct {
		SessionId uuid.UUID
		Username  Username `json:"username" form:"username"`
		Password  Password `json:"password" form:"password"`
	}
)
