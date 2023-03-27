package auth

import (
	"net/http"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	restful "github.com/emicklei/go-restful/v3"
)

// Infrastructure-level code tightly coupled with go-restul/v3
type AuthResource struct {
	authenticate *AuthenticateSessionAction
}

func NewAuthResource(authenticate *AuthenticateSessionAction) *AuthResource {
	return &AuthResource{
		authenticate: authenticate,
	}
}

func (r *AuthResource) WebService() *restful.WebService {
	ws := new(restful.WebService).
		Path("/auth").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/authenticate/{sessionId}").To(r.authenticate.AuthenticateSession).
		Doc("user authentication").
		Metadata(restfulspec.KeyOpenAPITags, []string{"auth", "user", "session"}).
		Writes(AuthToken{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), AuthToken{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil),
	)

	return ws
}
