package api

import (
	"file-manager/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterMenuRoutes(g *echo.Group) {
	g.GET("", func(c echo.Context) error {
		menu_list, err := models.GetSidebarMenu()
		if err != nil {
			return c.String(http.StatusNotFound, "Menu not found")
		}
		return c.JSON(http.StatusOK, menu_list)
	})

	// g.GET("", func(c echo.Context) error {
	// 	flat, _ := models.GetFlatMenus()
	// 	tree := models.BuildMenuTree(flat)
	// 	return c.JSON(http.StatusOK, tree)
	// })

	// g.GET("/:id", func(c echo.Context) error {
	// 	id, _ := strconv.Atoi(c.Param("id"))

	// 	menu, err := models.GetOneMenu(id)
	// 	if err != nil {
	// 		return c.String(http.StatusNotFound, "Menu not found")
	// 	}

	// 	return c.JSON(http.StatusOK, menu)
	// })

	// g.POST("", func(c echo.Context) error {
	// 	var payload models.MenuItem
	// 	if err := c.Bind(&payload); err != nil {
	// 		return c.String(http.StatusBadRequest, "Invalid input")
	// 	}
	// 	if err := models.InsertMenu(payload); err != nil {
	// 		return c.String(http.StatusInternalServerError, "Insert failed")
	// 	}
	// 	return c.Redirect(http.StatusSeeOther, "/")
	// 	// return c.Redirect(http.StatusSeeOther, c.Request().RequestURI)
	// })

	// g.PUT("/:id", func(c echo.Context) error {
	// 	var payload models.MenuItem
	// 	if err := c.Bind(&payload); err != nil {
	// 		return c.String(http.StatusBadRequest, "Invalid input")
	// 	}
	// 	payload.ID, _ = strconv.Atoi(c.Param("id"))
	// 	if err := models.UpdateMenu(payload); err != nil {
	// 		return c.String(http.StatusInternalServerError, "Update failed")
	// 	}
	// 	return c.Redirect(http.StatusSeeOther, "/")
	// 	// return c.Redirect(http.StatusSeeOther, c.Request().RequestURI)
	// })

	// g.DELETE("/:id", func(c echo.Context) error {
	// 	id, _ := strconv.Atoi(c.Param("id"))
	// 	if err := models.DeleteMenu(id); err != nil {
	// 		return c.String(http.StatusInternalServerError, "Delete failed")
	// 	}
	// 	return c.Redirect(http.StatusSeeOther, "/")
	// })
}
