package models

import (
	"database/sql"
	"file-manager/db"
	"fmt"
)

type MenuSidebar struct {
	Headfolder string         `json:"headfolder"`
	Name       string         `json:"name"`
	Uri        string         `json:"uri"`
	Type       string         `json:"type"`
	Seq        string         `json:"seq"`
	Children   []*MenuSidebar `json:"children, omitempty"`
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

		// Ambil semua anak folder dari headfolder
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
				continue
			}
			child.Uri = fmt.Sprintf("/folder/%d/0/0", folder_id)
			sm.Children = append(sm.Children, &child)
		}
		childRows.Close()

		sidebar = append(sidebar, sm)
	}

	return sidebar, nil
}

// select distinct headfolder, type, seq from itg_folder order by seq asc
