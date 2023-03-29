package auth

import (
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

// Adapter-level code to specifically handle HTTP requests.
// It is responsible to transform and validate input before passing it to
// underlying Service that is decoupled from Infrastructure and Adapter code.
type LoginAction struct {
	svc *AuthService
}

func NewLoginAction(svc *AuthService) *LoginAction {
	return &LoginAction{svc}
}

func (a LoginAction) Handler(req *restful.Request, res *restful.Response) {
	var (
		err         error
		credentials LoginInput
	)

	if err = req.ReadEntity(&credentials); err != nil {
		res.WriteHeaderAndJson(http.StatusBadRequest, err, restful.MIME_JSON)
		return
	}

	res.WriteAsJson(&credentials)
}
