package models_menu

import (
	"database/sql"
	"file-manager/db"
)

func GetOneMenu(payload UpdateMenuPayload) (string, error) {
	var name string
	var err error

	if payload.Type == BUDEPTFOLDER_LAST_CHILD {
		name, err = InsideBUdeptChild(payload)
	} else {
		name, err = HeadAndSubfolder(payload)
	}

	return name, err
}

func HeadAndSubfolder(payload UpdateMenuPayload) (string, error) {
	var name = ""
	row := db.DB_DEV.QueryRow(`
			select top 1 name
			from folder_list
			where folderoid = @folderoid
			`, sql.Named("folderoid", payload.Folderoid))

	err := row.Scan(&name)
	return name, err
}

func InsideBUdeptChild(payload UpdateMenuPayload) (string, error) {
	var name = ""
	row := db.DB_DEV.QueryRow(`
		select top 1 name 
		from dept_list 
		where divoid = @divoid
			and deptoid = @deptoid
		`, sql.Named("divoid", payload.Divoid),
		sql.Named("deptoid", payload.Deptoid))

	err := row.Scan(&name)

	return name, err
}
