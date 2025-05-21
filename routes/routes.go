package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ConfigureRoutes(e *echo.Echo) {
	e.GET("/", func(ctx echo.Context) error {

		props := map[string]interface{}{
			"phrase": "Selamat Datang",
		}

		return ctx.Render(http.StatusOK, "Index", props)
	})

	e.GET("/dashboard", func(ctx echo.Context) error {

		props := map[string]interface{}{
			"phrase": "Dashboard",
		}

		return ctx.Render(http.StatusOK, "Dashboard", props)
	})

	e.GET("/profile", func(ctx echo.Context) error {

		props := map[string]interface{}{
			"phrase": "Profile",
		}

		return ctx.Render(http.StatusOK, "Profile", props)
	})

	e.GET("/setup", func(ctx echo.Context) error {

		props := map[string]interface{}{
			"phrase": "Setup",
		}

		return ctx.Render(http.StatusOK, "Setup", props)
	})

	e.GET("/favorites", func(ctx echo.Context) error {

		props := map[string]interface{}{
			"phrase": "Favorites",
		}

		return ctx.Render(http.StatusOK, "Favorites", props)
	})

	e.GET("/allfiles", func(ctx echo.Context) error {

		props := map[string]interface{}{
			"phrase": "All Files",
		}

		return ctx.Render(http.StatusOK, "AllFiles", props)
	})
}
