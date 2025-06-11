package api

import (
	"file-manager/helpers"
	model_file "file-manager/models/file"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	// "path/filepath"
)

func RegisterFileRoutes(g *echo.Group) {
	g.GET("/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		filePath, fileName, ext, err := model_file.FileDownloadHarian(id)
		if err != nil {
			return c.String(http.StatusNotFound, "File not found")
		}

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return c.String(http.StatusNotFound, "File not found on disk")
		}

		var filePathOut string

		if ext == ".pdf" {
			filePathOut = fmt.Sprintf("C:/FileManager/out/%s", fileName)
		} else {
			filePathOut = filePath
		}

		var username = "admin"

		err = helpers.AddPDFWatermark(filePath, filePathOut, username)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to apply watermark to the file")
		}

		file, err := os.Open(filePathOut)
		if err != nil {
			return err
		}
		defer file.Close()

		// Set header agar PDF tampil di browser
		// Ubah supaya dinamis, jika filenya pdf gunakan kode dibawah ini namun jika non pdf maka download
		c.Response().Header().Set(echo.HeaderContentType, "application/pdf")
		c.Response().Header().Set(echo.HeaderContentDisposition, `inline; filename="`+fileName+`"`)
		return c.Stream(http.StatusOK, "application/pdf", file)
	})
}
