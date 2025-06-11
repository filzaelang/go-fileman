package api

import (
	"file-manager/helpers"
	"file-manager/models"
	model_file "file-manager/models/file"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
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

		var filePathOut = filePath
		var username = "admin"

		if ext == ".pdf" {
			filePathOut = fmt.Sprintf("C:/FileManager/%s", fileName)
			err = helpers.AddPDFWatermark(filePath, filePathOut, username)
			if err != nil {
				return c.String(http.StatusInternalServerError, "Failed to apply watermark to the file")
			}
		}

		file, err := os.Open(filePathOut)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to open file")
		}
		defer file.Close()

		disposition := "attachment"
		mimeType := helpers.DetectMimeType(ext)
		if ext == ".pdf" {
			disposition = "inline"
			mimeType = "application/pdf"
		}

		c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf(`%s; filename="%s"`, disposition, fileName))
		c.Response().Header().Set(echo.HeaderContentType, mimeType)

		if ext == ".pdf" {
			go helpers.DeleteFile(filePathOut)
		}

		return c.Stream(http.StatusOK, mimeType, file)
	})

	g.POST("", func(c echo.Context) error {
		// Parse multipart form
		fileHeader, err := c.FormFile("file")
		if err != nil {
			return c.String(http.StatusBadRequest, "File tidak ditemukan")
		}

		result, err := models.Upload(fileHeader, c)

		if err != nil {
			return c.String(http.StatusInternalServerError, result)
		}

		return c.String(http.StatusOK, result)
	})
}
