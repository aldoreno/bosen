package routes

import (
	AuthRouteHandler "bosen/auth/route/handler"
	"bosen/manifest"
	"bosen/model"
	SessionRouteHandler "bosen/session/route/handler"
	UserRouteHandler "bosen/user/route/handler"
	"net/http"

	"github.com/labstack/echo"
)

type RouteGroup struct{}

var API = RouteGroup{}

func (r RouteGroup) RegisterRoutes(fn func(model.Routes)) {
	routes := model.Routes{
		"inspect": model.Route{
			Method: "GET",
			Path:   "/inspect",
			Handler: func(context echo.Context) error {
				return context.JSON(http.StatusOK, manifest.Info())
			},
		},
		"get-users": model.Route{
			Method:  "GET",
			Path:    "/users",
			Handler: UserRouteHandler.ListUsers,
		},
		"create-user": model.Route{
			Method:  "PUT",
			Path:    "/users/:uid",
			Handler: UserRouteHandler.CreateNewUser,
		},
		"edit-user": model.Route{
			Method:  "POST",
			Path:    "/users/:uid",
			Handler: UserRouteHandler.UpdateUserInfo,
		},
		"detail-user": model.Route{
			Method:  "GET",
			Path:    "/users/:uid",
			Handler: UserRouteHandler.GetUserInfo,
		},
		"delete-user": model.Route{
			Method:  "DELETE",
			Path:    "/users/:uid",
			Handler: UserRouteHandler.DeleteUser,
		},
		"change-user-password": model.Route{
			Method:     "POST",
			Path:       "/users/:uid/password",
			Handler:    UserRouteHandler.ChangeUserPassword,
			Restricted: true,
		},

		// Session
		"get-sessions": model.Route{
			Method:  "GET",
			Path:    "/sessions",
			Handler: SessionRouteHandler.ListSessions,
		},
		"create-session": model.Route{
			Method:  "PUT",
			Path:    "/sessions/:uid",
			Handler: SessionRouteHandler.CreateNewSession,
		},
		"edit-session": model.Route{
			Method:  "POST",
			Path:    "/sessions/:uid",
			Handler: SessionRouteHandler.UpdateSessionInfo,
		},
		"detail-session": model.Route{
			Method:  "GET",
			Path:    "/sessions/:uid",
			Handler: SessionRouteHandler.GetSessionInfo,
		},
		"delete-session": model.Route{
			Method:  "DELETE",
			Path:    "/sessions/:uid",
			Handler: SessionRouteHandler.DeleteSession,
		},

		// Auth
		"authenticate-session": model.Route{
			Method:  "POST",
			Path:    "/auth/authenticate/:sessionId",
			Handler: AuthRouteHandler.AuthenticateSession,
		},
		"inspect-token": model.Route{
			Method:     "GET",
			Path:       "/auth/token",
			Handler:    AuthRouteHandler.InspectToken,
			Restricted: true,
		},
	}

	fn(routes)
}
