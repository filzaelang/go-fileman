package routes

import (
	"file-manager/api"
	"file-manager/middleware"
	"file-manager/models"
	model_file "file-manager/models/file"
	menu "file-manager/models/menu"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ConfigureRoutes(e *echo.Echo) {
	apiGroupMenu := e.Group("/api/menus")
	api.RegisterMenuRoutes(apiGroupMenu)
	apiGroupFile := e.Group("/api/files")
	api.RegisterFileRoutes(apiGroupFile)
	apiGroupDummy := e.Group("/api/dummy")
	api.RegisterFileRoutesDummy(apiGroupDummy)

	// Dynamic routes
	e.GET("/folder/:folderoid/:divoid/:deptoid", dynamicMenuHandler)

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
	}, middleware.RequireAuth)

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

func dynamicMenuHandler(ctx echo.Context) error {
	folderoid, err1 := strconv.Atoi(ctx.Param("folderoid"))
	divoid, err2 := strconv.Atoi(ctx.Param("divoid"))
	deptoid, err3 := strconv.Atoi(ctx.Param("deptoid"))

	if err1 != nil || err2 != nil || err3 != nil {
		return ctx.String(http.StatusBadRequest, "Invalid route parameters")
	}

	menus, err := menu.GetSidebarMenu()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to load menu")
	}

	items, err := model_file.GetFile(folderoid, divoid, deptoid)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to load data")
	}

	props := map[string]interface{}{
		"phrase": "Berikut daftar dokumen yang ditemukan",
		"menus":  menus,
		"role":   "super admin",
		"items":  items,
	}

	return ctx.Render(http.StatusOK, "GeneralPage", props)
}

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
