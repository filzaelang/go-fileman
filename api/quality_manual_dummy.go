package api

import (
	"file-manager/helpers"
	"file-manager/models"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterFileRoutes(g *echo.Group) {
	g.GET("/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		filePath, fileName, err := models.FileDownloadHarian(id)
		if err != nil {
			return c.String(http.StatusNotFound, "File not found")
		}

		// Validasi file ada
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return c.String(http.StatusNotFound, "File not found on disk")
		}

		var filePathOut = `C:\FileManager\out.pdf`

		err = helpers.AddPDFWatermark(filePath, filePathOut)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to apply watermark to the file")
		}

		// Buka File
		file, err := os.Open(filePathOut)
		if err != nil {
			return err
		}
		defer file.Close()

		// Set header agar PDF tampil di browser
		c.Response().Header().Set(echo.HeaderContentType, "application/pdf")
		c.Response().Header().Set(echo.HeaderContentDisposition, `inline; filename="`+fileName+`"`)
		return c.Stream(http.StatusOK, "application/pdf", file)
	})
}
