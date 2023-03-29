package auth

import (
	"bosen/pkg/auth/login"
	"net/http"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	restful "github.com/emicklei/go-restful/v3"
)

// Infrastructure-level code tightly coupled with go-restul/v3
type AuthResource struct {
	loginAction *login.LoginAction
}

func NewAuthResource(loginAction *login.LoginAction) *AuthResource {
	return &AuthResource{
		loginAction: loginAction,
	}
}

func (r *AuthResource) WebService() *restful.WebService {
	ws := new(restful.WebService).
		Path("/auth").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/login").To(r.loginAction.Handler).
		Doc("user authentication").
		Metadata(restfulspec.KeyOpenAPITags, []string{"auth", "user", "login"}).
		Writes(&login.LoginOutput{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), &login.LoginOutput{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil),
	)

	return ws
}
