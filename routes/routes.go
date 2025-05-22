package routes

import (
	"file-manager/models"
	"log"
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

	flatMenus, err := models.GetFlatMenus()
	if err != nil {
		log.Fatal("Failed to load dynamic routes:", err)
	}
	tree := models.BuildMenuTree(flatMenus)
	PropagateFullURI(tree, "")

	// Dynamic routes from DB
	for _, item := range flatMenus {
		if item.Uri != nil && *item.Uri != "" {
			path := *item.Uri
			name := item.Name

			e.GET(path, func(name string) echo.HandlerFunc {
				return func(ctx echo.Context) error {
					return renderWithMenus(ctx, "GeneralPage", name, tree)
				}
			}(name))
		}
	}

	// Static Route
	e.GET("/", func(ctx echo.Context) error {
		return renderWithMenus(ctx, "Index", "Selamat Datang", tree)
	})

	e.GET("/dashboard", func(ctx echo.Context) error {
		return renderWithMenus(ctx, "Dashboard", "Dashboard", tree)
	})

	e.GET("/profile", func(ctx echo.Context) error {
		return renderWithMenus(ctx, "Profile", "Profile", tree)
	})

	e.GET("/setup", func(ctx echo.Context) error {
		return renderWithMenus(ctx, "Setup", "Setup", tree)
	})

	e.GET("/favorites", func(ctx echo.Context) error {
		return renderWithMenus(ctx, "Favorites", "Favorites", tree)
	})

	e.GET("/allfiles", func(ctx echo.Context) error {
		return renderWithMenus(ctx, "AllFiles", "All Files", tree)
	})
}

func renderWithMenus(ctx echo.Context, component string, phrase string, tree []*models.MenuItem) error {
	props := map[string]interface{}{
		"phrase": phrase,
		"menus":  tree,
	}

	return ctx.Render(http.StatusOK, component, props)
}
