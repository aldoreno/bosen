package login

import (
	"bosen/pkg/response"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

// Adapter-level code to specifically handle HTTP requests.
// It is responsible to transform and validate input before passing it to
// underlying Service that is decoupled from Infrastructure and Adapter code.
type LoginAction struct {
	svc LoginService
}

func NewLoginAction(svc LoginService) *LoginAction {
	return &LoginAction{svc}
}

func (a LoginAction) Handler(req *restful.Request, res *restful.Response) {
	var input LoginInput
	if err := req.ReadEntity(&input); err != nil {
		response.WriteError(res, http.StatusBadRequest, err, restful.MIME_JSON)
		return
	}

	output, err := a.svc.Login(req.Request.Context(), input)
	if err != nil {
		response.WriteError(res, http.StatusUnauthorized, err, restful.MIME_JSON)
		return
	}

	response.WriteSuccess(res, http.StatusOK, &output, restful.MIME_JSON)
}
