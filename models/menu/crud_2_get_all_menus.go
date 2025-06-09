package models_menu

import (
	"database/sql"
	"file-manager/db"
	"fmt"
)

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
