package models_menu

import (
	"database/sql"
	"file-manager/db"
	"fmt"
)

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
	// Fetch only where divoid doesn't exist

	return listBU, nil
}
