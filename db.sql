create table folder_list (
    id int identity (1,1)
  , folder_id int
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