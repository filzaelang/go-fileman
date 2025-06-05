package models

import (
	"database/sql"
	"file-manager/db"
	"fmt"
)

type AddMenuPayload struct {
	Folder_id int    `json:"folder_id"`
	Div_id    int    `json:"div_id"`
	Name      string `json:"name"`
	User      string `json:"user"`
	Type      string `json:"type"`
}

type DeleteMenuPayload struct {
	Folder_id int    `json:"folder_id"`
	Div_id    int    `json:"div_id"`
	Dept_id   int    `json:"dept_id"`
	Type      string `json:"type"`
}

type UpdateMenuPayload struct {
	Div_id  int    `json:"div_id"`
	Dept_id int    `json:"dept_id"`
	Name    string `json:"name"`
	User    string `json:"user"`
	Type    string `json:"type"`
}

type MenuSidebar struct {
	Headfolder string         `json:"headfolder"`
	Name       string         `json:"name"`
	Folder_id  int            `json:"folder_id"`
	Div_id     int            `json:"div_id"`
	Dept_id    int            `json:"dept_id"`
	Uri        string         `json:"uri"`
	Type       string         `json:"type"`
	Seq        string         `json:"seq"`
	Children   []*MenuSidebar `json:"children"`
}

type BuFolderChild struct {
	Div_id int    `json:"div_id"`
	Name   string `json:"name"`
}

type BUList struct {
	Div_id   int    `json:"div_id"`
	Div_name string `json:"div_name"`
}

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
			folder_id, err := GetHeadfolder(sm.Headfolder)
			if err != nil {
				return nil, err
			}
			sm.Uri = fmt.Sprintf("/folder/%d/0/0", folder_id)
			sm.Name = sm.Headfolder
			sm.Folder_id = folder_id
			sm.Div_id = 0
			sm.Dept_id = 0
		}

		// Ambil semua children dari type "subfolder"
		if sm.Type == SUBFOLDER {
			children, err := GetSubFolder(sm.Headfolder)
			if err != nil {
				return nil, err
			}
			sm.Name = sm.Headfolder
			sm.Folder_id = 0
			sm.Div_id = 0
			sm.Dept_id = 0
			sm.Children = children
		}

		// Ambil semua children dari type "bufolder"
		if sm.Type == BUFOLDER {
			children, err := GetBuFolder(sm.Headfolder)
			if err != nil {
				return nil, err
			}
			sm.Name = sm.Headfolder
			sm.Folder_id = 0
			sm.Div_id = 0
			sm.Dept_id = 0
			sm.Children = children
		}

		// Ambil semua children dari type "budeptfolder"
		if sm.Type == BUDEPTFOLDER {
			children, err := GetBuDepthFolder(sm.Headfolder)
			if err != nil {
				return nil, err
			}
			sm.Name = sm.Headfolder
			sm.Folder_id = 0
			sm.Div_id = 0
			sm.Dept_id = 0
			sm.Children = children
		}

		sidebar = append(sidebar, sm)
	}

	return sidebar, nil
}

func GetHeadfolder(Headfolder string) (int, error) {
	var folder_id int
	headfolderRows := db.DB_DEV.QueryRow(`
				select top 1 folder_id 
				from folder_list 
				where headfolder = @headfolder
					and name = @headfolder
			`, sql.Named("headfolder", Headfolder))
	err := headfolderRows.Scan(&folder_id)
	if err != nil {
		return 0, err
	}
	return folder_id, nil
}

func GetSubFolder(Headfolder string) ([]*MenuSidebar, error) {
	childRows, err := db.DB_DEV.Query(`
			select name, folder_id
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
		var folder_id int
		if err := childRows.Scan(&child.Name, &folder_id); err != nil {
			return nil, err // continue
		}
		child.Uri = fmt.Sprintf("/folder/%d/0/0", folder_id)
		child.Folder_id = folder_id
		child.Div_id = 0
		child.Dept_id = 0
		child.Type = SUBFOLDER_CHILD
		children = append(children, &child)
	}
	childRows.Close()

	return children, nil
}

func GetBuFolder(Headfolder string) ([]*MenuSidebar, error) {
	var children []*MenuSidebar

	var divIdList []int
	divIdListRows, err := db.DB_DEV.Query(`
		select distinct div_id, seq 
		from bu_list 
		where div_id in (
			select distinct div_id 
			from folder_bu 
			where folder_id in (
				select top 1 folder_id 
				from folder_list 
				where headfolder = @headfolder
			)
		) 
		order by seq asc
	`, sql.Named("headfolder", Headfolder))
	if err != nil {
		return nil, err
	}
	for divIdListRows.Next() {
		var div_id int
		var seq int
		if err := divIdListRows.Scan(&div_id, &seq); err != nil {
			return nil, err
		}
		divIdList = append(divIdList, div_id)
	}
	divIdListRows.Close()

	var buChildren []BuFolderChild
	for _, div_id := range divIdList {
		buChildRows, err := db.DB_MIS.Query(`
					select divoid, divname
					from QL_mstdivision
					where activeFlag = 'ACTIVE'
						and divoid = @divoid
					order by divoid asc
				`, sql.Named("divoid", div_id))
		if err != nil {
			return nil, err
		}

		for buChildRows.Next() {
			var bc BuFolderChild
			if err := buChildRows.Scan(&bc.Div_id, &bc.Name); err != nil {
				return nil, err
			}
			buChildren = append(buChildren, bc)
		}
		buChildRows.Close()
	}

	folderRows, err := db.DB_DEV.Query(`
		select folder_id from folder_list where headfolder = @headfolder
	`, sql.Named("headfolder", Headfolder))
	if err != nil {
		return nil, err
	}
	defer folderRows.Close()

	var folderIDs []int
	for folderRows.Next() {
		var id int
		if err := folderRows.Scan(&id); err != nil {
			return nil, err
		}
		folderIDs = append(folderIDs, id)
	}

	for _, bu := range buChildren {
		for _, folderID := range folderIDs {
			child := &MenuSidebar{
				Name:      bu.Name,
				Uri:       fmt.Sprintf("/folder/%d/%d/0", folderID, bu.Div_id),
				Folder_id: folderID,
				Div_id:    bu.Div_id,
				Dept_id:   0,
				Type:      BUDEPTFOLDER_CHILD,
			}
			children = append(children, child)
		}
	}

	return children, nil
}

func GetBuDepthFolder(Headfolder string) ([]*MenuSidebar, error) {
	var children []*MenuSidebar

	var folder_id int
	rowsFolderId := db.DB_DEV.QueryRow(`
		select top 1 folder_id from folder_list where headfolder = @headfolder
	`, sql.Named("headfolder", Headfolder))
	err := rowsFolderId.Scan(&folder_id)
	if err != nil {
		return nil, err
	}

	var divIdList []int
	divIdListRows, err := db.DB_DEV.Query(`
		select distinct div_id, seq
		from bu_list
		where div_id in (
			select distinct div_id
			from folder_dept
			where folder_id = @folder_id
		)
		order by seq asc
	`, sql.Named("folder_id", folder_id))
	if err != nil {
		return nil, err
	}
	for divIdListRows.Next() {
		var div_id int
		var seq int
		if err := divIdListRows.Scan(&div_id, &seq); err != nil {
			return nil, err
		}
		divIdList = append(divIdList, div_id)
	}
	divIdListRows.Close()

	for _, div_id := range divIdList {
		buRowsTwo, err := db.DB_MIS.Query(`
					select divoid, divname
					from QL_mstdivision
					where activeFlag = 'ACTIVE'
						and divoid = @divoid
					order by divoid asc
				`, sql.Named("divoid", div_id))
		if err != nil {
			return nil, err
		}

		for buRowsTwo.Next() {
			var child MenuSidebar
			var divid int
			if err := buRowsTwo.Scan(&divid, &child.Name); err != nil {
				return nil, err
			}

			child.Folder_id = folder_id
			child.Div_id = div_id
			child.Dept_id = 0
			child.Type = BUDEPTFOLDER_CHILD

			// Sini untuk menambah menu lastchild budept
			// var deptIdList []int
			// buRowsThree, err := db.DB_DEV.Query(`
			// 	select distinct dept_id
			// 	from folder_dept
			// 	where div_id = @div_id
			// 		and folder_id = @folder_id
			// `, sql.Named("div_id", div_id),
			// 	sql.Named("folder_id", folder_id))
			// if err != nil {
			// 	return nil, err
			// }
			// for buRowsThree.Next() {
			// 	var dept int
			// 	if err := buRowsThree.Scan(&dept); err != nil {
			// 		return nil, err
			// 	}
			// 	deptIdList = append(deptIdList, dept)
			// }
			// buRowsThree.Close()

			// if len(deptIdList) > 0 {
			// 	buRowsFour, err := db.DB_DEV.Query(`
			// 				select dept_id, name
			// 				from dept_list
			// 				where div_id = @div_id
			// 					and activeflag = 'ACTIVE'
			// 					and dept_id in (
			// 						select distinct dept_id
			// 						from folder_dept
			// 						where div_id = @div_id
			// 							and folder_id in (
			// 								select top 1 folder_id
			// 								from folder_list
			// 								where headfolder = @headfolder
			// 							)
			// 					)
			// 				order by name asc
			// 			`, sql.Named("div_id", div_id),
			// 		sql.Named("headfolder", Headfolder))
			// 	if err != nil {
			// 		return nil, err
			// 	}

			// 	for buRowsFour.Next() {
			// 		var lastChild MenuSidebar
			// 		var dept_id int
			// 		if err := buRowsFour.Scan(&dept_id, &lastChild.Name); err != nil {
			// 			return nil, err
			// 		}
			// 		lastChild.Uri = fmt.Sprintf("/folder/%d/%d/%d", folder_id, div_id, dept_id)
			// 		lastChild.Folder_id = folder_id
			// 		lastChild.Div_id = divid
			// 		lastChild.Dept_id = dept_id
			// 		lastChild.Type = BUDEPTFOLDER_LAST_CHILD
			// 		child.Children = append(child.Children, &lastChild)
			// 	}
			// 	buRowsFour.Close()
			// }
			children = append(children, &child)
		}
		buRowsTwo.Close()
	}
	return children, nil
}

func GetOneMenu(payload UpdateMenuPayload) (string, error) {

	// if menu is inside bu_deptlist
	if payload.Type == BUDEPTFOLDER_LAST_CHILD {
		row := db.DB_DEV.QueryRow(`
		select top 1 name 
		from dept_list 
		where div_id = @div_id
			and dept_id = @dept_id
		`,
			sql.Named("div_id", payload.Div_id),
			sql.Named("dept_id", payload.Dept_id),
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

func GetBUList() ([]BUList, error) {
	var listBU []BUList
	listBURows, err := db.DB_MIS.Query(`
		select divoid, divname
		from QL_mstdivision
		where activeFlag = 'ACTIVE'
		order by divoid asc 
	`)
	if err != nil {
		return nil, err
	}

	for listBURows.Next() {
		var bu BUList
		if err := listBURows.Scan(&bu.Div_id, &bu.Div_name); err != nil {
			return nil, err
		}
		listBU = append(listBU, bu)
	}

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

	// if it is folder inside budeptfolder
	if payload.Type == BUDEPTFOLDER {

	}
	// if it is folder inside budeptfolder_child
	if payload.Type == BUDEPTFOLDER_CHILD {

		var lastDeptId int
		lastDeptIdrow := db.DB_DEV.QueryRow(`
			select top 1 dept_id from dept_list order by dept_id desc 
		`)
		err = lastDeptIdrow.Scan(&lastDeptId)
		if err != nil {
			return err
		}

		var folder_id = payload.Folder_id
		var div_id = payload.Div_id
		var dept_id = lastDeptId + 1
		var name = payload.Name
		var user = payload.User

		_, err = transaction.Exec(`
		insert into dept_list (
		    dept_id
		  , div_id
		  , name
		  , activeflag
		  , createuser
		  , createtime
		  , lastupdateuser
		  , lastupdatetime
		) values (
		 	@dept_id
		  , @div_id
		  , @name
		  , 'ACTIVE'
		  , @user
		  , getdate()
		  , @user
		  , getdate()
		)
		`,
			sql.Named("dept_id", dept_id),
			sql.Named("div_id", div_id),
			sql.Named("name", name),
			sql.Named("user", user),
		)
		if err != nil {
			return err
		}

		_, err = transaction.Exec(`
		insert into folder_dept (
			folder_id 
		  , div_id
		  , dept_id
		) values (
		    @folder_id
		  , @div_id
		  , @dept_id
		)
	`, sql.Named("folder_id", folder_id),
			sql.Named("div_id", div_id),
			sql.Named("dept_id", dept_id))
		if err != nil {
			return err
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
		where div_id = @div_id
			and dept_id = @dept_id
	`, sql.Named("name", payload.Name),
			sql.Named("user", payload.User),
			sql.Named("div_id", payload.Div_id),
			sql.Named("dept_id", payload.Dept_id))

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
		where folder_id = @folder_id
			and div_id = @div_id
			and dept_id = @dept_id`,
			sql.Named("folder_id", payload.Folder_id),
			sql.Named("div_id", payload.Div_id),
			sql.Named("dept_id", payload.Dept_id),
		)
		if err != nil {
			return err
		}

		_, err = transaction.Exec(`
		delete from dept_list
		where div_id = @div_id
			and dept_id = @dept_id
	`, sql.Named("div_id", payload.Div_id),
			sql.Named("dept_id", payload.Dept_id))
		if err != nil {
			return err
		}

		return transaction.Commit()
	}

	return nil
}
