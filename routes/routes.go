package routes

import (
	"file-manager/api"
	"file-manager/models"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// Gabungkan dua path jadi /a/b (tanpa double slash)
func joinURI(base, uri string) string {
	return strings.TrimRight(base, "/") + "/" + strings.TrimLeft(uri, "/")
}

func PropagateFullURI(nodes []*models.MenuItem, base string) {
	for _, node := range nodes {
		// Gunakan base + child path jika ada
		if node.Uri != nil {
			newUri := joinURI(base, *node.Uri)
			node.Uri = &newUri
		}

		// Rekursi ke children
		if len(node.Children) > 0 {
			PropagateFullURI(node.Children, *node.Uri)
		}
	}
}

func ConfigureRoutes(e *echo.Echo) {
	// apiGroup := e.Group("/api/menus")
	// api.RegisterMenuRoutes(apiGroup)
	apiGroupFiles := e.Group("/api/files")
	api.RegisterFileRoutes(apiGroupFiles)

	// // Dynamic routes from DB
	// e.GET("/*", dynamicMenuHandler)

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

// 	flatMenus, err := models.GetFlatMenus()
// 	if err != nil {
// 		return ctx.String(http.StatusInternalServerError, "Gagal ambil menu")
// 	}
// 	tree := models.BuildMenuTree(flatMenus)
// 	PropagateFullURI(tree, "")

// 	var matched *models.MenuItem
// 	for _, m := range flatMenus {
// 		if m.Uri != nil && *m.Uri == requestedUri {
// 			matched = &m
// 			break
// 		}
// 	}

// 	if matched == nil {
// 		return ctx.String(http.StatusNotFound, "Halaman tidak ditemukan")
// 	}

// 	props := map[string]interface{}{
// 		"phrase": matched.Name,
// 		"menus":  tree,
// 		"role":   "super admin",
// 	}

// 	return ctx.Render(http.StatusOK, "GeneralPage", props)
// }

func renderWithMenus(ctx echo.Context, component string, phrase string, items []models.FileItem) error {
	// flatMenus, err := models.GetFlatMenus()
	// if err != nil {
	// 	log.Fatal("Failed to load dynamic routes:", err)
	// }
	// tree := models.BuildMenuTree(flatMenus)
	// PropagateFullURI(tree, "")

	props := map[string]interface{}{
		"phrase": phrase,
		"menus":  nil,
		"role":   "super admin",
		"items":  items,
	}

	return ctx.Render(http.StatusOK, component, props)
}
