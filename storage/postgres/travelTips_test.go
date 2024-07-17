package postgres

import (
	pb "my_module/generated/content_service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTravelTips(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewTravelTipsRepo(db)

	req := pb.AddTravelTipsRequest{
		Title: "fdkasdnvjfk",
		Content: "afhkusdlsfan",
		Category: "Bali",
		AuthorId: "ade85737-0bd0-40b1-a11d-6d8515a41879",
	}

	res,err := repo.AddTravelTips(&req)
	
	assert.NoError(t,err)
	assert.NotEmpty(t,res)
}

func TestGetTravelTips(t *testing.T){
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewTravelTipsRepo(db)

	req := pb.GetTravelTipsRequest{
		Category: "Bali",
		Page: 1,
		Limit: 10,
	}

	res,err := repo.GetTravelTips(&req)
	assert.NoError(t,err)
	assert.NotEmpty(t,res)
}

func TestStatisticsUser(t *testing.T){
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewTravelTipsRepo(db)

	req := pb.StatisticsUserRequest{
		UserId: "2d194fd1-3fc9-4490-bd7a-4ad794acc207",
		TotalCountriesVisited: 35,
	}

	res,err := repo.StatisticsUser(&req)

	assert.NoError(t,err)
	assert.NotEmpty(t,res)
}

func TestGetTrendingDestinations(t *testing.T){
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewTravelTipsRepo(db)

	req := pb.GetTrendingDestinationsRequest{}

	res,err := repo.GetTrendingDestinations(&req)

	assert.NoError(t,err)
	assert.NotEmpty(t,res)
}