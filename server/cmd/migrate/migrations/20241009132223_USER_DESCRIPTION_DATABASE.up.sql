CREATE TABLE IF NOT EXISTS USER_DETAILS(
  userid VARCHAR(100) UNIQUE NOT NULL CONSTRAINT fk_user REFERENCES users(id),
  avatar VARCHAR(100) NOT NULL, 
  cover_photo VARCHAR(100) NOT NULL,
  prefered_color VARCHAR(10) NOT NULL,
  dark_mode int not null
);

