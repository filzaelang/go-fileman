package api

import (
	"file-manager/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterMenuRoutes(g *echo.Group) {
	g.GET("", func(c echo.Context) error {
		flat, _ := models.GetFlatMenus()
		tree := models.BuildMenuTree(flat)
		return c.JSON(http.StatusOK, tree)
	})

	g.POST("", func(c echo.Context) error {
		var payload models.MenuItem
		if err := c.Bind(&payload); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}
		if err := models.InsertMenu(payload); err != nil {
			return c.String(http.StatusInternalServerError, "Insert failed")
		}
		return c.NoContent(http.StatusCreated)
	})

	g.PUT("/:id", func(c echo.Context) error {
		var payload models.MenuItem
		if err := c.Bind(&payload); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}
		payload.ID, _ = strconv.Atoi(c.Param("id"))
		if err := models.UpdateMenu(payload); err != nil {
			return c.String(http.StatusInternalServerError, "Update failed")
		}
		return c.NoContent(http.StatusOK)
	})

	g.DELETE("/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := models.DeleteMenu(id); err != nil {
			return c.String(http.StatusInternalServerError, "Delete failed")
		}
		return c.NoContent(http.StatusOK)
	})
}
