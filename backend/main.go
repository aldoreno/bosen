package main

import (
	"bosen/application"
	"bosen/database"
	"bosen/model"
	"bosen/routes"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
)

func main() {
	builder := application.NewBuilder()

	builder.Services.AddDbConfig(func(cfg *application.DbConfig) {
		err := envconfig.Process("BACKEND_DB", cfg)
		if err != nil {
			zap.S().Fatal(err.Error())
		}
		zap.S().Infof("%+v", cfg)
	})

	app := builder.Build()
	app.Start()

	e := echo.New()

	const (
		DB_USER     = "cmsuser"
		DB_PASSWORD = "password"
		DB_NAME     = "cms"
		DB_HOST     = "localhost"
		DB_PORT     = 5432
	)

	if _, err := database.NewDBConnection("postgres", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME); err != nil {
		panic(err)
	}

	db := database.GetDBConnection("postgres")
	defer db.Close()

	// Allow requests from any origin
	// See: https://github.com/labstack/echox/blob/master/cookbook/cors/server.go
	e.Use(middleware.CORS())

	e.Use(func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			appContext := application.NewApplicationContext(context, db)
			return handler(appContext)
		}
	})

	secret := []byte("seqCEeCMUK8cd4RjMaARwqf8XnuZqkL567FysqRvZJyMTkM3H9uztKtykqtSJkqVNFhnERva")
	JWTMiddleware := middleware.JWT(secret)

	routes.API.RegisterRoutes(func(routes model.Routes) {
		for routeName, route := range routes {
			if route.Restricted {
				// Restricted route reference implementation (minus echo.Group)
				// See: https://github.com/labstack/echox/blob/master/cookbook/jwt/map-claims/server.go#L43
				switch route.Method {
				case "GET":
					e.GET(route.Path, route.Handler, JWTMiddleware).Name = routeName
				case "POST":
					e.POST(route.Path, route.Handler, JWTMiddleware).Name = routeName
				case "PUT":
					e.PUT(route.Path, route.Handler, JWTMiddleware).Name = routeName
				case "DELETE":
					e.DELETE(route.Path, route.Handler, JWTMiddleware).Name = routeName
				}

				continue
			}

			switch route.Method {
			case "GET":
				e.GET(route.Path, route.Handler).Name = routeName
			case "POST":
				e.POST(route.Path, route.Handler).Name = routeName
			case "PUT":
				e.PUT(route.Path, route.Handler).Name = routeName
			case "DELETE":
				e.DELETE(route.Path, route.Handler).Name = routeName
			}
		}
	})

	e.Logger.Fatal(e.Start(":8080"))
}
