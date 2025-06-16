package model_file

import "database/sql"

func LogDownload(log Log, transaction *sql.Tx) error {
	var deptoid *int
	if log.Deptoid == 0 {
		deptoid = nil
	} else {
		deptoid = &log.Deptoid
	}

	var counter *int
	if log.Counter == 0 {
		counter = nil
	} else {
		counter = &log.Counter
	}

	_, err := transaction.Exec(`
		insert into log (
		    fileoid
		  , [user]
		  , [action]
		  , [datetime]
		  , deptoid
		  , [counter]
		) values (
		    @fileoid
		  , @user
		  , @action
		  , getdate()
		  , @deptoid
		  , @counter
		) 
		`, sql.Named("fileoid", log.Fileoid),
		sql.Named("user", log.User),
		sql.Named("action", log.Action),
		sql.Named("deptoid", deptoid),
		sql.Named("counter", counter))

	return err
}
