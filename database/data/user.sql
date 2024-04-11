CREATE TABLE user  IF  NOT EXISTS (
    id  int  primary key,                 --
    username varchar(255) NOT NULL ,      --
    password varchar(255) NOT NULL ,      --
    user_role  int  NOT NULL ,

)