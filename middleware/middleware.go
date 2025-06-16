package middleware

import (
	"file-manager/inertiaMiddleware"
	"log"
	"net/url"
	"os"

	"github.com/gorilla/sessions"
	session "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigureMiddleware(e *echo.Echo) {

	assetVersion := "1"
	e.Use(inertiaMiddleware.InertiaMiddleware(e, assetVersion))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret-key"))))
	e.Static("/", "views/public")

	isDevelopment := os.Getenv("BUILD_ENV") == "development"

	if isDevelopment {
		viteDevServer, err := url.Parse("http://localhost:5173")
		if err != nil {
			log.Fatal("Could not parse Vite dev server url", err)
		}

		devAssets := e.Group("/src/assets")
		target := []*middleware.ProxyTarget{
			{Name: "viteProxyTarget",
				URL: viteDevServer}}

		loadbalancer := middleware.NewRoundRobinBalancer(target)

		devAssets.Use(middleware.Proxy(loadbalancer))
		return
	}
}
