CREATE TABLE USER_DETAILS(
userid VARCHAR(100) UNIQUE NOT NULL,
  avatar VARCHAR(100) NOT NULL, 
  coverphoto VARCHAR(100) NOT NULL,
  prefered_color VARCHAR(10) NOT NULL,
created_at DATETIME DEFAULT current_timestamp,
  updated_at DATETIME DEFAULT current_timestamp ON UPDATE current_timestamp,
  dark_mode int not null,

  FOREIGN KEY(userid) REFERENCES user(id)
)
