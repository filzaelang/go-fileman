package model_file

import (
	"database/sql"
	"file-manager/db"
	"file-manager/helpers"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func UploadFile(fileHeader *multipart.FileHeader, c echo.Context) (string, string, error) {

	filenameInput := c.FormValue("document_name")
	filenumberInput := c.FormValue("document_number")
	filerevnumberInput := c.FormValue("revision_number")
	filerevdateInput := c.FormValue("revision_date")
	filerevdate, err := time.Parse("2006-01-02", filerevdateInput)
	if err != nil {
		return "", "", err
	}
	folderoid, err := strconv.Atoi(c.FormValue("folderoid"))
	if err != nil {
		return "", "", err
	}
	divoid, err := strconv.Atoi(c.FormValue("divoid"))
	if err != nil {
		return "", "", err
	}
	deptoid, err := strconv.Atoi(c.FormValue("deptoid"))
	if err != nil {
		return "", "", err
	}

	transaction, err := db.DB_DEV.Begin()
	if err != nil {
		return "", "", err
	}
	defer func() {
		if err != nil {
			transaction.Rollback()
		}
	}()

	// check folder_type etc
	var title, titlehead, foldertype, folderhidebudept, divname, deptname string
	row := transaction.QueryRow(`
		select top 1 
			  [name]
			, headfolder
			, type
			, folderhidebudept 
		from folder_list where folderoid = @folderoid`,
		sql.Named("folderoid", folderoid),
	)

	err = row.Scan(&title, &titlehead, &foldertype, &folderhidebudept)
	if err != nil {
		return "", "", err
	}

	if foldertype == "budeptfolder" || foldertype == "bufolder" {
		if divoid != 0 {
			row = db.DB_MIS.QueryRow(`
				select top 1 divname from QL_mstdivision where divoid = @divoid`,
				sql.Named("divoid", divoid),
			)

			err = row.Scan(&divname)
			if err != nil {
				return "", "", err
			}
		}
		if deptoid != 0 {
			row = transaction.QueryRow(`
				select top 1 name from dept_list where divoid=@divoid and deptoid=@deptoid`,
				sql.Named("divoid", divoid),
				sql.Named("deptoid", deptoid),
			)

			err = row.Scan(&deptname)
			if err != nil {
				return "", "", err
			}
		}
	}

	var redirectUrl string

	src, err := fileHeader.Open()
	if err != nil {
		redirectUrl = ResponseRedirect(foldertype, folderoid, divoid, deptoid)
		return redirectUrl, "UPLOAD FAILED! No file uploaded.", err
	}
	defer src.Close()

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))

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

	row = transaction.QueryRow(query, args...)
	err = row.Scan(&filenumber)

	if err != sql.ErrNoRows {
		errorUploadMessage := fmt.Sprintf("UPLOAD FAILED! file uploaded is exist in this folder with number %s", filenumber)
		redirectUrl = ResponseRedirect(foldertype, folderoid, divoid, deptoid)
		return redirectUrl, errorUploadMessage, nil
	}

	fileNumberInput := c.FormValue("document_number")

	row = transaction.QueryRow(
		"select top 1 filenumber from file_list where folderoid <> 0 and filenumber = @filenumber",
		sql.Named("filenumber", fileNumberInput))
	err = row.Scan(&filenumber)

	if err != sql.ErrNoRows {
		errorUploadMessage := fmt.Sprintf("UPLOAD FAILED! file number %s is already exist. please check in All Files", fileNumberInput)
		redirectUrl = ResponseRedirect(foldertype, folderoid, divoid, deptoid)
		return redirectUrl, errorUploadMessage, nil
	}

	titlehead = helpers.CharReplace(titlehead)
	title = helpers.CharReplace(title)
	divname = helpers.CharReplace(divname)
	if foldertype == "budeptfolder" {
		deptname = helpers.CharReplace(deptname)
	}

	// get last fileoid
	var lastFileoid int
	rows := transaction.QueryRow("select top 1 fileoid from file_list order by fileoid desc")
	err = rows.Scan(&lastFileoid)
	if err != nil {
		return "", "", err
	}

	newFileoid := lastFileoid + 1

	_, err = transaction.Exec(`
		insert into file_list (
			fileoid
		  , divoid
		  , deptoid
		  , leveloid
		  , folderoid
		  , filename
		  , fileurl
		  , createuser
		  , createtime
		  , lastupdateuser
		  , lastupdatetime
		  , filenumber
		  , filerevnumber
		  , filerevdate
		  , fileoldnumber
		  , filevisible
		) values (
		 	@fileoid
		  , @divoid
		  , @deptoid
		  , @leveloid
		  , @folderoid
		  , @filename
		  , @fileurl
		  , @user
		  , getdate()
		  , @user
		  , getdate()
		  , @filenumber
		  , @filerevnumber
		  , @filerevdate
		  , @fileoldnumber
		  , @filevisible
		)
		`, sql.Named("fileoid", newFileoid),
		sql.Named("divoid", divoid),
		sql.Named("deptoid", deptoid),
		sql.Named("leveloid", 0), //Sementara diisi nol => $this->session->userdata('leveloid')
		sql.Named("folderoid", folderoid),
		sql.Named("filename", filenameInput),
		sql.Named("fileurl", file), //Nanti di research lagi
		sql.Named("user", "admin"), //Nanti dibenarkan setelah auth selesai
		sql.Named("filenumber", filenumberInput),
		sql.Named("filerevnumber", filerevnumberInput),
		sql.Named("filerevdate", filerevdate),
		sql.Named("fileoldnumber", ""),
		sql.Named("filevisible", "True"))

	if err != nil {
		return "", "", err
	}

	// $config['allowed_types'] = 'xls|xlsx|doc|docx|ppt|pptx|pdf|zip|rar|txt';
	// $config['max_size']     = '102400000';

	upload_path, err := UploadPath(foldertype, title, titlehead, divname, deptname)
	if err != nil {
		return "", "", err
	}

	fullFilePath := filepath.Join(upload_path, file)

	if ext == ".pdf" || ext == ".png" {
		timestamp := time.Now().Format("20060102_150405")
		safeFileName := fmt.Sprintf("%s_%s_%s_%s%s",
			filenumberInput,
			filenameInput,
			filerevnumberInput,
			timestamp,
			filepath.Ext(fileHeader.Filename),
		)
		targetPath := filepath.Join("C:/FileManager", safeFileName)

		dst, err := os.Create(targetPath)
		if err != nil {
			return "Gagal menyimpan file", "", err
		}

		if _, err = io.Copy(dst, src); err != nil {
			dst.Close()
			return "Gagal menyalin file", "", err
		}
		dst.Close()

		if ext == ".pdf" {
			if err := helpers.CompressPdf(targetPath, fullFilePath); err != nil {
				return "Gagal mengompresi PDF", "", err
			}
		} else if ext == ".png" {
			if err := helpers.CompressPng(targetPath, fullFilePath); err != nil {
				return "Gagal mengompresi PNG", "", err
			}
		}

		// Hapus file asli
		if err := os.Remove(targetPath); err != nil {
			return "Gagal menghapus file asli", "", err
		}
	} else {
		dst, err := os.Create(fullFilePath)
		if err != nil {
			return "Gagal menyimpan file", "", err
		}

		if _, err = io.Copy(dst, src); err != nil {
			dst.Close()
			return "Gagal menyalin file", "", err
		}
		dst.Close()
	}

	transaction.Commit()

	url_redirect := ResponseRedirect(foldertype, folderoid, divoid, deptoid)

	return url_redirect, "", nil
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

		upload_path = fmt.Sprintf("C:/FileManager/%s/%s/%s/", title, divname, deptname)
	}

	return upload_path, nil
}

func ResponseRedirect(foldertype string, folderoid int, divoid int, deptoid int) string {
	var redirectUrl string = ""
	if foldertype == "headfolder" || foldertype == "subfolder" {
		redirectUrl = fmt.Sprintf("/folder/%s/0/0", strconv.Itoa(folderoid))
	} else if foldertype == "bufolder" {
		redirectUrl = fmt.Sprintf("/folder/%s/%s/0", strconv.Itoa(folderoid), strconv.Itoa(divoid))
	} else if foldertype == "budeptfolder" {
		redirectUrl = fmt.Sprintf("/folder/%s/%s/%s", strconv.Itoa(folderoid), strconv.Itoa(divoid), strconv.Itoa(deptoid))
	}

	return redirectUrl
}
