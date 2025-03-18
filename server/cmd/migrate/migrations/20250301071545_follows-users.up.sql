CREATE TABLE IF NOT EXISTS follows (
  id varchar(255) primary key,
  follow_id varchar(255) not null,
  follower_id varchar(255) not null,
  createdat TIMESTAMP default CURRENT_TIMESTAMP,
  updatedat TIMESTAMP default CURRENT_TIMESTAMP,

  CONSTRAINT fk_follow_id FOREIGN KEY(follow_id) REFERENCES users(id),
  CONSTRAINT fk_follower_id FOREIGN KEY(follower_id) REFERENCES users(id),
  CONSTRAINT na_follow_follower CHECK ( follow_id != follower_id )
);

CREATE OR REPLACE TRIGGER update_updatedat 
BEFORE UPDATE ON follows FOR EACH ROW EXECUTE FUNCTION update_updated_at();
