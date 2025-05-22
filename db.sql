create table menu_list (
    id int identity (1,1)
  , [name] varchar(50)
  , uri varchar(max)
  
  , constraint menu_list_pk primary key (id)
)