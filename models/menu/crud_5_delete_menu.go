package models_menu

import (
	"database/sql"
	"file-manager/db"
	"fmt"
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

	// if folder is headfolder
	if payload.Type == HEADFOLDER {
		_, err = transaction.Exec(`
			delete from folder_list where folderoid = @folderoid
		`, sql.Named("folderoid", payload.Folderoid))
		if err != nil {
			return err
		}

		return transaction.Commit()
	}

	// if folder is BUFOLDER
	if payload.Type == BUFOLDER {
		_, err = transaction.Exec(`
			delete from folder_bu where folderoid = @folderoid
		`, sql.Named("divoid", payload.Divoid), sql.Named("folderoid", payload.Folderoid))
		if err != nil {
			return err
		}

		_, err = transaction.Exec(`
			delete from folder_list where folderoid = @folderoid
		`, sql.Named("folderoid", payload.Folderoid))
		if err != nil {
			return err
		}

		return transaction.Commit()
	}

	// if folder inside bufolder
	if payload.Type == BUFOLDER_CHILD {
		_, err = transaction.Exec(`
			delete from folder_bu where divoid = @divoid and folderoid = @folderoid
		`, sql.Named("divoid", payload.Divoid), sql.Named("folderoid", payload.Folderoid))
		if err != nil {
			return err
		}

		return transaction.Commit()
	}

	// if folder is subfolder
	if payload.Type == SUBFOLDER {
		_, err = transaction.Exec(`
			delete from folder_list where headfolder = @headfolder
		`, sql.Named("headfolder", payload.Headfolder))
		if err != nil {
			return err
		}

		return transaction.Commit()
	}

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

	// if folder budeptfolder
	if payload.Type == BUDEPTFOLDER {
		_, err = transaction.Exec(`
			delete from folder_dept where folderoid = @folderoid
		`, sql.Named("folderoid", payload.Folderoid))
		if err != nil {
			return err
		}

		_, err = transaction.Exec(`
			delete from folder_list where folderoid = @folderoid
		`, sql.Named("folderoid", payload.Folderoid))
		if err != nil {
			return err
		}

		return transaction.Commit()
	}

	// if folder inside budept
	if payload.Type == BUDEPTFOLDER_CHILD {
		var listFolderDeptid []int
		rows, err := transaction.Query(`
			select deptoid from folder_dept where folderoid = @folderoid and divoid = @divoid
		`, sql.Named("folderoid", payload.Folderoid),
			sql.Named("divoid", payload.Divoid))
		if err != nil {
			return err
		}

		for rows.Next() {
			var deptoid int
			if err := rows.Scan(&deptoid); err != nil {
				return err
			}
			listFolderDeptid = append(listFolderDeptid, deptoid)
		}

		_, err = transaction.Exec(`
			delete from folder_dept where folderoid = @folderoid and divoid = @divoid
		`, sql.Named("folderoid", payload.Folderoid), sql.Named("divoid", payload.Divoid))
		if err != nil {
			return err
		}

		// delete in dept_list
		placeholders := ""
		args := []interface{}{}
		for i, id := range listFolderDeptid {
			if i > 0 {
				placeholders += ", "
			}
			placeholders += fmt.Sprintf("@id%d", i)
			args = append(args, sql.Named(fmt.Sprintf("id%d", i), id))
		}
		query := "delete from dept_list where deptoid in (" + placeholders + ")"

		_, err = transaction.Exec(query, args...)
		if err != nil {
			return err
		}

		return transaction.Commit()
	}

	// if folder inside budeptchild
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
