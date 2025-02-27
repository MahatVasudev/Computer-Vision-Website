ALTER TABLE posts
ADD COLUMN IF NOT EXISTS user_id varchar(255),
ADD CONSTRAINT fk_post_user FOREIGN KEY(user_id) REFERENCES users(id);
