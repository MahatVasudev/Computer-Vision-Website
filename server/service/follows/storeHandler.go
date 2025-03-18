package follows

import (
	"context"
	"log"

	"github.com/MahatVasudev/Computer-Vision-Website/server/msg"
	"github.com/MahatVasudev/Computer-Vision-Website/server/typestore"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// GetAllFollowers implements store.FollowStore.
func (s *Store) GetAllFollowers(
	ctx context.Context,
	userid string,
) (typestore.FollowUserDetails, error) {
	panic("unimplemented")
}

func (s *Store) GetAggregate(
	ctx context.Context,
	username string,
) (*typestore.FollowingAggDetails, error) {
	query := `
  SELECT u.id, u.username, COALESCE(f.counts, 0) as followers, COALESCE(fo.counts, 0) as followings
  from users as u 
  left join (select follow_id, count(*) as counts from follows group by follow_id) as f 
    on f.follow_id = u.id
  left join (select follower_id, count(*) as counts from follows group by follower_id) as fo 
    on fo.follower_id = u.id
  where u.username = $1 
  `

	row := s.db.QueryRowContext(ctx, query, username)

	var fv typestore.FollowingAggDetails

	if err := row.Scan(&fv.Userid, &fv.Username, &fv.FollowerCount, &fv.FollowingCount); err != nil {

		log.Println(err)
		return nil, err

	}

	if fv.Username == "" {

		return nil, msg.ErrorNotFound
	}
	log.Println(fv)
	return &fv, nil

}

func (s *Store) FollowSomeOne(
	ctx context.Context,
	whosfollowingid string,
	tofollowid string,
) error {
	query := `INSERT INTO follows (id, follow_id, follower_id) VALUES($1,$2,$3);`

	_, err := s.db.ExecContext(ctx, query, uuid.NewString(), tofollowid, whosfollowingid)

	if err == nil {
		return nil
	}
	if errType, ok := err.(*pq.Error); ok {

		log.Println("Error Type Struct", errType)
		switch errType.Code {
		case pq.ErrorCode(msg.PQForeignKeyViolation):
			return msg.ErrorNotFound

		case pq.ErrorCode(msg.PQErrUniqueKeyViolation):
			return msg.ErrorConflict
		default:
			return msg.ErrorServerSide
		}
	}

	return msg.ErrorServerSide
}

func (s *Store) IsFollowingOrFollowed(
	ctx context.Context,
	userid string,
	tocheck_userid string,
) (bool, bool, error) {
	query := `SELECT 
    EXISTS (
        SELECT 1 FROM follows 
        JOIN users u1 ON follows.follower_id = u1.id 
        JOIN users u2 ON follows.follow_id = u2.id 
        WHERE u1.username = $1 AND u2.username = $2
    ) AS is_following,
    EXISTS (
        SELECT 1 FROM follows 
        JOIN users u1 ON follows.follower_id = u1.id 
        JOIN users u2 ON follows.follow_id = u2.id 
        WHERE u1.username = $2 AND u2.username = $1
    ) AS is_followed_by;`

	row := s.db.QueryRowContext(ctx, query, userid, tocheck_userid)

	is_following, is_followed_by := false, false

	if err := row.Scan(
		&is_following, &is_followed_by,
	); err != nil {

		if errType, ok := err.(*pq.Error); ok {
			log.Println("Error Type Struct", errType)
		}
		return false, false, msg.ErrorNotFound
	}

	return is_following, is_followed_by, nil

}
