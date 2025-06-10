package api

import (
	menu "file-manager/models/menu"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterMenuRoutes(g *echo.Group) {
	g.GET("", func(c echo.Context) error {
		menu_list, err := menu.GetSidebarMenu()
		if err != nil {
			return c.String(http.StatusNotFound, "Menu not found")
		}
		return c.JSON(http.StatusOK, menu_list)
	})

	g.POST("/getone", func(c echo.Context) error {
		var payload menu.UpdateMenuPayload
		if err := c.Bind(&payload); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		name, err := menu.GetOneMenu(payload)
		if err != nil {
			return c.String(http.StatusNotFound, "Menu not found")
		}

		return c.JSON(http.StatusOK, name)
	})

	g.POST("", func(c echo.Context) error {
		var payload menu.AddMenuPayload
		if err := c.Bind(&payload); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}
		if err := menu.InsertMenu(payload); err != nil {
			return c.String(http.StatusInternalServerError, "Insert failed")
		}
		return c.Redirect(http.StatusSeeOther, "/")
		// return c.Redirect(http.StatusSeeOther, c.Request().RequestURI)
	})

	g.PUT("", func(c echo.Context) error {
		var payload menu.UpdateMenuPayload
		if err := c.Bind(&payload); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}
		if err := menu.UpdateMenu(payload); err != nil {
			return c.String(http.StatusInternalServerError, "Update failed")
		}
		return c.Redirect(http.StatusSeeOther, "/")
		// return c.Redirect(http.StatusSeeOther, c.Request().RequestURI)
	})

	g.POST("/delete", func(c echo.Context) error {
		var payload menu.DeleteMenuPayload

		if err := c.Bind(&payload); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		if err := menu.DeleteMenu(payload); err != nil {
			return c.String(http.StatusInternalServerError, "Delete failed")
		}
		return c.Redirect(http.StatusSeeOther, "/")
	})

	g.POST("/bulist", func(c echo.Context) error {
		var payload menu.BuChildList
		if err := c.Bind(&payload); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		bu_list, err := menu.GetBUList(payload)
		if err != nil {
			return c.String(http.StatusNotFound, "BU's not found")
		}
		return c.JSON(http.StatusOK, bu_list)
	})
}
