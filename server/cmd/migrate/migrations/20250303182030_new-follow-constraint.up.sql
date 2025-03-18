ALTER TABLE follows
ADD CONSTRAINT follower_pair UNIQUE (follower_id, follow_id)
