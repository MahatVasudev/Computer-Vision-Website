package comments

import (
	"context"
	"log"

	"github.com/MahatVasudev/Computer-Vision-Website/server/msg"
	"github.com/MahatVasudev/Computer-Vision-Website/server/typestore"
)

func (s *Store) AmountOfCommentsInPost(ctx context.Context, post_id string) int {
	query := `SELECT COUNT(*) FROM comments where postid = $1 group by postid;`

	row := s.db.QueryRowContext(ctx, query, post_id)

	counts := 0

	if err := row.Scan(&counts); err != nil {
		log.Println(err)
		return 0
	}

	return counts

}

func (s *Store) CreateNewComment(
	ctx context.Context,
	post_id string,
	commented string,
	user_id string,
) error {
	query := `
  INSERT INTO comments(commented, postid, userid ) VALUES ($1, $2, $3)
  `

	_, err := s.db.ExecContext(ctx, query, commented, post_id, user_id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetCommentsOfPosts(
	ctx context.Context,
	post_id string,
) (*[]typestore.Comment, error) {

	query := `
  SELECT c.id, c.commented, c.userid, u.username FROM comments as c 
  JOIN users as u 
    on u.id = c.user_id
  WHERE postid = $1
  `

	rows, err := s.db.QueryContext(ctx, query, post_id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Comments := []typestore.Comment{}

	for rows.Next() {
		c := new(typestore.Comment)
		if err := rows.Scan(&c.Id, &c.Commented, &c.UserId, &c.Username); err != nil {
			return nil, err
		}

		Comments = append(Comments, *c)
	}

	if len(Comments) == 0 {
		return nil, msg.ErrorNotFound
	}

	return &Comments, nil
}
