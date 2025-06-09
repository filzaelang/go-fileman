create table folder_list (
    id int identity (1,1)
  , folderoid int null
  , divoid int
  , deptoid int null
  , leveloid int
  , headfolder varchar(30) null
  , [name] varchar(50) null
  , divzip varchar(20)
  , createuser varchar(30) null
  , createtime datetime
  , lastupdateuser varchar(30) null
  , lastupdatetime datetime
  , [type] varchar(50) null
  , seq int 
  , folderhidebudept varchar(1) null

  , constraint folder_list_pk primary key (id)
)

create table dept_list (
    id int identity (1,1)
  , deptoid int
  , divoid int
  , [name] varchar(50)
  , activeflag varchar(20)
  , createuser varchar(20)
  , createtime datetime
  , lastupdateuser varchar(20)
  , lastupdatetime datetime
  , deptdistcode varchar(20) null

  , constraint dept_list_pk primary key (id)
)

create table folder_dept (
    id int identity (1,1)
  , folderoid int
  , divoid int
  , deptoid int

  , constraint folder_dept_pk primary key (id)
)

create table bu_list (
    id int identity (1,1)
  , divoid int
  , seq int
  , divname varchar(50)

  , constraint bu_list_pk primary key (id)
)



create table folder_bu (
    id int identity (1,1)
  , divoid int
  , folderoid int

  , constraint folder_bu_pk primary key (id)
)