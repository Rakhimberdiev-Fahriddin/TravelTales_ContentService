package postgres

import (
	"testing"

	pb "my_module/generated/content_service"

	"github.com/stretchr/testify/assert"
)

func TestCreateStories(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewStoriesRepo(db)

	req := pb.CreateStoriesRequest{
		Title:    "Test Story",
		Content:  "This is a test story.",
		Location: "Test Location",
		UserId:   "3315fb4a-a1cf-41df-9ca0-361b5da096dc",
		Tags:     []string{"tag1", "tag2"},
	}

	res, err := repo.CreateStories(&req)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Id)
	assert.Equal(t, req.Title, res.Title)
	assert.Equal(t, req.Content, res.Content)
	assert.Equal(t, req.Location, res.Location)
	assert.Equal(t, req.UserId, res.AuthorId)
	assert.Equal(t, len(req.Tags), len(res.Tags))
}

func TestUpdateStories(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewStoriesRepo(db)

	req := pb.UpdateStoriesRequest{
		Id:      "020d265d-863e-4b2b-b34d-db1aa893b7ba",
		Title:   "Updated Title",
		Content: "Updated Content",
	}

	res, err := repo.UpdateStories(&req)
	assert.NoError(t, err)
	assert.Equal(t, req.Id, res.Id)
	assert.Equal(t, req.Title, res.Title)
	assert.Equal(t, req.Content, res.Content)
}

func TestDeleteStories(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewStoriesRepo(db)

	req := pb.DeleteStoriesRequest{
		Id: "3d142a3b-f84d-4460-a341-eed9bd306dfe",
	}

	_,_ = repo.DeleteStories(&req)
}

func TestGetStories(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewStoriesRepo(db)

	req := pb.GetStoriesRequest{
		Id: "333b0da9-7c1c-4b19-9346-70f0962a767d",
	}

	res, err := repo.GetStories(&req)
	assert.NoError(t, err)
	assert.Equal(t, req.Id, res.Id)
}

func TestGetsStories(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewStoriesRepo(db)

	req := pb.GetsStoriesRequest{
		Page:  1,
		Limit: 10,
	}

	res, err := repo.GetsStories(&req)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Stories)
	assert.Equal(t, req.Page, res.Page)
	assert.Equal(t, req.Limit, res.Limit)
}

func TestCommentStories(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewStoriesRepo(db)

	req := pb.CommentStoriesRequest{
		StoryId:  "333b0da9-7c1c-4b19-9346-70f0962a767d",
		AuthorId: "ade85737-0bd0-40b1-a11d-6d8515a41879",
		Content:  "This is a test comment.",
	}

	res, err := repo.CommentStories(&req)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Id)
	assert.Equal(t, req.Content, res.Content)
}

func TestGetCommentStories(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewStoriesRepo(db)

	req := pb.GetCommentStoriesRequest{
		Page:  1,
		Limit: 10,
	}

	res, err := repo.GetCommentStories(&req)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Comment)
	assert.Equal(t, req.Page, res.Page)
	assert.Equal(t, req.Limit, res.Limit)
}

func TestLikeStories(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewStoriesRepo(db)

	req := pb.LikeStoriesRequest{
		StoriesId: "0dd0a88e-720e-42da-8efd-666a9d73362d",
	}

	res, err := repo.LikeStories(&req)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.StoryId)
}
