package api

import (
	"net/http"

	folder "file-manager/models/folder"

	"github.com/labstack/echo/v4"
)

func RegisterFolderRoutes(g *echo.Group) {
	g.GET("/:folderoid/:divoid/:deptoiddept", func(c echo.Context) error {
		var folderoid int
		var divoid int
		var deptoid int
		file_list, err := folder.Folder(folderoid, divoid, deptoid)
		if err != nil {
			return c.String(http.StatusNotFound, "Menu not found")
		}
		return c.JSON(http.StatusOK, file_list)
	})
}
