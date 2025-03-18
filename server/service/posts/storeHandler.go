package posts

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/MahatVasudev/Computer-Vision-Website/server/msg"
	"github.com/MahatVasudev/Computer-Vision-Website/server/typestore"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (s *Store) GetPostDetail(
	ctx context.Context,
	post_id string,
) (*typestore.PostFullPicture, error) {

	query := `SELECT p.id, 
    title, 
    user_id, 
    description, 
    u.id, 
    u.username, 
    i.location, 
    p.createdat, 
    p.updatedat FROM posts as p 
  JOIN users as u
    on p.user_id = u.id
  JOIN images as i 
    on p.id = i.post_id
  WHERE p.id = $1
  `

	row, err := s.sqlDB.QueryContext(ctx, query, post_id)

	if err != nil {

		if errType, ok := err.(*pq.Error); ok {

			log.Println("Error Type Struct", errType)
			switch errType.Code {
			case pq.ErrorCode(msg.PQForeignKeyViolation):
				return nil, msg.ErrorNotFound

			case pq.ErrorCode(msg.PQErrUniqueKeyViolation):
				return nil, msg.ErrorConflict
			default:
				return nil, msg.ErrorServerSide
			}
		}

		return nil, err
	}

	defer row.Close()

	var Posts typestore.PostFullPicture

	for row.Next() {
		if err = row.Scan(
			&Posts.PostID,
			&Posts.Title,
			&Posts.UserID,
			&Posts.Description,
			&Posts.UserID,
			&Posts.UserName,
			&Posts.Images,
			&Posts.CreatedAt,
			&Posts.UpdatedAt,
		); err != nil {

			if errType, ok := err.(*pq.Error); ok {

				log.Println("Error Type Struct", errType)
				switch errType.Code {
				case pq.ErrorCode(msg.PQForeignKeyViolation):
					return nil, msg.ErrorNotFound

				case pq.ErrorCode(msg.PQErrUniqueKeyViolation):
					return nil, msg.ErrorConflict
				default:
					return nil, msg.ErrorServerSide
				}

			}

			return nil, err
		}
	}

	return &Posts, nil
}

func (s *Store) Get_All_Posts_From_User(
	ctx context.Context,
	user_name string,
	limit int,
) (*[]typestore.Post, error) {
	query := `
  SELECT p.id, p.title, i.location, u.username, p.createdat, p.updatedat FROM posts as p 
  JOIN users as u 
    on u.id = p.user_id
  JOIN images as i 
    on i.post_id = p.id 
  WHERE u.username = $1
  LIMIT $2
  `

	rows, err := s.sqlDB.QueryContext(ctx, query, user_name, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var Posts []typestore.Post

	for rows.Next() {
		p := new(typestore.Post)

		if err := rows.Scan(&p.ID, &p.Title, &p.Location, &p.Username, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}

		Posts = append(Posts, *p)
	}

	if len(Posts) == 0 {
		return &Posts, msg.ErrorNotFound
	}

	return &Posts, nil
}

func (s *Store) Get_All_Posts(ctx context.Context, limit int) (*[]typestore.Post, error) {

	query := `
  SELECT p.id, p.title, i.location, u.username, p.createdat, p.updatedat FROM posts as p 
  JOIN users as u 
    on u.id = p.user_id
  JOIN images as i 
    on i.post_id = p.id 
  LIMIT $1
  `

	rows, err := s.sqlDB.QueryContext(ctx, query, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var Posts []typestore.Post

	for rows.Next() {
		p := new(typestore.Post)

		if err := rows.Scan(&p.ID, &p.Title, &p.Location, &p.Username, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}

		Posts = append(Posts, *p)
	}

	if len(Posts) == 0 {
		return nil, msg.ErrorNotFound
	}

	return &Posts, nil

}

func (s *Store) CreatePost(ctx context.Context, p *typestore.PostFullPicture) error {

	query1 := `
  INSERT INTO posts(id, title, user_id, description)
  VALUES ($1, $2, $3, $4);
  `
	query2 := `

  INSERT INTO images(location, post_id) VALUES ($1, $2);

  `

	_, err := s.sqlDB.ExecContext(ctx, query1, p.PostID, p.Title, p.UserID, p.Description)

	if err != nil {
		return err
	}

	_, err = s.sqlDB.ExecContext(ctx, query2, p.Images, p.PostID)

	if err != nil {
		return err
	}

	return nil

}

func (s *Store) CountOfPostOfEachUser(ctx context.Context, user_id string) (int, error) {

	query := `
  SELECT count(*) FROM posts as p  
  JOIN users as u 
    on u.id = p.user_id
  WHERE u.username = $1
  group by user_id;
  `

	row := s.sqlDB.QueryRowContext(ctx, query, user_id)

	count := 0

	err := row.Scan(&count)

	if err != nil {
		return 0, msg.ErrorNotFound
	}

	return count, nil
}

func (s *Store) UploadImages(ctx context.Context, r *http.Request) (string, error) {

	err := r.ParseMultipartForm(10 << 20)

	if err != nil {
		return "", err
	}

	files, fileHeader, err := r.FormFile("images")

	if err != nil {
		return "", err
	}

	defer files.Close()

	if ok := IsSupportedMedia(fileHeader.Filename); !ok {
		return "", msg.ErrorBadRequest
	}

	outFileName := fmt.Sprintf("%s.%s", uuid.NewString(), GetExtention(fileHeader.Filename))

	outFile, err := os.Create(filepath.Join(s.postLib, outFileName))

	if err != nil {
		fmt.Println(err)
		return "", msg.ErrorServerSide
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, files)

	if err != nil {
		fmt.Println(err)
		return "", msg.ErrorServerSide
	}

	return outFileName, nil
}

func GetExtention(file string) string {

	splits := strings.Split(file, ".")
	extention := splits[len(splits)-1]

	return extention
}

func IsSupportedMedia(file string) bool {

	extention := GetExtention(file)
	if result, ok := ListOfSupportedImages[extention]; ok {
		return result
	}

	return false
}

type UploadedImage struct {
	FileName string
	Status   bool
}

var ListOfSupportedImages = map[string]bool{
	"jpg":  true,
	"jpeg": true,
	"png":  true,
	"webp": true,
	"mp4":  true, //support videos as well
}
