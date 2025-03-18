ALTER TABLE users
ALTER COLUMN username TYPE CITEXT,
ADD CONSTRAINT username_length_check CHECK ( length(username) >= 5 and length(username) <= 20);
