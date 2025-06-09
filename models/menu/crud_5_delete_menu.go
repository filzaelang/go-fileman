package models_menu

import (
	"database/sql"
	"file-manager/db"
)

func DeleteMenu(payload DeleteMenuPayload) error {
	transaction, err := db.DB_DEV.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			transaction.Rollback()
		}
	}()

	// if folder inside subfolder
	if payload.Type == SUBFOLDER_CHILD {
		_, err = transaction.Exec(`
		delete from folder_list where folderoid = @folderoid
		`, sql.Named("folderoid", payload.Folderoid))

		if err != nil {
			return err
		}

		return transaction.Commit()
	}

	// if folder inside budept
	if payload.Type == BUDEPTFOLDER_LAST_CHILD {
		_, err = transaction.Exec(`
		delete from folder_dept 
		where folderoid = @folderoid
			and divoid = @divoid
			and deptoid = @deptoid`,
			sql.Named("folderoid", payload.Folderoid),
			sql.Named("divoid", payload.Divoid),
			sql.Named("deptoid", payload.Deptoid),
		)
		if err != nil {
			return err
		}

		_, err = transaction.Exec(`
		delete from dept_list
		where divoid = @divoid
			and deptoid = @deptoid
	`, sql.Named("divoid", payload.Divoid),
			sql.Named("deptoid", payload.Deptoid))
		if err != nil {
			return err
		}

		return transaction.Commit()
	}

	return nil
}
