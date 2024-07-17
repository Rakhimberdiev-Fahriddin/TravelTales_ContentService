package service

import (
	pb "my_module/generated/content_service"

	"golang.org/x/net/context"
)

func (d *StoriesService) GetsDestinations(ctx context.Context, req *pb.GetsDestinationsRequest) (*pb.GetsDestinationsResponce, error) {
	d.Logger.Info("Gets Destinations")

	res, err := d.Destinations.GetsDestinations(req)
	if err != nil {
		d.Logger.Error("Failed gets destinations", "error", err.Error())
		return nil, err
	}
	return res, nil
}

func (d *StoriesService) GetDestinations(ctx context.Context,req *pb.GetDestinationsRequest) (*pb.GetDestinationsResponce, error) {
	d.Logger.Info("Get Destinations")

	res, err := d.Destinations.GetDestinations(req)
	if err != nil {
		d.Logger.Error("Failed get destinations", "error", err.Error())
		return nil, err
	}
	return res, nil
}
