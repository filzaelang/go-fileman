package model_file

import (
	"database/sql"
	"file-manager/db"
	"file-manager/helpers"
	"fmt"
	"os"
	"path/filepath"
)

func FileDownloadHarian(id int) (string, string, string, error) {
	var title, titlehead, foldertype, divname, deptname string
	var folderoid, divoid, deptoid int

	row := db.DB_DEV.QueryRow(`
		select top 1 [name], headfolder, type, folderoid
		from folder_list
		where folderoid in
			(select folderoid
			from file_list
			where fileoid = @id)`,
		sql.Named("id", id),
	)

	err := row.Scan(&title, &titlehead, &foldertype, &folderoid)
	if err != nil {
		return "", "", "", err
	}

	row = db.DB_DEV.QueryRow(`
		select top 1 divoid, deptoid from file_list where fileoid = @id
		`, sql.Named("id", id),
	)

	err = row.Scan(&divoid, &deptoid)
	if err != nil {
		return "", "", "", err
	}

	if foldertype == "budeptfolder" || foldertype == "bufolder" {
		if divoid != 0 {
			row := db.DB_MIS.QueryRow(`
				select top 1 divname from QL_mstdivision where divoid = @divoid
			`, sql.Named("divoid", divoid))

			err := row.Scan(&divname)
			if err != nil {
				return "", "", "", err
			}
		}

		if deptoid != 0 {
			row := db.DB_DEV.QueryRow(`
				select top 1 [name] from dept_list where divoid=@divoid and deptoid=@deptoid
			`, sql.Named("divoid", divoid),
				sql.Named("deptoid", deptoid))

			err := row.Scan(&deptname)
			if err != nil {
				return "", "", "", err
			}
		}
	}

	titlehead = helpers.CharReplace(titlehead)
	title = helpers.CharReplace(title)

	if foldertype == "budeptfolder" {
		divname = helpers.CharReplace(divname)
		deptname = helpers.CharReplace(deptname)
	} else if foldertype == "bufolder" {
		divname = helpers.CharReplace(divname)
	}

	var fileurl string
	row = db.DB_DEV.QueryRow("select top 1 fileurl from file_list where fileoid = @id", sql.Named("id", id))
	err = row.Scan(&fileurl)
	if err != nil {
		return "", "", "", err
	}

	var path string
	switch foldertype {
	case "subfolder":
		path = filepath.Join("C:/FileManager", titlehead, title, fileurl) // fmt.Sprintf("C:/FileManager/%s/%s/%s", titlehead, title, fileurl)
	case "headfolder":
		path = filepath.Join("C:/FileManager", title, fileurl) // fmt.Sprintf("C:/FileManager/%s/%s", title, fileurl)
	case "budeptfolder":
		path = filepath.Join("C:/FileManager", title, divname, deptname, fileurl) // fmt.Sprintf("C:/FileManager/%s/%s/%s/%s", title, divname, deptname, fileurl)
	case "bufolder":
		path = filepath.Join("C:/FileManager", title, divname, fileurl) // fmt.Sprintf("C:/FileManager/%s/%s/%s", title, divname, fileurl)
	default:
		return "", "", "", fmt.Errorf("unknown foldertype")
	}

	content, err := os.Open(path)
	if err != nil {
		return "", "", "", err
	}
	defer content.Close()

	ext := filepath.Ext(fileurl)

	return path, fileurl, ext, nil
}

// filepathStr := `C:\FileManager\5_dip.pdf`
// filename := filepath.Base(filepathStr)

// return filepathStr, filename, nil
