package service

import (
	pb "my_module/generated/content_service"

	"golang.org/x/net/context"
)

func (t *StoriesService) AddTravelTips(ctx context.Context,req *pb.AddTravelTipsRequest) (*pb.AddTravelTipsResponce, error) {
	t.Logger.Info("Add Travel Tips")
	res, err := t.TravelTips.AddTravelTips(req)
	if err != nil {
		t.Logger.Error("Failed add travel tips", "error", err.Error())
		return nil, err
	}
	return res, nil
}

func (t *StoriesService) GetTravelTips(ctx context.Context,req *pb.GetTravelTipsRequest) (*pb.GetTravelTipsResponce, error) {
	t.Logger.Info("Get Travel Tips")
	res, err := t.TravelTips.GetTravelTips(req)
	if err != nil {
		t.Logger.Error("Failed get travel tips", "error", err.Error())
		return nil, err
	}
	return res, err
}

func (t *StoriesService) StatisticsUser(ctx context.Context,req *pb.StatisticsUserRequest) (*pb.StatisticsUserResponce, error) {
	t.Logger.Info("Statistics User")
	res, err := t.TravelTips.StatisticsUser(req)
	if err != nil {
		t.Logger.Error("Failed statistics user", "error", err.Error())
		return nil, err
	}
	return res, nil
}

func (t *StoriesService) GetTrendingDestinations(ctx context.Context,req *pb.GetTrendingDestinationsRequest) (*pb.GetTrendingDestinationsResponce, error) {
	t.Logger.Info("Get Trending Destinations")
	res, err := t.TravelTips.GetTrendingDestinations(req)
	if err != nil {
		t.Logger.Error("Failed get trending destinations", "error", err.Error())
		return nil, err
	}
	return res, nil
}
