package auth

import (
	"bosen/pkg/response"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

// Adapter-level code to specifically handle HTTP requests.
// It is responsible to transform and validate input before passing it to
// underlying Service that is decoupled from Infrastructure and Adapter code.
type LoginAction struct {
	svc *authService
}

func NewLoginAction(svc *authService) *LoginAction {
	return &LoginAction{svc}
}

func (a LoginAction) Handler(req *restful.Request, res *restful.Response) {
	var credentials LoginInput
	if err := req.ReadEntity(&credentials); err != nil {
		response.WriteError(res, http.StatusBadRequest, err, restful.MIME_JSON)
		return
	}

	var token AuthToken
	if err := a.svc.Login(credentials, &token); err != nil {
		response.WriteError(res, http.StatusUnauthorized, err, restful.MIME_JSON)
		return
	}

	response.WriteSuccess(res, http.StatusUnauthorized, &token, restful.MIME_JSON)
}
