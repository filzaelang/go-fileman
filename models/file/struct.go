package model_file

import "time"

type FileItem struct {
	Fileoid        int       `json:"fileoid"`
	Divoid         int       `json:"divoid"`
	Deptoid        int       `json:"deptoid"`
	Leveloid       int       `json:"leveloid"`
	Folderoid      int       `json:"folderoid"`
	Filename       string    `json:"filename"`
	Fileurl        string    `json:"fileurl"`
	Createuser     string    `json:"createuser"`
	Createtime     time.Time `json:"createtime"`
	Lastupdatetime time.Time `json:"lastupdatetime"`
	Filenumber     string    `json:"filenumber"`
	Filerevdate    time.Time `json:"filerevdate"`
	Fileoldnumber  string    `json:"fileoldnumber"`
	Filevisible    bool      `json:"filevisible"`
}

type ItgFile struct {
	FileURL string
}

type Log struct {
	Fileoid int    `json:"fileoid"`
	User    string `json:"user"`
	Action  string `json:"action"`
	Deptoid int    `json:"deptoid"`
	Counter int    `json:"counter"`
}
