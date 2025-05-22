// models/menu.go
package models

import (
	"file-manager/db"
)

type MenuItem struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Uri      *string     `json:"uri,omitempty"`
	ParentID *int        `json:"parent_id,omitempty"`
	Children []*MenuItem `json:"children,omitempty"`
}

func GetFlatMenus() ([]MenuItem, error) {
	rows, err := db.DB.Query("SELECT id, name, uri, parent_id FROM menu_list ORDER BY id")
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
