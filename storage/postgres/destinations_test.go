package postgres

import (
	"fmt"
	"testing"

	pb "my_module/generated/content_service"

	"github.com/stretchr/testify/assert"
)

func TestGetsDestinations(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewDestinationsRepo(db)

	req := pb.GetsDestinationsRequest{
		Query: "sample destination",
		Page:  1,
		Limit: 10,
	}

	res, err := repo.GetsDestinations(&req)
	fmt.Println(err)
	assert.NoError(t, err)
	assert.Equal(t, req.Page, res.Page)
	assert.Equal(t, req.Limit, res.Limit)
}

func TestGetDestinations(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewDestinationsRepo(db)

	req := pb.GetDestinationsRequest{
		Id: "950dd78b-5e7d-4edd-807a-01a41e7014cf",
	}

	res, err := repo.GetDestinations(&req)
	fmt.Println(err)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Id)
	assert.Equal(t, req.Id, res.Id)
}