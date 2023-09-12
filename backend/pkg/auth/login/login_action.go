package login

import (
	"bosen/internal/trace"
	errs "bosen/pkg/errors"
	"bosen/pkg/response"
	"bosen/pkg/runtime"
	"errors"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
	"github.com/sourcegraph/log"
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
	span, ctx := trace.New(req.Request.Context(), runtime.GetCurrentFunctionName())
	defer span.End()

	l := log.Scoped("loginAction", "authentication handler").WithTrace(trace.Context(ctx))
	l.Info("authentication request received")

	var input LoginInput
	if err := req.ReadEntity(&input); err != nil {
		response.WriteError(res, http.StatusBadRequest, err, restful.MIME_JSON)
		return
	}

	output, err := a.svc.Login(ctx, input)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrAuthCredentials):
			response.WriteError(res, http.StatusUnauthorized, err, restful.MIME_JSON)
		default:
			response.WriteError(res, http.StatusInternalServerError, err, restful.MIME_JSON)
		}
		return
	}

	response.WriteSuccess(res, http.StatusOK, &output, restful.MIME_JSON)
}
