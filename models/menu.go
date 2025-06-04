package models

import (
	"database/sql"
	"file-manager/db"
	"fmt"
)

type MenuSidebar struct {
	Headfolder string         `json:"headfolder"`
	Name       string         `json:"name"`
	Div_id     int            `json:"div_id"`
	Uri        string         `json:"uri"`
	Type       string         `json:"type"`
	Seq        string         `json:"seq"`
	Children   []*MenuSidebar `json:"children"`
}

type BuFolderChild struct {
	Div_id int    `json:"div_id"`
	Name   string `json:"name"`
}

func GetSidebarMenu() ([]MenuSidebar, error) {
	rows, err := db.DB_DEV.Query("select distinct headfolder, type, seq from folder_list where headfolder is not null order by seq asc")
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

		// Ambil semua children dari type "subfolder"
		if sm.Type == "subfolder" {
			childRows, err := db.DB_DEV.Query(`
			select name, folder_id
			from folder_list
			where headfolder = @headfolder
				and headfolder != name`,
				sql.Named("headfolder", sm.Headfolder),
			)
			if err != nil {
				return nil, err
			}

			for childRows.Next() {
				var child MenuSidebar
				var folder_id int
				if err := childRows.Scan(&child.Name, &folder_id); err != nil {
					return nil, err // continue
				}
				child.Uri = fmt.Sprintf("/folder/%d/0/0", folder_id)
				sm.Children = append(sm.Children, &child)
			}
			childRows.Close()
		}

		// Ambil semua children dari type "bufolder"
		if sm.Type == "bufolder" {
			buRows, err := db.DB_DEV.Query(`
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
			`, sql.Named("headfolder", sm.Headfolder))
			if err != nil {
				return nil, err
			}

			var buDivIdList []int

			for buRows.Next() {
				var div_id int
				var seq int
				if err := buRows.Scan(&div_id, &seq); err != nil {
					return nil, err
				}
				buDivIdList = append(buDivIdList, div_id)
			}
			buRows.Close()

			var BuChild []BuFolderChild
			for _, value := range buDivIdList {
				buRowsTwo, err := db.DB_MIS.Query(`
					select divoid, divname
					from QL_mstdivision
					where activeFlag = 'ACTIVE'
						and divoid = @divoid
					order by divoid asc
				`, sql.Named("divoid", value))
				if err != nil {
					return nil, err
				}

				for buRowsTwo.Next() {
					var bc BuFolderChild
					if err := buRowsTwo.Scan(&bc.Div_id, &bc.Name); err != nil {
						return nil, err // continue
					}
					BuChild = append(BuChild, bc)
				}
				buRowsTwo.Close()
			}

			for _, value := range BuChild {
				buRowsThree, err := db.DB_DEV.Query(`
					select folder_id
					from folder_list
					where headfolder = @headfolder
				`, sql.Named("headfolder", sm.Headfolder))
				if err != nil {
					return nil, err
				}

				for buRowsThree.Next() {
					var child MenuSidebar
					var folder_id int
					if err := buRowsThree.Scan(&folder_id); err != nil {
						return nil, err
					}
					child.Uri = fmt.Sprintf("/folder/%d/%d/0", folder_id, value.Div_id)
					child.Name = value.Name
					sm.Children = append(sm.Children, &child)
				}
				buRowsThree.Close()
			}
		}

		// Ambil semua children dari type "budeptfolder"
		if sm.Type == "budeptfolder" {
			var folder_id int
			rowsFolderId := db.DB_DEV.QueryRow(`
					select top 1 folder_id from folder_list where headfolder = @headfolder
				`, sql.Named("headfolder", sm.Headfolder))
			err = rowsFolderId.Scan(&folder_id)
			if err != nil {
				return nil, err
			}

			var buDivIdList []int
			buRows, err := db.DB_DEV.Query(`
				select distinct div_id, seq
				from bu_list
				where div_id in (
					select distinct div_id
					from folder_dept
					where folder_id in (
						select top 1 folder_id
						from folder_list where headfolder = @headfolder
					)
				)
				order by seq asc
			`, sql.Named("headfolder", sm.Headfolder))
			if err != nil {
				return nil, err
			}
			for buRows.Next() {
				var div_id int
				var seq int
				if err := buRows.Scan(&div_id, &seq); err != nil {
					return nil, err
				}
				buDivIdList = append(buDivIdList, div_id)
			}
			buRows.Close()

			for _, value := range buDivIdList {
				buRowsTwo, err := db.DB_MIS.Query(`
					select divoid, divname
					from QL_mstdivision
					where activeFlag = 'ACTIVE'
						and divoid = @divoid
					order by divoid asc
				`, sql.Named("divoid", value))
				if err != nil {
					return nil, err
				}

				for buRowsTwo.Next() {
					var child MenuSidebar
					var div_id int
					if err := buRowsTwo.Scan(&div_id, &child.Name); err != nil {
						return nil, err
					}

					var deptIdList []int
					buRowsThree, err := db.DB_DEV.Query(`
						select distinct dept_id
						from folder_dept
						where div_id = @div_id
							and folder_id in (
								select top 1 folder_id
								from folder_list
								where headfolder = @headfolder
							)
					`, sql.Named("div_id", div_id),
						sql.Named("headfolder", sm.Headfolder))
					if err != nil {
						return nil, err
					}
					for buRowsThree.Next() {
						var dept int
						if err := buRowsThree.Scan(&dept); err != nil {
							return nil, err
						}
						deptIdList = append(deptIdList, dept)
					}
					buRowsThree.Close()

					if len(deptIdList) > 0 {
						buRowsFour, err := db.DB_DEV.Query(`
							select dept_id, name
							from folder_list
							where div_id = @div_id
								and activeflag = 'ACTIVE'
								and dept_id in (
									select distinct dept_id
									from folder_dept
									where div_id = @div_id
										and folder_id in (
											select top 1 folder_id
											from folder_list
											where headfolder = @headfolder
										)
								)
							order by name asc
						`, sql.Named("div_id", div_id),
							sql.Named("headfolder", sm.Headfolder))
						if err != nil {
							return nil, err
						}

						for buRowsFour.Next() {
							var lastChild MenuSidebar
							var dept_id int
							if err := buRowsFour.Scan(&dept_id, &lastChild.Name); err != nil {
								return nil, err
							}
							lastChild.Uri = fmt.Sprintf("/folder/%d/%d/%d", folder_id, div_id, dept_id)
							child.Children = append(child.Children, &lastChild)
						}
						buRowsFour.Close()
					}
					sm.Children = append(sm.Children, &child)
				}
				buRowsTwo.Close()
			}
		}

		sidebar = append(sidebar, sm)
	}

	return sidebar, nil
}

func SubFolder(sm MenuSidebar) error {
	childRows, err := db.DB_DEV.Query(`
			select name, folder_id
			from folder_list
			where headfolder = @headfolder
				and headfolder != name`,
		sql.Named("headfolder", sm.Headfolder),
	)
	if err != nil {
		return err
	}

	for childRows.Next() {
		var child MenuSidebar
		var folder_id int
		if err := childRows.Scan(&child.Name, &folder_id); err != nil {
			return err // continue
		}
		child.Uri = fmt.Sprintf("/folder/%d/0/0", folder_id)
		sm.Children = append(sm.Children, &child)
	}
	childRows.Close()

	return nil
}
