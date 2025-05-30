package models

import (
	"file-manager/db"
	"file-manager/helpers"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type FileItem struct {
	Fileoid        int       `json:"fileoid"`
	Divoid         int       `json:"divoid"`
	Deptoid        int       `json:"deptoid"`
	Leveloid       int       `json:"leveloid"`
	Folderoid      int       `json:"folderoid"`
	Filename       string    `json:"filename"`
	Fileurl        string    `json:"fileurl"`
	Createuser     string    `json:"createuser"`
	Createtime     time.Time `json:"createtime"`
	Lastupdatetime time.Time `json:"lastupdatetime"`
	Divzip         string    `json:"divzip"`
	Filenumber     string    `json:"filenumber"`
	Filerevdate    time.Time `json:"filerevdate"`
	Fileoldnumber  string    `json:"fileoldnumber"`
	Filevisible    bool      `json:"filevisible"`
}

type ItgFile struct {
	FileURL string
}

func GetFile() ([]FileItem, error) {
	rows, err := db.DB.Query(`
		select 
			fileoid
		  , divoid
		  , deptoid
		  , leveloid
		  , folderoid
		  , filename
		  , fileurl
		  , createuser
		  , createtime
		  , lastupdatetime
		  , divzip
		  , filenumber
		  , filerevdate
		  , fileoldnumber
		  , filevisible  
		from itg_file
		where folderoid = 63
			and filevisible = 'True'
		order by 
		case
			when fileurl LIKE '%.htm' then 4
			when fileurl LIKE '%.pps' then 3
			when fileurl LIKE '%.mp4' then 2
			when fileurl LIKE '%.pdf' then 1
		end desc,
		lastupdatetime desc
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []FileItem

	for rows.Next() {
		var item FileItem
		if err := rows.Scan(
			&item.Fileoid,
			&item.Divoid,
			&item.Deptoid,
			&item.Leveloid,
			&item.Folderoid,
			&item.Filename,
			&item.Fileurl,
			&item.Createuser,
			&item.Createtime,
			&item.Lastupdatetime,
			&item.Divzip,
			&item.Filenumber,
			&item.Filerevdate,
			&item.Fileoldnumber,
			&item.Filevisible,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func FileDownloadHarian(id int) (string, string, error) {
	// var title, titlehead, foldertype, divname, deptname string
	// var folderoid, divoid, deptoid int

	// row := db.DB.QueryRow(`
	// 	select top 1 foldername, headfolder, type, folderoid
	// 	from itg_folder
	// 	where folderoid in
	// 		(select folderoid
	// 		from itg_file
	// 		where fileoid = @id)
	// 	)`,
	// 	sql.Named("id", id),
	// )

	// err := row.Scan(&title, &titlehead, &foldertype, &folderoid)
	// if err != nil {
	// 	return "", err
	// }

	// row = db.DB.QueryRow(`
	// 	select top 1 divoid, deptoid from itg_file where fileoid = @id
	// 	`, sql.Named("id", id),
	// )

	// err = row.Scan(&divoid, &deptoid)
	// if err != nil {
	// 	return "", err
	// }

	// if foldertype == "budeptfolder" || foldertype == "bufolder" {
	// 	if divoid != 0 {
	// 		row := db.DB_MIS.QueryRow(`
	// 			select top 1 divname from QL_mstdivision where divoid = @divoid
	// 		`, sql.Named("divoid", divoid))

	// 		err := row.Scan(&divname)
	// 		if err != nil {
	// 			return "", err
	// 		}
	// 	}

	// 	if deptoid != 0 {
	// 		row := db.DB.QueryRow(`
	// 			select top 1 * from itg_mstdept where divoid=@divoid and deptoid=@deptoid
	// 		`, sql.Named("divoid", divoid),
	// 			sql.Named("deptoid", deptoid))

	// 		err := row.Scan(&deptname)
	// 		if err != nil {
	// 			return "", err
	// 		}
	// 	}
	// }

	// titlehead = helpers.CharReplace(titlehead)
	// title = helpers.CharReplace(title)

	// if foldertype == "budeptfolder" {
	// 	divname = helpers.CharReplace(divname)
	// 	deptname = helpers.CharReplace(deptname)
	// } else if foldertype == "bufolder" {
	// 	divname = helpers.CharReplace(divname)
	// }

	// var fileData string
	// row = db.DB.QueryRow("select top 1 fileurl from itg_file where fileoid = @id", sql.Named("id", id))
	// err = row.Scan(&fileData)
	// if err != nil {
	// 	return "", err
	// }

	// content, err := os.ReadFile(filepathStr)
	// if err != nil {
	// 	return "", nil
	// }
	filepathStr := `C:\FileManager\5_dip.pdf`
	filename := filepath.Base(filepathStr)

	return filepathStr, filename, nil
}

func Upload(fileHeader *multipart.FileHeader, c echo.Context) (string, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return "Gagal membuka file", err
	}
	defer src.Close()

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))

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
		return "Gagal menyimpan file", err
	}

	if _, err = io.Copy(dst, src); err != nil {
		dst.Close()
		return "Gagal menyalin file", err
	}
	dst.Close()

	compressedPath := filepath.Join("C:\\FileManager", (documentName + ext))

	if ext == ".pdf" {
		if err := helpers.CompressPdf(targetPath, compressedPath); err != nil {
			return "Gagal mengompresi PDF", err
		}
	} else if ext == ".png" {
		if err := helpers.CompressPng(targetPath, compressedPath); err != nil {
			return "Gagal mengompresi PNG", err
		}
	}

	// Hapus file asli
	if err := os.Remove(targetPath); err != nil {
		return "Gagal menghapus file asli", err
	}

	return "File berhasil diupload", nil
}

// userdeptoid := 0

// var filepathStr string = ""

// // Buat full path berdasarkan foldertype
// if foldertype == "subfolder" {
// 	filepathStr = filepath.Join("C:\\FileManager", titlehead, title, fileData)
// } else if foldertype == "headfolder" {
// 	filepathStr = filepath.Join("C:\\FileManager", title, fileData)
// } else if foldertype == "budeptfolder" {
// 	filepathStr = filepath.Join("C:\\FileManager", title, divname, deptname, fileData)
// } else if foldertype == "bufolder" {
// 	filepathStr = filepath.Join("C:\\FileManager", title, divname, fileData)
// }

// // Buka dan baca isi file
// content, err := os.ReadFile(filepathStr)
// if err != nil {
// 	return "", nil
// }

// return filepathStr, content, nil
