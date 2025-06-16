package api

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func RegisterAuthRoutes(g *echo.Group) {
	g.POST("/login", loginHandler)
	g.POST("/logout", logoutHandler)
}

func loginHandler(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid input")
	}

	// Dummy check (use real logic later)
	if req.Username == "user" && req.Password == "user" {
		sess, _ := session.Get("session", c)
		sess.Values["authenticated"] = true
		sess.Values["username"] = req.Username
		sess.Save(c.Request(), c.Response())
		return c.JSON(http.StatusOK, map[string]string{"message": "Login success"})
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
}

func logoutHandler(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())
	return c.JSON(http.StatusOK, map[string]string{"message": "Logged out"})
}
