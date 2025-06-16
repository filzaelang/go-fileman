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
  , constraint folder_list_uq unique (folderoid)
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
  , constraint dept_list_uq unique (deptoid)
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

create table file_list (
    id int identity (1,1)
  , fileoid int
  , divoid int
  , deptoid int
  , leveloid int
  , folderoid int
  , [filename] varchar(max)
  , fileurl varchar(max)
  , createuser varchar(30)
  , createtime datetime
  , lastupdateuser varchar(30)
  , lastupdatetime datetime
  , filenumber varchar(max)
  , filerevnumber int
  , filerevdate datetime
  , fileoldnumber varchar(max) null
  , filevisible varchar(5)

  , constraint file_list_pk primary key (id)
  , constraint file_list_uq unique (fileoid)
)

-- "set identity_insert log on"
create table log (
    logoid int identity (1,1)
  , fileoid int
  , [user] varchar(20)
  , [action] varchar(20)
  , [datetime] datetime
  , deptoid int null
  , counter int null

  , constraint log_pk primary key (id)
)
-- "set identity_insert log off"

create table role_list (
    profoid varchar(20)
  , proftype varchar(20)

  , constraint role_list_pk primary key(profoid)
)