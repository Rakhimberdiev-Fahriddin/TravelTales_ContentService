package postgres

import (
	"database/sql"
	pb "my_module/generated/content_service"
)

type StoriesRepo struct {
	DB *sql.DB
}

func NewStoriesRepo(db *sql.DB) *StoriesRepo {
	return &StoriesRepo{DB: db}
}

func (s *StoriesRepo) CreateStories(reqStories *pb.CreateStoriesRequest) (*pb.CreateStoriesResponce, error) {
	resStories := pb.CreateStoriesResponce{}
	err := s.DB.QueryRow(`
	INSERT INTO 
		Stories(
			title,
			content,
			location,
			author_id
			)
		VALUES(
			$1,
			$2,
			$3,
			$4
			)
		Returning
			id,
			title,
			content,
			location,
			author_id,
			created_at`,
		reqStories.Title,
		reqStories.Content,
		reqStories.Location,
		reqStories.UserId,
	).Scan(
		&resStories.Id,
		&resStories.Title,
		&resStories.Content,
		&resStories.Location,
		&resStories.AuthorId,
		&resStories.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	var tags []*pb.Tag
	for _, v := range reqStories.Tags {
		var tag pb.Tag
		err := s.DB.QueryRow(`
		INSERT INTO
			Story_tags(
				story_id,
				tag
			)
			VALUES(
				$1,
				$2
			)
			Returning
				story_id,
				tag
			`,
			resStories.Id,
			v).Scan(
			&tag.StoryId,
			&tag.Tag,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, &tag)
	}
	resStories.Tags = tags
	return &resStories, nil
}

func (s *StoriesRepo) UpdateStories(updateStories *pb.UpdateStoriesRequest) (*pb.UpdateStoriesResponce, error) {
	resUpdate := pb.UpdateStoriesResponce{}
	err := s.DB.QueryRow(`
	UPDATE 
		Stories
	SET
		title = $1,
		content = $2
	WHERE
		id = $3
	Returning
		id,
		title,
		content,
		location,
		author_id,
		updated_at
	`,
		updateStories.Title,
		updateStories.Content,
		updateStories.Id).Scan(
		&resUpdate.Id,
		&resUpdate.Title,
		&resUpdate.Content,
		&resUpdate.Location,
		&resUpdate.AuthorId,
		&resUpdate.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	rows, err := s.DB.Query(`
	SELECT 
		story_id,
		tag
	FROM
		Story_tags
	WHERE
		story_id = $1
	`,resUpdate.Id)

	if err != nil {
		return nil, err
	}

	var tags []*pb.Tag

	for rows.Next() {
		var tag pb.Tag
		err := rows.Scan(&tag.StoryId, &tag.Tag)

		if err != nil {
			return nil, err
		}

		tags = append(tags, &tag)
	}
	resUpdate.Tags = tags

	return &resUpdate, nil
}

func (s *StoriesRepo) DeleteStories(id *pb.DeleteStoriesRequest) (*pb.DeleteStoriesResponce, error) {
	_, err := s.DB.Exec(`
	DELETE FROM 
		Stories
	WHERE
		id = $1
	`, id.Id)

	if err != nil {
		return &pb.DeleteStoriesResponce{
			Message: "Failed deleted Stories",
		}, err
	}

	return &pb.DeleteStoriesResponce{
		Message: "Succesfully deleted Stories",
	}, nil
}

func (s *StoriesRepo) GetStories(id *pb.GetStoriesRequest) (*pb.GetStoriesResponce, error) {
	resStories := pb.GetStoriesResponce{}

	err := s.DB.QueryRow(`
	SELECT
		id,
		title,
		content,
		location,
		likes_count,
		comments_count,
		created_at,
		updated_at
	FROM
		Stories
	WHERE
		id = $1
		`, id.Id).Scan(
		&resStories.Id,
		&resStories.Title,
		&resStories.Content,
		&resStories.Location,
		&resStories.LikesCount,
		&resStories.CommentsCount,
		&resStories.CreatedAt,
		&resStories.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	rows, err := s.DB.Query(`
	SELECT 
		story_id,
		tag
	FROM
		Story_tags
	WHERE
		story_id = $1
	`, id.Id)

	if err != nil {
		return nil, err
	}

	var tags []*pb.Tag

	for rows.Next() {
		var tag pb.Tag
		err := rows.Scan(&tag.StoryId, &tag.Tag)

		if err != nil {
			return nil, err
		}

		tags = append(tags, &tag)
	}

	resStories.Tags = tags

	var user_id string
	err = s.DB.QueryRow(`
	SELECT
		author_id
	FROM
		Stories
	WHERE
		id = $1
	`, id.Id).Scan(&user_id)

	if err != nil {
		return nil, err
	}

	var userName string
	err = s.DB.QueryRow(`
	SELECT
		username
	FROM
		users
	WHERE
		id = $1
	`, user_id).Scan(&userName)

	if err != nil {
		return nil, err
	}

	author := pb.Author{}

	author.Id = user_id
	author.UserName = userName

	resStories.Author = &author

	return &resStories, nil
}

func (s *StoriesRepo) GetsStories(reqStories *pb.GetsStoriesRequest) (*pb.GetsStoriesResponce, error) {
	rows, err := s.DB.Query(`
		SELECT
			id,
			title,
			content,
			location,
			likes_count,
			comments_count,
			created_at,
			updated_at
		FROM
			Stories
		OFFSET $1
		LIMIT $2
	`, (reqStories.Page-1)*reqStories.Limit, reqStories.Limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stories []*pb.GetStoriesResponce
	for rows.Next() {
		var storie pb.GetStoriesResponce
		err = rows.Scan(
			&storie.Id,
			&storie.Title,
			&storie.Content,
			&storie.Location,
			&storie.LikesCount,
			&storie.CommentsCount,
			&storie.CreatedAt,
			&storie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		tagRows, err := s.DB.Query(`
			SELECT
				story_id,
				tag
			FROM
				Story_tags
			WHERE
				story_id = $1
		`, storie.Id)

		if err != nil {
			return nil, err
		}
		defer tagRows.Close()

		var tags []*pb.Tag
		for tagRows.Next() {
			var tag pb.Tag
			err := tagRows.Scan(&tag.StoryId, &tag.Tag)
			if err != nil {
				return nil, err
			}
			tags = append(tags, &tag)
		}
		storie.Tags = tags

		var user_id string
		err = s.DB.QueryRow(`
			SELECT
				author_id
			FROM
				Stories
			WHERE
				id = $1
		`, storie.Id).Scan(&user_id)

		if err != nil {
			return nil, err
		}

		var userName string
		err = s.DB.QueryRow(`
			SELECT
				username
			FROM
				users
			WHERE
				id = $1
		`, user_id).Scan(&userName)

		if err != nil {
			return nil, err
		}

		author := pb.Author{
			Id:       user_id,
			UserName: userName,
		}
		storie.Author = &author

		stories = append(stories, &storie)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var total int32
	err = s.DB.QueryRow(`
		SELECT
			COUNT(*)
		FROM
			Stories
	`).Scan(&total)

	if err != nil {
		return nil, err
	}

	return &pb.GetsStoriesResponce{
		Stories: stories,
		Total:   total,
		Page:    reqStories.Page,
		Limit:   reqStories.Limit,
	}, nil
}


func (s *StoriesRepo) CommentStories(reqComment *pb.CommentStoriesRequest) (*pb.CommentStoriesResponce, error) {
	resComment := pb.CommentStoriesResponce{}
	err := s.DB.QueryRow(`
	INSERT INTO
		Comments(
			story_id,
			author_id,
			content	
		)
		VALUES(
			$1,
			$2,
			$3	
		)
		Returning
			id,
			content,
			author_id,
			story_id,
			created_at
	`, reqComment.StoryId,reqComment.AuthorId,reqComment.Content).Scan(
		&resComment.Id,
		&resComment.Content,
		&resComment.AuthorId,
		&resComment.StoryId,
		&resComment.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &resComment, nil
}

func (s *StoriesRepo) GetCommentStories(reqComment *pb.GetCommentStoriesRequest) (*pb.GetCommentStoriesResponce, error) {
	rows, err := s.DB.Query(`
	SELECT
		id,
		content,
		created_at
	FROM
		Comments
	OFFSET $1
	LIMIT $2
	`, (reqComment.Page-1)*reqComment.Limit, reqComment.Limit)

	if err != nil {
		return nil, err
	}

	var comments []*pb.Comment
	for rows.Next() {
		var comment pb.Comment
		err = rows.Scan(&comment.Id, &comment.Content, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}

		var storiesId string

		err = s.DB.QueryRow(`
		SELECT
			story_id
		FROM
			Comments
		WHERE
			id = $1
		`, comment.Id).Scan(&storiesId)

		if err != nil{
			return nil,err
		}

		var user_id string
		err = s.DB.QueryRow(`
		SELECT
			author_id
		FROM
			Stories
		WHERE
			id = $1
		`, storiesId).Scan(&user_id)

		if err != nil {
			return nil, err
		}

		var userName string
		err = s.DB.QueryRow(`
		SELECT
			username
		FROM
			users
		WHERE
			id = $1
		`, user_id).Scan(&userName)

		if err != nil {
			return nil, err
		}

		author := pb.Author{}

		author.Id = user_id
		author.UserName = userName

		comment.Author = &author

		comments = append(comments, &comment)
	}

	var total int32
	err = s.DB.QueryRow(`
	SELECT
		COUNT(*)
	FROM
		Comments
	`).Scan(&total)

	if err != nil {
		return nil, err
	}

	return &pb.GetCommentStoriesResponce{
		Comment: comments,
		Total:   total,
		Page:    reqComment.Page,
		Limit:   reqComment.Limit,
	}, nil
}

func (s *StoriesRepo) LikeStories(likeStories *pb.LikeStoriesRequest) (*pb.LikeStoriesResponce, error) {
	var user_id string
	err := s.DB.QueryRow(`
	SELECT 
		author_id
	FROM
		Stories
	WHERE 
		id = $1
	`, likeStories.StoriesId).Scan(&user_id)

	if err != nil {
		return nil, err
	}

	like := pb.LikeStoriesResponce{}

	err = s.DB.QueryRow(`
	INSERT INTO
		Likes(
			story_id,
			user_id 
		)
		VALUES(
			$1,
			$2
			)
		Returning
			story_id,
			user_id,
			Liked_at	
	`, likeStories.StoriesId, user_id).Scan(
		&like.StoryId,
		&like.UserId,
		&like.LikedAt,
	)

	if err != nil {
		return nil, err
	}

	return &like, nil
}
