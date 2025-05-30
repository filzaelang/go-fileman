package api

import (
	"file-manager/helpers"
	"file-manager/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func RegisterFileRoutes(g *echo.Group) {
	g.GET("/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		filePath, fileName, err := models.FileDownloadHarian(id)
		if err != nil {
			return c.String(http.StatusNotFound, "File not found")
		}

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return c.String(http.StatusNotFound, "File not found on disk")
		}

		var filePathOut = `C:\FileManager\out.pdf`
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
		c.Response().Header().Set(echo.HeaderContentType, "application/pdf")
		c.Response().Header().Set(echo.HeaderContentDisposition, `inline; filename="`+fileName+`"`)
		return c.Stream(http.StatusOK, "application/pdf", file)
	})

	g.POST("", func(c echo.Context) error {
		// Parse multipart form
		fileHeader, err := c.FormFile("file")
		if err != nil {
			return c.String(http.StatusBadRequest, "File tidak ditemukan")
		}

		src, err := fileHeader.Open()
		if err != nil {
			return c.String(http.StatusInternalServerError, "Gagal membuka file")
		}
		defer src.Close()

		// get metadata from form
		ducumentNumber := c.FormValue("document_number")
		documentName := c.FormValue("document_name")
		revisionNumber := c.FormValue("revision_number")
		// revisionDate := c.FormValue("revision_date")

		// Save file to C:\FileManager\ with unique name
		timestamp := time.Now().Format("20060102_150405")
		safeFileName := fmt.Sprintf("%s_%s_%s_%s%s",
			ducumentNumber,
			documentName,
			revisionNumber,
			timestamp,
			filepath.Ext(fileHeader.Filename),
		)
		targetPath := filepath.Join("C:\\FileManager", safeFileName)

		dst, err := os.Create(targetPath)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Gagal menyimpan file")
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return c.String(http.StatusInternalServerError, "Gagal menyalin file")
		}

		return c.String(http.StatusOK, "File berhasil diupload")
	})
}
