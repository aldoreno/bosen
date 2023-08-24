package login

import (
	"bosen/manifest"
	errs "bosen/pkg/errors"
	"bosen/pkg/response"
	"bosen/pkg/runtime"
	"errors"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
	"go.opentelemetry.io/otel"
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
	ctx, span := otel.Tracer(manifest.AppName).Start(req.Request.Context(), runtime.GetCurrentFunctionName())
	defer span.End()

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
