package models

import (
	"database/sql"
	"file-manager/db"
	"fmt"
	"slices"
)

var HEADFOLDER = "headfolder"
var SUBFOLDER = "subfolder"
var SUBFOLDER_CHILD = "subfolder_child"
var BUFOLDER = "bufolder"
var BUFOLDER_CHILD = "bufolder_child"
var BUDEPTFOLDER = "budeptfolder"
var BUDEPTFOLDER_CHILD = "budeptfolder_child"
var BUDEPTFOLDER_LAST_CHILD = "budeptfolder_last_child"

func GetSidebarMenu() ([]MenuSidebar, error) {
	rows, err := db.DB_DEV.Query("select distinct headfolder, type, seq from folder_list order by seq asc")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sidebar []MenuSidebar
	for rows.Next() {
		var sm MenuSidebar
		if err := rows.Scan(&sm.Headfolder, &sm.Type, &sm.Seq); err != nil {
			return nil, err
		}

		// Untuk type "headfolder"
		if sm.Type == HEADFOLDER {
			folderoid, err := GetHeadfolder(sm.Headfolder)
			if err != nil {
				return nil, err
			}
			sm.Uri = fmt.Sprintf("/folder/%d/0/0", folderoid)
			sm.Name = sm.Headfolder
			sm.Folderoid = folderoid
			sm.Divoid = 0
			sm.Deptoid = 0
		}

		// Ambil semua children dari type "subfolder"
		if sm.Type == SUBFOLDER {
			children, err := GetSubFolder(sm.Headfolder)
			if err != nil {
				return nil, err
			}
			sm.Name = sm.Headfolder
			sm.Folderoid = 0
			sm.Divoid = 0
			sm.Deptoid = 0
			sm.Children = children
		}

		// Ambil semua children dari type "bufolder"
		if sm.Type == BUFOLDER {
			children, folderoid, err := GetBuFolder(sm.Headfolder)
			if err != nil {
				return nil, err
			}
			sm.Name = sm.Headfolder
			sm.Folderoid = folderoid
			sm.Divoid = 0
			sm.Deptoid = 0
			sm.Children = children
		}

		// Ambil semua children dari type "budeptfolder"
		if sm.Type == BUDEPTFOLDER {
			children, folderoid, err := GetBuDepthFolder(sm.Headfolder)
			if err != nil {
				return nil, err
			}
			sm.Name = sm.Headfolder
			sm.Folderoid = folderoid
			sm.Divoid = 0
			sm.Deptoid = 0
			sm.Children = children
		}

		sidebar = append(sidebar, sm)
	}

	return sidebar, nil
}

func GetHeadfolder(Headfolder string) (int, error) {
	var folderoid int
	headfolderRows := db.DB_DEV.QueryRow(`
				select top 1 folderoid 
				from folder_list 
				where headfolder = @headfolder
					and name = @headfolder
			`, sql.Named("headfolder", Headfolder))
	err := headfolderRows.Scan(&folderoid)
	if err != nil {
		return 0, err
	}
	return folderoid, nil
}

func GetSubFolder(Headfolder string) ([]*MenuSidebar, error) {
	childRows, err := db.DB_DEV.Query(`
			select name, folderoid
			from folder_list
			where headfolder = @headfolder
				and headfolder != name`,
		sql.Named("headfolder", Headfolder),
	)
	if err != nil {
		return nil, err
	}

	var children []*MenuSidebar

	for childRows.Next() {
		var child MenuSidebar
		var folderoid int
		if err := childRows.Scan(&child.Name, &folderoid); err != nil {
			return nil, err // continue
		}
		child.Uri = fmt.Sprintf("/folder/%d/0/0", folderoid)
		child.Folderoid = folderoid
		child.Divoid = 0
		child.Deptoid = 0
		child.Type = SUBFOLDER_CHILD
		children = append(children, &child)
	}
	childRows.Close()

	return children, nil
}

func GetBuFolder(Headfolder string) ([]*MenuSidebar, int, error) {
	var children []*MenuSidebar

	var divIdList []int
	divIdListRows, err := db.DB_DEV.Query(`
		select distinct divoid, seq 
		from bu_list 
		where divoid in (
			select distinct divoid 
			from folder_bu 
			where folderoid in (
				select top 1 folderoid 
				from folder_list 
				where headfolder = @headfolder
			)
		) 
		order by seq asc
	`, sql.Named("headfolder", Headfolder))
	if err != nil {
		return nil, 0, err
	}
	for divIdListRows.Next() {
		var divoid int
		var seq int
		if err := divIdListRows.Scan(&divoid, &seq); err != nil {
			return nil, 0, err
		}
		divIdList = append(divIdList, divoid)
	}
	divIdListRows.Close()

	var buChildren []BuFolderChild
	for _, divoid := range divIdList {
		buChildRows, err := db.DB_MIS.Query(`
					select divoid, divname
					from QL_mstdivision
					where activeFlag = 'ACTIVE'
						and divoid = @divoid
					order by divoid asc
				`, sql.Named("divoid", divoid))
		if err != nil {
			return nil, 0, err
		}

		for buChildRows.Next() {
			var bc BuFolderChild
			if err := buChildRows.Scan(&bc.Divoid, &bc.Name); err != nil {
				return nil, 0, err
			}
			buChildren = append(buChildren, bc)
		}
		buChildRows.Close()
	}

	folderRows, err := db.DB_DEV.Query(`
		select folderoid from folder_list where headfolder = @headfolder
	`, sql.Named("headfolder", Headfolder))
	if err != nil {
		return nil, 0, err
	}
	defer folderRows.Close()

	var folderIDs []int
	var folderoid int
	for folderRows.Next() {
		var id int
		if err := folderRows.Scan(&id); err != nil {
			return nil, 0, err
		}
		folderoid = id
		folderIDs = append(folderIDs, id)
	}

	for _, bu := range buChildren {
		for _, folderID := range folderIDs {
			child := &MenuSidebar{
				Name:      bu.Name,
				Uri:       fmt.Sprintf("/folder/%d/%d/0", folderID, bu.Divoid),
				Folderoid: folderID,
				Divoid:    bu.Divoid,
				Deptoid:   0,
				Type:      BUDEPTFOLDER_CHILD,
			}
			children = append(children, child)
		}
	}

	return children, folderoid, nil
}

func GetBuDepthFolder(Headfolder string) ([]*MenuSidebar, int, error) {
	var children []*MenuSidebar

	var folderoid int
	rowsFolderId := db.DB_DEV.QueryRow(`
		select top 1 folderoid from folder_list where headfolder = @headfolder
	`, sql.Named("headfolder", Headfolder))
	err := rowsFolderId.Scan(&folderoid)
	if err != nil {
		return nil, 0, err
	}

	var divIdList []int
	divIdListRows, err := db.DB_DEV.Query(`
		select distinct divoid, seq
		from bu_list
		where divoid in (
			select distinct divoid
			from folder_dept
			where folderoid = @folderoid
		)
		order by seq asc
	`, sql.Named("folderoid", folderoid))
	if err != nil {
		return nil, 0, err
	}
	for divIdListRows.Next() {
		var divoid int
		var seq int
		if err := divIdListRows.Scan(&divoid, &seq); err != nil {
			return nil, 0, err
		}
		divIdList = append(divIdList, divoid)
	}
	divIdListRows.Close()

	for _, divoid := range divIdList {
		buRowsTwo, err := db.DB_MIS.Query(`
					select divoid, divname
					from QL_mstdivision
					where activeFlag = 'ACTIVE'
						and divoid = @divoid
					order by divoid asc
				`, sql.Named("divoid", divoid))
		if err != nil {
			return nil, 0, err
		}

		for buRowsTwo.Next() {
			var child MenuSidebar
			var divid int
			if err := buRowsTwo.Scan(&divid, &child.Name); err != nil {
				return nil, 0, err
			}

			child.Folderoid = folderoid
			child.Divoid = divoid
			child.Deptoid = 0
			child.Type = BUDEPTFOLDER_CHILD

			// Sini untuk menambah menu lastchild budept
			var deptIdList []int
			buRowsThree, err := db.DB_DEV.Query(`
				select distinct deptoid
				from folder_dept
				where divoid = @divoid
					and folderoid = @folderoid
			`, sql.Named("divoid", divoid),
				sql.Named("folderoid", folderoid))
			if err != nil {
				return nil, 0, err
			}
			for buRowsThree.Next() {
				var dept int
				if err := buRowsThree.Scan(&dept); err != nil {
					return nil, 0, err
				}
				deptIdList = append(deptIdList, dept)
			}
			buRowsThree.Close()

			if len(deptIdList) > 0 {
				buRowsFour, err := db.DB_DEV.Query(`
							select deptoid, name
							from dept_list
							where divoid = @divoid
								and activeflag = 'ACTIVE'
								and deptoid in (
									select distinct deptoid
									from folder_dept
									where divoid = @divoid
										and folderoid in (
											select top 1 folderoid
											from folder_list
											where headfolder = @headfolder
										)
								)
							order by name asc
						`, sql.Named("divoid", divoid),
					sql.Named("headfolder", Headfolder))
				if err != nil {
					return nil, 0, err
				}

				for buRowsFour.Next() {
					var lastChild MenuSidebar
					var dept_id int
					if err := buRowsFour.Scan(&dept_id, &lastChild.Name); err != nil {
						return nil, 0, err
					}
					lastChild.Uri = fmt.Sprintf("/folder/%d/%d/%d", folderoid, divoid, dept_id)
					lastChild.Folderoid = folderoid
					lastChild.Divoid = divid
					lastChild.Deptoid = dept_id
					lastChild.Type = BUDEPTFOLDER_LAST_CHILD
					child.Children = append(child.Children, &lastChild)
				}
				buRowsFour.Close()
			}
			children = append(children, &child)
		}
		buRowsTwo.Close()
	}
	return children, folderoid, nil
}

func GetOneMenu(payload UpdateMenuPayload) (string, error) {

	// if menu is inside bu_deptlist
	if payload.Type == BUDEPTFOLDER_LAST_CHILD {
		row := db.DB_DEV.QueryRow(`
		select top 1 name 
		from dept_list 
		where divoid = @divoid
			and deptoid = @deptoid
		`,
			sql.Named("divoid", payload.Divoid),
			sql.Named("deptoid", payload.Deptoid),
		)

		var name string
		err := row.Scan(&name)

		if err != nil {
			return "", err
		}

		return name, nil
	}

	return "", nil
}

func GetBUList(payload FolderID) ([]BUList, error) {
	var listBU []BUList

	// Take folder which exist in this folder
	var divIdList []int
	divIdListRows, err := db.DB_DEV.Query(`
		select distinct divoid, seq
		from bu_list
		where divoid in (
			select distinct divoid
			from folder_dept
			where folderoid = @folderoid
		)
		order by seq asc
	`, sql.Named("folderoid", payload.Folderoid))
	if err != nil {
		return nil, err
	}
	for divIdListRows.Next() {
		var divoid int
		var seq int
		if err := divIdListRows.Scan(&divoid, &seq); err != nil {
			return nil, err
		}
		divIdList = append(divIdList, divoid)
	}
	divIdListRows.Close()
	// Take folder which exist in this folder

	// Fetch only where divoid doesn exist
	placeholders := ""
	args := []interface{}{}
	for i, id := range divIdList {
		if i > 0 {
			placeholders += ", "
		}
		placeholders += fmt.Sprintf("@id%d", i)
		args = append(args, sql.Named(fmt.Sprintf("id%d", i), id))
	}

	query := `
		select divoid, divname
		from QL_mstdivision
		where activeFlag = 'ACTIVE'`

	// Filter not in if divoid exist
	if len(divIdList) > 0 {
		query += " and divoid not in (" + placeholders + ")"
	}

	query += " order by divoid asc"

	listBURows, err := db.DB_MIS.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer listBURows.Close()

	for listBURows.Next() {
		var bu BUList
		if err := listBURows.Scan(&bu.Divoid, &bu.Divname); err != nil {
			return nil, err
		}
		listBU = append(listBU, bu)
	}
	// Fetch only where divoid doesn exist

	return listBU, nil
}

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

	// if folder inside subfolder
	if payload.Type == SUBFOLDER {
		// Butuh folderoid (folder id terakhir + 1)
		// Headfolder = didapat dari payload
		// Name = didapat dari payload
		// divzip ('Sementara isi CORP')
		// c
	}

	//if folder inside bufolder
	if payload.Type == BUFOLDER {
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
		if err != nil {
			return err
		}

		return transaction.Commit()
	}

	// if folder inside budeptfolder
	if payload.Type == BUDEPTFOLDER {
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
		if err != nil {
			return err
		}

		return transaction.Commit()
	}

	// if it is folder inside budeptfolder_child
	if payload.Type == BUDEPTFOLDER_CHILD {

		var lastDeptId int
		lastDeptIdrow := db.DB_DEV.QueryRow(`
			select top 1 deptoid from dept_list order by deptoid desc 
		`)
		err = lastDeptIdrow.Scan(&lastDeptId)
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

		return transaction.Commit()
	}

	return nil
}

func UpdateMenu(payload UpdateMenuPayload) error {
	// if it is folder inside budept
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

func DeleteMenu(payload DeleteMenuPayload) error {

	// if it is folder inside budept
	if payload.Type == BUDEPTFOLDER_LAST_CHILD {
		transaction, err := db.DB_DEV.Begin()
		if err != nil {
			return err
		}
		defer func() {
			if err != nil {
				transaction.Rollback()
			}
		}()

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
