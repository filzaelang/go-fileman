package models_menu

import (
	"database/sql"
	"file-manager/db"
)

func UpdateMenu(payload UpdateMenuPayload) error {

	// if folder inside subfolder
	if payload.Type == SUBFOLDER_CHILD {
		_, err := db.DB_DEV.Exec(`
		update folder_list
		set [name] = @name
		where folderoid = @folderoid
		`, sql.Named("name", payload.Name),
			sql.Named("folderoid", payload.Folderoid))

		return err
	}

	// if folder inside budept
	if payload.Type == BUDEPTFOLDER_LAST_CHILD {
		_, err := db.DB_DEV.Exec(`
		update dept_list
		set name = @name
		  , lastupdateuser = @user
		  , lastupdatetime = getdate()
		where divoid = @divoid
			and deptoid = @deptoid
	`, sql.Named("name", payload.Name),
			sql.Named("user", payload.User),
			sql.Named("divoid", payload.Divoid),
			sql.Named("deptoid", payload.Deptoid))

		return err
	}

	return nil
}
