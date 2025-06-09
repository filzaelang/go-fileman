package models_menu

import (
	"database/sql"
	"file-manager/db"
)

func GetOneMenu(payload UpdateMenuPayload) (string, error) {
	var name = ""
	// if menu is inside subfolder
	if payload.Type == SUBFOLDER_CHILD {
		row := db.DB_DEV.QueryRow(`
			select top 1 name
			from folder_list
			where folderoid = @folderoid
			`, sql.Named("folderoid", payload.Folderoid))

		err := row.Scan(&name)

		if err != nil {
			return name, nil
		}

		return name, nil
	}
	// if menu is inside budeptchild
	if payload.Type == BUDEPTFOLDER_LAST_CHILD {
		row := db.DB_DEV.QueryRow(`
		select top 1 name 
		from dept_list 
		where divoid = @divoid
			and deptoid = @deptoid
		`, sql.Named("divoid", payload.Divoid),
			sql.Named("deptoid", payload.Deptoid))

		err := row.Scan(&name)

		if err != nil {
			return name, err
		}

		return name, nil
	}

	return name, nil
}
