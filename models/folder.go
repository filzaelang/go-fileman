package models

import (
	"database/sql"
	"file-manager/db"
)

type FolderData struct {
	Title            string `json:"title"`
	User             string `json:"user"`
	Lap_harisan_saya string `json:"lap_harisan_saya"`
	Folderhidebudept string `json:"folderhidebudept"`
	Divoid           int    `json:"divoid"`
	Deptoid          int    `json:"deptoid"`
	Roleid           int    `json:"roleid"`
}

func Folder(id, divoid, deptoid int) (FolderData, error) {
	var result FolderData
	var title, titlehead, foldertype, folderhidebudept, divname, divzip, deptname string

	row := db.DB.QueryRow(`
		select top 1 
			  foldername
			, headfolder
			, type
			, folderhidebudept 
		from itg_folder where folderoid = @id`,
		sql.Named("id", id),
	)

	err := row.Scan(&title, &titlehead, &foldertype, &folderhidebudept)
	if err != nil {
		return result, err
	}

	if foldertype == "budeptfolder" || foldertype == "bufolder" {
		if divoid != 0 {
			row = db.DB_MIS.QueryRow(`
				select top 1 divname, divzip from QL_mstdivision where divoid = @divoid`,
				sql.Named("divoid", divoid),
			)

			err = row.Scan(&divname, &divzip)
			if err != nil {
				return result, err
			}
		}
		if deptoid != 0 {
			row = db.DB.QueryRow(`
				select top 1 deptname from itg_mstdept where divoid=@divoid and deptoid=@deptoid`,
				sql.Named("divoid", divoid),
				sql.Named("deptoid", deptoid),
			)

			err = row.Scan(&deptname)
			if err != nil {
				return result, err
			}
		}
	}

	if foldertype == "subfolder" {
		result.Title = titlehead + " -> " + title
	} else if foldertype == "headfolder" {
		result.Title = title
	} else if foldertype == "budeptfolder" {
		result.Title = title + " -> " + divname + " -> " + deptname
	} else if foldertype == "bufolder" {
		result.Title = title + " -> " + divname
	}

	result.User = "admin"            //sementara dummy
	result.Lap_harisan_saya = "asdf" // sementara dummy
	result.Folderhidebudept = folderhidebudept
	result.Divoid = divoid
	result.Deptoid = deptoid
	result.Roleid = 1 // sementara dummy

	return result, nil
}
