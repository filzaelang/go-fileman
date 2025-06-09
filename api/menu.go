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

	g.POST("/getone", func(c echo.Context) error {
		var payload models.UpdateMenuPayload
		if err := c.Bind(&payload); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		name, err := models.GetOneMenu(payload)
		if err != nil {
			return c.String(http.StatusNotFound, "Menu not found")
		}

		return c.JSON(http.StatusOK, name)
	})

	g.POST("", func(c echo.Context) error {
		var payload models.AddMenuPayload
		if err := c.Bind(&payload); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}
		if err := models.InsertMenu(payload); err != nil {
			return c.String(http.StatusInternalServerError, "Insert failed")
		}
		return c.Redirect(http.StatusSeeOther, "/")
		// return c.Redirect(http.StatusSeeOther, c.Request().RequestURI)
	})

	g.PUT("", func(c echo.Context) error {
		var payload models.UpdateMenuPayload
		if err := c.Bind(&payload); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}
		if err := models.UpdateMenu(payload); err != nil {
			return c.String(http.StatusInternalServerError, "Update failed")
		}
		return c.Redirect(http.StatusSeeOther, "/")
		// return c.Redirect(http.StatusSeeOther, c.Request().RequestURI)
	})

	g.POST("/delete", func(c echo.Context) error {
		var payload models.DeleteMenuPayload

		if err := c.Bind(&payload); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		if err := models.DeleteMenu(payload); err != nil {
			return c.String(http.StatusInternalServerError, "Delete failed")
		}
		return c.Redirect(http.StatusSeeOther, "/")
	})

	g.POST("/bulist", func(c echo.Context) error {
		var payload models.FolderID
		if err := c.Bind(&payload); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		bu_list, err := models.GetBUList(payload)
		if err != nil {
			return c.String(http.StatusNotFound, "BU's not found")
		}
		return c.JSON(http.StatusOK, bu_list)
	})
}
