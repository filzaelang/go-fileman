package model_file

import (
	"database/sql"
	"file-manager/db"
	"file-manager/helpers"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

func Upload(fileHeader *multipart.FileHeader, c echo.Context, folderoid int, divoid int, deptoid int) (string, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return "Gagal membuka file", err //nanti dibenahi yang jelas redirect
	}
	defer src.Close()

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))

	fileraw := helpers.PrepFilename(fileHeader.Filename)
	file := helpers.NormalizeFilename(fileraw)
	file = strings.ReplaceAll(file, ",", "")

	// check folder_type
	var foldertype string
	row := db.DB.QueryRow(`
		select top 1 type
		from folder_list where folderoid = @folderoid`,
		sql.Named("folderoid", folderoid),
	)

	err = row.Scan(&foldertype)
	if err != nil {
		return "", err
	}

	var query string
	args := []interface{}{}
	var count int

	if foldertype == "headfolder" || foldertype == "subfolder" {
		query = "select top 1 * from file_list where folderoid = @folderoid and fileurl=@fileurl"
		args = append(args, sql.Named("folderoid", folderoid))
		args = append(args, sql.Named("fileurl", file))
	} else if foldertype == "bufolder" {
		query = "select top 1 * from file_list where folderoid = @folderoid and divoid = @divoid and fileurl=@fileurl"
		args = append(args, sql.Named("folderoid", folderoid))
		args = append(args, sql.Named("divoid", divoid))
		args = append(args, sql.Named("fileurl", file))
	} else if foldertype == "budeptfolder" {
		query = "select top 1 * from file_list where folderoid = $folderoid and divoid = $divoid and deptoid = @deptoid and fileurl = @fileurl"
		args = append(args, sql.Named("folderoid", folderoid))
		args = append(args, sql.Named("divoid", divoid))
		args = append(args, sql.Named("deptoid", deptoid))
		args = append(args, sql.Named("fileurl", file))
	}

	countRows := db.DB_DEV.QueryRow(query, args...)
	err = countRows.Scan(&count)
	if err != nil {
		return "", err
	}

	if count == 0 {

	} else {
		if foldertype == "subfolder" {

		}
	}

	// // get metadata from form
	// ducumentNumber := c.FormValue("document_number")
	// documentName := c.FormValue("document_name")
	// revisionNumber := c.FormValue("revision_number")
	// revisionDate := c.FormValue("revision_date")

	// // Save file to C:\FileManager\ with unique name
	// timestamp := time.Now().Format("20060102_150405")
	// safeFileName := fmt.Sprintf("%s_%s_%s_%s%s",
	// 	ducumentNumber,
	// 	documentName,
	// 	revisionNumber,
	// 	timestamp,
	// 	filepath.Ext(fileHeader.Filename),
	// )
	// targetPath := filepath.Join("C:\\FileManager", safeFileName)

	// dst, err := os.Create(targetPath)
	// if err != nil {
	// 	return "Gagal menyimpan file", err
	// }

	// if _, err = io.Copy(dst, src); err != nil {
	// 	dst.Close()
	// 	return "Gagal menyalin file", err
	// }
	// dst.Close()

	// compressedPath := filepath.Join("C:\\FileManager", (documentName + ext))

	// if ext == ".pdf" {
	// 	if err := helpers.CompressPdf(targetPath, compressedPath); err != nil {
	// 		return "Gagal mengompresi PDF", err
	// 	}
	// } else if ext == ".png" {
	// 	if err := helpers.CompressPng(targetPath, compressedPath); err != nil {
	// 		return "Gagal mengompresi PNG", err
	// 	}
	// }

	// // Hapus file asli
	// if err := os.Remove(targetPath); err != nil {
	// 	return "Gagal menghapus file asli", err
	// }

	// return "File berhasil diupload", nil
}
