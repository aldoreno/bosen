package auth

import (
	"fmt"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
	uuid "github.com/satori/go.uuid"
)

// Adapter-level code to specifically handle HTTP requests.
// It is responsible to transform and validate input before passing it to
// underlying Service that is decoupled from Infrastructure and Adapter code.
type AuthenticateSessionAction struct {
	svc *AuthService
}

func NewAuthSessAction(svc *AuthService) *AuthenticateSessionAction {
	return &AuthenticateSessionAction{svc}
}

func (a AuthenticateSessionAction) AuthenticateSession(req *restful.Request, res *restful.Response) {
	var (
		err         error
		credentials AuthenticateSessionInput
	)

	if err = req.ReadEntity(&credentials); err != nil {
		res.WriteError(http.StatusBadRequest, err)
		return
	}

	credentials.SessionId, err = uuid.FromString(req.PathParameter("sessionId"))
	if err != nil {
		res.WriteError(http.StatusBadRequest, fmt.Errorf("%w", err))
		return
	}
}
