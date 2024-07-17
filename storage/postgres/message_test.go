package postgres

import (
	pb "my_module/generated/content_service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewMessageRepo(db)

	req := pb.MessageRequest{
		RecipientId: "ade85737-0bd0-40b1-a11d-6d8515a41879",
		Content:     "Test Message",
		SenderId:    "2d194fd1-3fc9-4490-bd7a-4ad794acc207",
	}

	res, err := repo.Message(&req)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Id)
	assert.Equal(t, req.RecipientId, res.RecipientId)
	assert.Equal(t, req.Content, res.Content)
	assert.Equal(t, req.SenderId, res.SenderId)
}

func TestGetMessage(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewMessageRepo(db)

	req := pb.GetMessageRequest{
		Page:  1,
		Limit: 10,
	}

	res, err := repo.GetMessage(&req)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Messages)
	assert.Equal(t, req.Page, res.Page)
	assert.Equal(t, req.Limit, res.Limit)
	assert.True(t, len(res.Messages) <= int(req.Limit))
}
