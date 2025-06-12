package models_folder

import (
	"database/sql"
	"file-manager/db"
	model_file "file-manager/models/file"
)

type FolderData struct {
	Title            string                `json:"title"`
	User             string                `json:"user"`
	Lap_harisan_saya []model_file.FileItem `json:"lap_harisan_saya"`
	Folderhidebudept string                `json:"folderhidebudept"`
	Divoid           int                   `json:"divoid"`
	Deptoid          int                   `json:"deptoid"`
	Roleid           int                   `json:"roleid"`
}

func Folder(folderoid, divoid, deptoid int) (FolderData, error) {
	// $this->updlastact();
	// $this->form_validation->set_rules('createuser', 'ID', 'required|trim');

	var result FolderData
	var title, titlehead, foldertype, folderhidebudept, divname, deptname string

	row := db.DB.QueryRow(`
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
		return result, err
	}

	if foldertype == "budeptfolder" || foldertype == "bufolder" {
		if divoid != 0 {
			row = db.DB_MIS.QueryRow(`
				select top 1 divname from QL_mstdivision where divoid = @divoid`,
				sql.Named("divoid", divoid),
			)

			err = row.Scan(&divname)
			if err != nil {
				return result, err
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

	LapHarian, err := model_file.GetFile(folderoid, divoid, deptoid)
	if err != nil {
		return result, err
	}

	result.User = "admin"               // sementara dummy
	result.Lap_harisan_saya = LapHarian // sementara dummy (diambil dari models file/getLapHarian (folderoid, divoid, deptoid))
	result.Folderhidebudept = folderhidebudept
	result.Divoid = divoid
	result.Deptoid = deptoid
	result.Roleid = 1 // sementara dummy

	return result, nil
}
