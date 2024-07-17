package service

import (
	pb "my_module/generated/content_service"

	"golang.org/x/net/context"
)

func (i *StoriesService) CreateItineraries(ctx context.Context,req *pb.CreateItinerariesRequest) (*pb.CreateItinerariesResponce, error) {
	i.Logger.Info("Create Itineraries")
	res, err := i.Itineraries.CreateItineraries(req)
	if err != nil {
		i.Logger.Error("Failed create itineraries", "error", err.Error())
		return nil, err
	}
	return res, nil
}

func (i *StoriesService) UpdateItineraries(ctx context.Context,req *pb.UpdateItinerariesRequest) (*pb.UpdateItinerariesResponce, error) {
	i.Logger.Info("Update Itineraries")
	res, err := i.Itineraries.UpdateItineraries(req)
	if err != nil {
		i.Logger.Error("Failed update itineraries", "error", err.Error())
		return nil, err
	}
	return res, nil
}

func (i *StoriesService) DeleteItineraries(ctx context.Context,req *pb.DeleteItinerariesRequest) (*pb.DeleteItinerariesResponce, error) {
	i.Logger.Info("Deleted Itineraries")
	res, err := i.Itineraries.DeleteItineraries(req)
	if err != nil {
		i.Logger.Error("Failed deleted itineraries", "error", err.Error())
		return nil, err
	}
	return res, nil
}

func (i *StoriesService) GetsItineraries(ctx context.Context,req *pb.GetsItinerariesRequest) (*pb.GetsItinerariesResponce, error) {
	i.Logger.Info("Gets Itineraries")
	res, err := i.Itineraries.GetsItineraries(req)

	if err != nil {
		i.Logger.Error("Failed gets itineraries", "error", err.Error())
		return nil, err
	}
	return res, nil
}

func (i *StoriesService) GetItineraries(ctx context.Context,req *pb.GetItinerariesRequest) (*pb.GetItinerariesResponce, error) {
	i.Logger.Info("Get Itineraries")
	res, err := i.Itineraries.GetItineraries(req)
	if err != nil {
		i.Logger.Error("Failed get itineraries", "error", err.Error())
		return nil, err
	}
	return res, nil
}
