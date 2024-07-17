package postgres

import (
	pb "my_module/generated/content_service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateItineraries(t *testing.T) {
	db, err := ConnectDB() 
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewItinerariesRepo(db)

	req := pb.CreateItinerariesRequest{
		Title:       "Test Itinerary",
		Description: "Test Description",
		StartDate:   "2024-07-20",
		EndDate:     "2024-07-25",
		AuthorId:    "ade85737-0bd0-40b1-a11d-6d8515a41879",
		Destinations: []*pb.Destination{
			{
				Name:      "Paris",
				StartDate: "2024-07-20",
				EndDate:   "2024-07-22",
				Activities: []string{
					"Eiffel Tower Visit",
					"Louvre Museum",
				},
			},
		},
	}

	res, err := repo.CreateItineraries(&req)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Id)
	assert.Equal(t, req.Title, res.Title)
	assert.Equal(t, req.Description, res.Description)
}

func TestUpdateItineraries(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewItinerariesRepo(db)

	req := pb.UpdateItinerariesRequest{
		Id:          "50c782fb-e891-480a-a09f-3007ef2d937c",
		Title:       "Updated Itinerary",
		Description: "Updated Description",
	}

	res, err := repo.UpdateItineraries(&req)
	assert.NoError(t, err)
	assert.Equal(t, req.Id, res.Id)
	assert.Equal(t, req.Title, res.Title)
	assert.Equal(t, req.Description, res.Description)
}

func TestDeleteItineraries(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewItinerariesRepo(db)

	req := pb.DeleteItinerariesRequest{
		Id: "50c782fb-e891-480a-a09f-3007ef2d937c",
	}

	res, err := repo.DeleteItineraries(&req)
	assert.NoError(t, err)
	assert.Equal(t, "Succesfully deleted Itineraries", res.Message)
}

func TestGetsItineraries(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
		return
	}
	defer db.Close()

	repo := NewItinerariesRepo(db)

	req := pb.GetsItinerariesRequest{
		Page:  1,
		Limit: 10,
	}

	res, err := repo.GetsItineraries(&req)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Itineraries)
	assert.Equal(t, req.Page, res.Page)
	assert.Equal(t, req.Limit, res.Limit)
}
