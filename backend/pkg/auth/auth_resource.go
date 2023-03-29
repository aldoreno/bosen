package auth

import (
	"net/http"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	restful "github.com/emicklei/go-restful/v3"
)

// Infrastructure-level code tightly coupled with go-restul/v3
type AuthResource struct {
	loginAction *LoginAction
}

func NewAuthResource(loginAction *LoginAction) *AuthResource {
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
		Writes(AuthToken{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), &AuthToken{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil),
	)

	return ws
}
