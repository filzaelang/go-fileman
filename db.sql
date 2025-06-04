create table folder_list (
    id int identity (1,1)
  , folder_id int null
  , dept_id int null
  , div_id int
  , level_id int
  , headfolder varchar(30) null
  , [name] varchar(50) null
  , createuser varchar(30) null
  , createtime datetime
  , lastupdateuser varchar(30) null
  , lastupdatetime datetime
  , [type] varchar(50) null
  , seq int null
  , folderhidebudept varchar(1) null
  , activeflag varchar(10) null

  , constraint folder_list_pk primary key (id)
)

create table bu_list (
    id int identity (1,1)
  , div_id int
  , seq int
  , divname varchar(50)

  , constraint bu_list_pk primary key (id)
)

create table folder_dept (
    id int identity (1,1)
  , div_id int
  , dept_id int
  , folder_id int

  , constraint folder_dept_pk primary key (id)
)

create table folder_bu (
    id int identity (1,1)
  , div_id int
  , folder_id int

  , constraint folder_bu_pk primary key (id)
)