package model_file

import (
	"database/sql"
	"file-manager/db"
)

func GetFile(folderoid, divoid, deptoid int) ([]FileItem, error) {

	var roleId int = 1 // di dapat dari session login, sementara dummy

	query := `
		select 
			fileoid
		  , divoid
		  , deptoid
		  , leveloid
		  , folderoid
		  , filename
		  , fileurl
		  , createuser
		  , createtime
		  , lastupdatetime
		  , filenumber
		  , filerevdate
		  , fileoldnumber
		  , filevisible  
		from file_list
		where folderoid = @folderoid`

	args := []interface{}{}
	args = append(args, sql.Named("folderoid", folderoid))

	if divoid != 0 {
		query += " and divoid = @divoid"
		args = append(args, sql.Named("divoid", divoid))
	}

	if deptoid != 0 {
		query += " and deptoid = @deptoid"
		args = append(args, sql.Named("deptoid", deptoid))
	}

	if roleId == 2 {
		query += "and filevisible = 'True'"
	}

	query += " order by"

	if folderoid == 0 {
		query += `case
			when fileurl LIKE '%.htm' then 4
			when fileurl LIKE '%.pps' then 3
			when fileurl LIKE '%.mp4' then 2
			when fileurl LIKE '%.pdf' then 1
		end desc`
	}

	query += " lastupdatetime desc"

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []FileItem

	for rows.Next() {
		var item FileItem
		if err := rows.Scan(
			&item.Fileoid,
			&item.Divoid,
			&item.Deptoid,
			&item.Leveloid,
			&item.Folderoid,
			&item.Filename,
			&item.Fileurl,
			&item.Createuser,
			&item.Createtime,
			&item.Lastupdatetime,
			&item.Filenumber,
			&item.Filerevdate,
			&item.Fileoldnumber,
			&item.Filevisible,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
