package routes

import (
	"file-manager/api"
	"file-manager/models"
	menu "file-manager/models/menu"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ConfigureRoutes(e *echo.Echo) {
	apiGroupMenu := e.Group("/api/menus")
	api.RegisterMenuRoutes(apiGroupMenu)
	apiGroupFolder := e.Group("/api/folder")
	api.RegisterFileRoutes(apiGroupFolder)

	// // Dynamic routes
	// e.GET("/folder/:folderoid/:divoid/:deptoid", dynamicMenuHandler)

	// Static Route
	e.GET("/", func(ctx echo.Context) error {
		items, _ := models.GetFile()
		return renderWithMenus(ctx, "Index", "Selamat Datang", items)
	})

	e.GET("/dashboard", func(ctx echo.Context) error {
		items, _ := models.GetFile()
		return renderWithMenus(ctx, "Dashboard", "Dashboard", items)
	})

	e.GET("/profile", func(ctx echo.Context) error {
		items, _ := models.GetFile()
		return renderWithMenus(ctx, "Profile", "Profile", items)
	})

	e.GET("/setup", func(ctx echo.Context) error {
		items, _ := models.GetFile()
		return renderWithMenus(ctx, "Setup", "Setup", items)
	})

	e.GET("/favorites", func(ctx echo.Context) error {
		items, _ := models.GetFile()
		return renderWithMenus(ctx, "Favorites", "Favorites", items)
	})

	e.GET("/allfiles", func(ctx echo.Context) error {
		items, _ := models.GetFile()
		return renderWithMenus(ctx, "AllFiles", "All Files", items)
	})

	e.GET("/setup-menu", func(ctx echo.Context) error {
		items, _ := models.GetFile()
		return renderWithMenus(ctx, "SetupMenu", "SetupMenu", items)
	})
}

// func dynamicMenuHandler(ctx echo.Context) error {
// 	requestedUri := ctx.Request().URL.Path

// 	menus, err := menu.GetSidebarMenu()
// 	if err != nil {
// 		log.Fatal("Failed to load menus", err)
// 	}

// 	props := map[string]interface{}{
// 		"phrase": ,
// 		"menus":  menus,
// 		"role":   "super admin",
// 	}

// 	return ctx.Render(http.StatusOK, "GeneralPage", props)
// }

func renderWithMenus(ctx echo.Context, component string, phrase string, items []models.FileItem) error {
	menus, err := menu.GetSidebarMenu()
	if err != nil {
		log.Fatal("Failed to load menus", err)
	}

	props := map[string]interface{}{
		"phrase": phrase,
		"menus":  menus,
		"role":   "super admin",
		"items":  items,
	}

	return ctx.Render(http.StatusOK, component, props)
}
