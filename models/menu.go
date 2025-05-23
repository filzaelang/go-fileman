// models/menu.go
package models

import (
	"database/sql"
	"file-manager/db"
)

type MenuItem struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Uri      *string     `json:"uri,omitempty"`
	ParentID *int        `json:"parent_id"`
	Children []*MenuItem `json:"children,omitempty"`
}

type MenuInput struct {
	Name     string `json:"name"`
	Uri      string `json:"uri"`
	ParentID *int   `json:"parent_id"`
}

func GetFlatMenus() ([]MenuItem, error) {
	rows, err := db.DB.Query("select id, name, uri, parent_id from menu_list order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flat []MenuItem
	for rows.Next() {
		var m MenuItem
		if err := rows.Scan(&m.ID, &m.Name, &m.Uri, &m.ParentID); err != nil {
			return nil, err
		}
		flat = append(flat, m)
	}

	return flat, nil
}

func BuildMenuTree(flat []MenuItem) []*MenuItem {
	idMap := make(map[int]*MenuItem)
	visited := make(map[int]bool)

	// Index semua menu berdasarkan ID
	for i := range flat {
		idMap[flat[i].ID] = &flat[i]
	}

	var roots []*MenuItem

	for _, item := range flat {
		// Prevent circular reference
		if func() bool {
			var _ map[int]bool = visited
			return detectCircular(item.ID, idMap)
		}() {
			continue
		}

		if item.ParentID != nil {
			parent := idMap[*item.ParentID]
			if parent != nil {
				parent.Children = append(parent.Children, idMap[item.ID])
			}
		} else {
			roots = append(roots, idMap[item.ID])
		}
	}

	return roots
}

// Cegah loop antar parent_id yang menyebabkan infinite recursion
func detectCircular(id int, idMap map[int]*MenuItem) bool {
	current := id
	seen := map[int]bool{}

	for {
		if seen[current] {
			// Detected cycle!
			return true
		}
		seen[current] = true

		node := idMap[current]
		if node == nil || node.ParentID == nil {
			break
		}

		current = *node.ParentID
	}

	return false
}

func InsertMenu(menu MenuItem) error {
	_, err := db.DB.Exec(`
		insert into menu_list (name, uri, parent_id)
		values (@name, @uri, @parent_id)
		`,
		sql.Named("name", menu.Name),
		sql.Named("uri", menu.Uri),
		sql.Named("parent_id", menu.ParentID),
	)
	return err
}

func GetOneMenu(id int) (*MenuItem, error) {
	row := db.DB.QueryRow(`
		select id, name, uri, parent_id from menu_list where id = @id 
		`,
		sql.Named("id", id),
	)

	var menu MenuItem
	err := row.Scan(&menu.ID, &menu.Name, &menu.Uri, &menu.ParentID)

	if err != nil {
		return nil, err
	}

	return &menu, nil
}

func UpdateMenu(menu MenuItem) error {
	_, err := db.DB.Exec(`
		update menu_list SET 
			  name = @name
			, uri = @uri
			, parent_id = @parent_id 
		where id = @id
		`,
		sql.Named("name", menu.Name),
		sql.Named("uri", menu.Uri),
		sql.Named("parent_id", menu.ParentID),
		sql.Named("id", menu.ID),
	)
	return err
}

func DeleteMenu(id int) error {
	_, err := db.DB.Exec(`
		with RecursiveMenu AS (
			select id
			from menu_list
			where id = @id
			union all
			select m.id
			from menu_list m
			inner join RecursiveMenu rm ON m.parent_id = rm.id
		)
		delete from menu_list
		where id in (select id from RecursiveMenu);
		`,
		sql.Named("id", id),
	)
	return err
}
