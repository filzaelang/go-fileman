package models_menu

import (
	"database/sql"
	"file-manager/db"
)

func UpdateMenu(payload UpdateMenuPayload) error {
	var err error

	// if folder is headfolder
	if payload.Type == HEADFOLDER {
		err = UpdateHeadfolder(payload)
	}

	// if folder is subfolder
	if payload.Type == SUBFOLDER {
		err = UpdateSubfolder(payload)
	}

	// if folder is bufolder
	if payload.Type == BUFOLDER {
		err = UpdateBUFolder(payload)
	}

	// if folder inside subfolder
	if payload.Type == SUBFOLDER_CHILD {
		err = UpdateBUSubfolderchild(payload)
	}

	// if folder inside budept
	if payload.Type == BUDEPTFOLDER_LAST_CHILD {
		err = UpdateBUdeptlastchild(payload)
	}

	return err
}

func UpdateHeadfolder(payload UpdateMenuPayload) error {
	_, err := db.DB_DEV.Exec(`
		update folder_list
		set headfolder = @name
		  , [name] = @name
		  , lastupdateuser = @user
		  , lastupdatetime = getdate()
		where headfolder = @previous_headfolder
			and type = 'headfolder' 
		`, sql.Named("name", payload.Name),
		sql.Named("user", payload.User),
		sql.Named("previous_headfolder", payload.Headfolder))
	return err
}

func UpdateSubfolder(payload UpdateMenuPayload) error {
	_, err := db.DB_DEV.Exec(`
		update folder_list
		set headfolder = @headfolder
		  , lastupdateuser = @user
		  , lastupdatetime = getdate()
		where headfolder = @previous_headfolder
			and type = 'subfolder'
		`, sql.Named("headfolder", payload.Name),
		sql.Named("user", payload.User),
		sql.Named("previous_headfolder", payload.Headfolder))
	return err
}

func UpdateBUSubfolderchild(payload UpdateMenuPayload) error {
	var previousName string

	rows := db.DB_DEV.QueryRow(`select top 1 name from folder_list where folderoid = @folderoid`,
		sql.Named("folderoid", payload.Folderoid))

	err := rows.Scan(&previousName)
	if err != nil {
		return err
	}

	_, err = db.DB_DEV.Exec(`
		update folder_list
		set [name] = @name
		  , lastupdateuser = @user
		  , lastupdatetime = getdate()
		where name = @previous_name
			and type = 'subfolder'
		`, sql.Named("name", payload.Name),
		sql.Named("user", payload.User),
		sql.Named("previous_name", previousName))

	return err
}

func UpdateBUFolder(payload UpdateMenuPayload) error {
	_, err := db.DB_DEV.Exec(`
		update folder_list
		set headfolder = @name
		  , [name] = @name
		  , lastupdateuser = @user
		  , lastupdatetime = getdate()
		where headfolder = @previous_headfolder
			and type = 'bufolder' 
		`, sql.Named("name", payload.Name),
		sql.Named("user", payload.User),
		sql.Named("previous_headfolder", payload.Headfolder))
	return err
}

func UpdateBUdeptlastchild(payload UpdateMenuPayload) error {
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
