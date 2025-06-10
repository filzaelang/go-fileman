package models_menu

import (
	"database/sql"
	"file-manager/db"
	"slices"
)

func InsertMenu(payload AddMenuPayload) error {
	transaction, err := db.DB_DEV.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			transaction.Rollback()
		}
	}()

	// if folder not inside anywhere
	if payload.IsBase {
		err = AddMainMenu(payload, transaction)
		if err != nil {
			return err
		}
	}

	// if folder inside subfolder
	if payload.Type == SUBFOLDER {
		err = AddSubfolderChild(payload, transaction)
		if err != nil {
			return err
		}
	}

	//if folder inside bufolder
	if payload.Type == BUFOLDER {
		err = AddBUfolderChild(payload, transaction)
		if err != nil {
			return err
		}
	}

	// if folder inside budeptfolder
	if payload.Type == BUDEPTFOLDER {
		err = AddBUdeptfolderChild(payload, transaction)
		if err != nil {
			return err
		}
	}

	// if it is folder inside budeptfolder_child
	if payload.Type == BUDEPTFOLDER_CHILD {
		err = AddBUdeptfolderLastChild(payload, transaction)
		if err != nil {
			return err
		}
	}

	return transaction.Commit()
}

func AddMainMenu(payload AddMenuPayload, transaction *sql.Tx) error {
	var lastFolderoid int
	lastFolderoidRows := transaction.QueryRow(`select top 1 folderoid from folder_list order by folderoid desc`)
	err := lastFolderoidRows.Scan(&lastFolderoid)
	if err != nil {
		return err
	}

	var lastSeq int
	lastSeqRows := transaction.QueryRow(`select distinct top 1 seq from folder_list order by seq desc`)
	err = lastSeqRows.Scan(&lastSeq)
	if err != nil {
		return err
	}

	var folderoid = lastFolderoid + 1
	var seq = lastSeq + 1
	var folderhidebudept string
	if payload.Type == "headfolder" || payload.Type == "subfolder" {
		folderhidebudept = "Y"
	} else {
		folderhidebudept = "N"
	}

	_, err = transaction.Exec(`
		insert into folder_list (
			  folderoid
			, divoid
			, deptoid
			, leveloid
			, headfolder
			, [name]
			, divzip
			, createuser
			, createtime
			, lastupdateuser
			, lastupdatetime
			, type
			, seq
			, folderhidebudept
		) values (
			  @folderoid
			, @divoid
			, @deptoid
			, @leveloid
			, @headfolder
			, @name
			, @divzip
			, @user
			, getdate()
			, @user
			, getdate()
			, @type
			, @seq
			, @folderhidebudept
		)`, sql.Named("folderoid", folderoid),
		sql.Named("divoid", 0),
		sql.Named("deptoid", 0),
		sql.Named("leveloid", 0),
		sql.Named("headfolder", payload.Name),
		sql.Named("name", payload.Name),
		sql.Named("divzip", "CORP"), // Sementara otomatis isi "CORP"
		sql.Named("user", payload.User),
		sql.Named("type", payload.Type),
		sql.Named("seq", seq),
		sql.Named("folderhidebudept", folderhidebudept))
	return err
}

func AddSubfolderChild(payload AddMenuPayload, transaction *sql.Tx) error {
	var lastFolderoid int
	lastFolderoidRows := transaction.QueryRow(`select top 1 folderoid from folder_list order by folderoid desc`)
	err := lastFolderoidRows.Scan(&lastFolderoid)
	if err != nil {
		return err
	}

	var folderoid = lastFolderoid + 1

	var subfolderList []SubFolderList

	subfolderListRows, err := transaction.Query(`
			select distinct headfolder, seq from folder_list where type = 'subfolder' order by seq asc
		`)
	if err != nil {
		return err
	}

	for subfolderListRows.Next() {
		var item SubFolderList
		if err := subfolderListRows.Scan(&item.Headfolder, &item.Seq); err != nil {
			return err
		}
		subfolderList = append(subfolderList, item)
	}

	for _, subfolder := range subfolderList {
		_, err := transaction.Exec(`
				insert into folder_list (
					folderoid
				  , divoid
				  , deptoid
				  , leveloid
				  , headfolder
				  , [name]
				  , divzip
				  , createuser
				  , createtime
				  , lastupdateuser
				  , lastupdatetime
				  , type
				  , seq
				  , folderhidebudept
				) values (
				    @folderoid
				  , @divoid
				  , @deptoid
				  , @leveloid
				  , @headfolder
				  , @name
				  , @divzip
				  , @user
				  , getdate()
				  , @user
				  , getdate()
				  , @type
				  , @seq
				  , @folderhidebudept
				)
			`, sql.Named("folderoid", folderoid),
			sql.Named("divoid", 0),
			sql.Named("deptoid", 0),
			sql.Named("leveloid", 0),
			sql.Named("headfolder", subfolder.Headfolder),
			sql.Named("name", payload.Name),
			sql.Named("divzip", "CORP"), // Sementara otomatis isi "CORP"
			sql.Named("user", payload.User),
			sql.Named("type", payload.Type),
			sql.Named("seq", subfolder.Seq),
			sql.Named("folderhidebudept", "Y")) // Sementara otomatis isi "Y"
		if err != nil {
			return err
		}
	}

	return err
}

func AddBUfolderChild(payload AddMenuPayload, transaction *sql.Tx) error {
	var div_list []int
	var seq int

	listBURows, err := db.DB_DEV.Query(`select divoid, seq from bu_list order by seq asc`)
	if err != nil {
		return err
	}

	for listBURows.Next() {
		var divoid int
		if err := listBURows.Scan(&divoid, &seq); err != nil {
			return err
		}
		div_list = append(div_list, divoid)
	}

	// check if BU already exist in bu_list
	if slices.Contains(div_list, payload.Divoid) {
		// Do nothing
	} else {
		_, err := transaction.Exec(`
				insert into bu_list (
				  divoid 
				, seq
				, divname
				) values (
				  @divoid
				, @seq
				, @divname
				)
			`, sql.Named("divoid", payload.Divoid),
			sql.Named("seq", seq+1),
			sql.Named("divname", payload.Name))

		if err != nil {
			return err
		}
	}

	_, err = transaction.Exec(`
			insert into folder_bu (
				divoid
			  , folderoid
			) values (
			    @divoid
			  , @folderoid
			)
		`, sql.Named("divoid", payload.Divoid),
		sql.Named("folderoid", payload.Folderoid))
	return err
}

func AddBUdeptfolderChild(payload AddMenuPayload, transaction *sql.Tx) error {
	var div_list []int
	var seq int

	listBURows, err := db.DB_DEV.Query(`select divoid, seq from bu_list order by seq asc`)
	if err != nil {
		return err
	}

	for listBURows.Next() {
		var divoid int
		if err := listBURows.Scan(&divoid, &seq); err != nil {
			return err
		}
		div_list = append(div_list, divoid)
	}

	// check if BU already exist in bu_list
	if slices.Contains(div_list, payload.Divoid) {
		// Do nothing
	} else {
		_, err := transaction.Exec(`
				insert into bu_list (
				  divoid 
				, seq
				, divname
				) values (
				  @divoid
				, @seq
				, @divname
				)
			`, sql.Named("divoid", payload.Divoid),
			sql.Named("seq", seq+1),
			sql.Named("divname", payload.Name))

		if err != nil {
			return err
		}
	}

	// add to folder_dept
	_, err = transaction.Exec(`
			insert into folder_dept (
				  folderoid 
				, divoid
				, deptoid
			) values (
				  @folderoid
				, @divoid
				, @deptoid
			)
		`, sql.Named("folderoid", payload.Folderoid),
		sql.Named("divoid", payload.Divoid),
		sql.Named("deptoid", 0))
	return err
}

func AddBUdeptfolderLastChild(payload AddMenuPayload, transaction *sql.Tx) error {
	var lastDeptId int
	lastDeptIdrow := db.DB_DEV.QueryRow(`
			select top 1 deptoid from dept_list order by deptoid desc 
		`)
	err := lastDeptIdrow.Scan(&lastDeptId)
	if err != nil {
		return err
	}

	var folderoid = payload.Folderoid
	var divoid = payload.Divoid
	var deptoid = lastDeptId + 1
	var name = payload.Name
	var user = payload.User

	_, err = transaction.Exec(`
		insert into dept_list (
		    deptoid
		  , divoid
		  , name
		  , activeflag
		  , createuser
		  , createtime
		  , lastupdateuser
		  , lastupdatetime
		) values (
		 	@deptoid
		  , @divoid
		  , @name
		  , 'ACTIVE'
		  , @user
		  , getdate()
		  , @user
		  , getdate()
		)
		`,
		sql.Named("deptoid", deptoid),
		sql.Named("divoid", divoid),
		sql.Named("name", name),
		sql.Named("user", user),
	)
	if err != nil {
		return err
	}

	// check if deptoid = 0 is exist in the current directory
	var check_dept_id int
	checkRows := transaction.QueryRow(`
			select top 1 deptoid
			from folder_dept
			where folderoid = @folderoid
				and divoid = @divoid
			order by id asc
		`, sql.Named("folderoid", payload.Folderoid),
		sql.Named("divoid", payload.Divoid))

	err = checkRows.Scan(&check_dept_id)
	if err != nil {
		return err
	}

	args := []interface{}{
		sql.Named("folderoid", folderoid),
		sql.Named("divoid", divoid),
		sql.Named("deptoid", deptoid),
	}

	if check_dept_id == 0 {
		args = append(args, sql.Named("check_dept_id", check_dept_id))
		_, err = transaction.Exec(`
				update folder_dept
				set deptoid = @deptoid
				where folderoid = @folderoid
					and divoid = @divoid
					and deptoid = @check_dept_id
			`, args...)
		if err != nil {
			return err
		}
	} else {
		_, err = transaction.Exec(`
				insert into folder_dept (
				  folderoid 
				, divoid
				, deptoid
				) values (
				  @folderoid
				, @divoid
				, @deptoid
				)
			`, args...)
		if err != nil {
			return err
		}
	}

	return err
}
