package models_menu

type AddMenuPayload struct {
	Folderoid  int    `json:"folderoid"`
	Divoid     int    `json:"divoid"`
	Name       string `json:"name"`
	User       string `json:"user"`
	Type       string `json:"type"`
	Headfolder string `json:"headfolder"`
	IsBase     bool   `json:"isbase"`
}

type DeleteMenuPayload struct {
	Folderoid  int    `json:"folderoid"`
	Divoid     int    `json:"divoid"`
	Deptoid    int    `json:"deptoid"`
	Type       string `json:"type"`
	Headfolder string `json:"headfolder"`
}

type UpdateMenuPayload struct {
	Folderoid  *int   `json:"folderoid"`
	Divoid     *int   `json:"divoid"`
	Deptoid    *int   `json:"deptoid"`
	Headfolder string `json:"headfolder"`
	Name       string `json:"name"`
	User       string `json:"user"`
	Type       string `json:"type"`
}

type MenuSidebar struct {
	Headfolder string         `json:"headfolder"`
	Name       string         `json:"name"`
	Folderoid  int            `json:"folderoid"`
	Divoid     int            `json:"divoid"`
	Deptoid    int            `json:"deptoid"`
	Uri        string         `json:"uri"`
	Type       string         `json:"type"`
	Seq        string         `json:"seq"`
	Children   []*MenuSidebar `json:"children"`
}

type BuFolderChild struct {
	Divoid int    `json:"divoid"`
	Name   string `json:"name"`
}

type BUList struct {
	Divoid  int    `json:"divoid"`
	Divname string `json:"divname"`
	Seq     *int   `json:"seq"`
}

type FolderID struct {
	Folderoid int `json:"folderoid"`
}

type SubFolderList struct {
	Headfolder string `json:"headfolder"`
	Seq        int    `json:"seq"`
}
