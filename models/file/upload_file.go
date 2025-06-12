package model_file

import (
	"database/sql"
	"file-manager/db"
	"file-manager/helpers"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func Upload(fileHeader *multipart.FileHeader, c echo.Context, folderoid int, divoid int, deptoid int) (string, error, string) {

	var upload_path string
	// check folder_type etc
	var title, titlehead, foldertype, folderhidebudept, divname, deptname string
	row := db.DB_DEV.QueryRow(`
		select top 1 
			  [name]
			, headfolder
			, type
			, folderhidebudept 
		from folder_list where folderoid = @folderoid`,
		sql.Named("folderoid", folderoid),
	)

	err := row.Scan(&title, &titlehead, &foldertype, &folderhidebudept)
	if err != nil {
		return "", err, ""
	}

	if foldertype == "budeptfolder" || foldertype == "bufolder" {
		if divoid != 0 {
			row = db.DB_MIS.QueryRow(`
				select top 1 divname from QL_mstdivision where divoid = @divoid`,
				sql.Named("divoid", divoid),
			)

			err = row.Scan(&divname)
			if err != nil {
				return "", err, ""
			}
		}
		if deptoid != 0 {
			row = db.DB_DEV.QueryRow(`
				select top 1 name from dept_list where divoid=@divoid and deptoid=@deptoid`,
				sql.Named("divoid", divoid),
				sql.Named("deptoid", deptoid),
			)

			err = row.Scan(&deptname)
			if err != nil {
				return "", err, ""
			}
		}
	}

	var redirectUrl string

	src, err := fileHeader.Open()
	if err != nil {
		redirectUrl = ErrorRedirect(foldertype, folderoid, divoid, deptoid)
		return redirectUrl, err, "UPLOAD FAILED! No file uploaded."
	}
	defer src.Close()

	// ext := strings.ToLower(filepath.Ext(fileHeader.Filename))

	fileraw := helpers.PrepFilename(fileHeader.Filename)
	file := helpers.NormalizeFilename(fileraw)
	file = strings.ReplaceAll(file, ",", "")

	var query string
	args := []interface{}{}
	var filenumber string

	if foldertype == "headfolder" || foldertype == "subfolder" {
		query = "select top 1 filenumber from file_list where folderoid = @folderoid and fileurl=@fileurl"
		args = append(args, sql.Named("folderoid", folderoid))
		args = append(args, sql.Named("fileurl", file))
	} else if foldertype == "bufolder" {
		query = "select top 1 filenumber from file_list where folderoid = @folderoid and divoid = @divoid and fileurl=@fileurl"
		args = append(args, sql.Named("folderoid", folderoid))
		args = append(args, sql.Named("divoid", divoid))
		args = append(args, sql.Named("fileurl", file))
	} else if foldertype == "budeptfolder" {
		query = "select top 1 filenumber from file_list where folderoid = @folderoid and divoid = @divoid and deptoid = @deptoid and fileurl = @fileurl"
		args = append(args, sql.Named("folderoid", folderoid))
		args = append(args, sql.Named("divoid", divoid))
		args = append(args, sql.Named("deptoid", deptoid))
		args = append(args, sql.Named("fileurl", file))
	}

	row = db.DB_DEV.QueryRow(query, args...)
	err = row.Scan(&filenumber)

	if err != sql.ErrNoRows {
		errorUploadMessage := fmt.Sprintf("UPLOAD FAILED! file uploaded is exist in this folder with number %s", filenumber)
		redirectUrl = ErrorRedirect(foldertype, folderoid, divoid, deptoid)
		return redirectUrl, nil, errorUploadMessage
	} else if err != nil {
		return "", err, "Filenumber doesn't exist"
	}

	fileNumberInput := c.FormValue("document_number")

	row = db.DB_DEV.QueryRow(
		"select top 1 filenumber from file_list where folderoid <> 0 and filenumber = @filenumber",
		sql.Named("filenumber", fileNumberInput))
	err = row.Scan(&filenumber)

	if err != sql.ErrNoRows {
		errorUploadMessage := fmt.Sprintf("UPLOAD FAILED! file number %s is already exist. please check in All Files", fileNumberInput)
		redirectUrl = ErrorRedirect(foldertype, folderoid, divoid, deptoid)
		return redirectUrl, nil, errorUploadMessage
	} else if err != nil {
		return "", err, ""
	}

	titlehead = helpers.CharReplace(titlehead)
	title = helpers.CharReplace(title)
	divname = helpers.CharReplace(divname)
	if foldertype == "budeptfolder" {
		deptname = helpers.CharReplace(deptname)
	}

	// $config['allowed_types'] = 'xls|xlsx|doc|docx|ppt|pptx|pdf|zip|rar|txt';
	// $config['max_size']     = '102400000';

	upload_path, err = UploadPath(foldertype, title, titlehead, divname, deptname)
	if err != nil {
		return "", err, ""
	}

	// return filenumber, nil, ""
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

func UploadPath(foldertype, title, titlehead, divname, deptname string) (string, error) {
	var upload_path string
	if foldertype == "headfolder" {
		dirpath := filepath.Join("C:/FileManager", title)
		if _, err := os.Stat(dirpath); os.IsNotExist(err) {
			err := os.MkdirAll(dirpath, os.ModePerm)
			if err != nil {
				return "", err
			}
		}
		upload_path = dirpath + string(os.PathSeparator) //fmt.Sprintf("C:/FileManager/%s/", title)
	} else if foldertype == "subfolder" {
		dirpath := filepath.Join("C:/FileManager", titlehead, title)
		if _, err := os.Stat(dirpath); os.IsNotExist(err) {
			err := os.MkdirAll(dirpath, os.ModePerm)
			if err != nil {
				return "", err
			}
		}
		upload_path = dirpath + string(os.PathSeparator) //fmt.Sprintf("C:/FileManager/%s/%s/", titlehead, title)
	} else if foldertype == "bufolder" {
		dirpath := filepath.Join("C:/FileManager", title, divname)
		if _, err := os.Stat(dirpath); os.IsNotExist(err) {
			err := os.MkdirAll(dirpath, os.ModePerm)
			if err != nil {
				return "", err
			}
		}
		upload_path = dirpath + string(os.PathSeparator) //fmt.Sprintf("C:/FileManager/%s/%s/", title, divname)
	} else if foldertype == "budeptfolder" {
		dirpath := filepath.Join("C:/FileManager", title, divname, deptname)
		if _, err := os.Stat(dirpath); os.IsNotExist(err) {
			err := os.MkdirAll(dirpath, os.ModePerm)
			if err != nil {
				return "", err
			}
		}

		upload_path = dirpath + string(os.PathSeparator) //fmt.Sprintf("C:/FileManager/%s/%s/%s/", title, divname, deptname)
	}

	return upload_path, nil
}

func ErrorRedirect(foldertype string, folderoid int, divoid int, deptoid int) string {
	var redirectUrl string = ""
	if foldertype == "headfolder" || foldertype == "subfolder" {
		redirectUrl = "/folder/" + strconv.Itoa(folderoid) + "/0/0"
	} else if foldertype == "bufolder" {
		redirectUrl = fmt.Sprintf("/folder/%s/%s/0", folderoid, divoid)
	} else if foldertype == "budeptfolder" {
		redirectUrl = fmt.Sprintf("/folder/%s/%s/%s", folderoid, divoid, deptoid)
	}

	return redirectUrl
}
