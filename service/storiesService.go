package service

import (
	"database/sql"
	"log/slog"
	pb "my_module/generated/content_service"
	"my_module/logs"
	"my_module/storage/postgres"

	"golang.org/x/net/context"
)

type StoriesService struct {
	pb.UnimplementedContentServer
	Stories      postgres.StoriesRepo
	Destinations postgres.DestinationsRepo
	Itineraries  postgres.ItinerariesRepo
	Messag       postgres.MessageRepo
	TravelTips   postgres.TravelTipsRepo
	Logger       *slog.Logger
}

func NewStoriesService(db *sql.DB) *StoriesService {
	logs.InitLogger()
	return &StoriesService{
		Stories:      *postgres.NewStoriesRepo(db),
		Destinations: *postgres.NewDestinationsRepo(db),
		Itineraries:  *postgres.NewItinerariesRepo(db),
		Messag:       *postgres.NewMessageRepo(db),
		TravelTips:   *postgres.NewTravelTipsRepo(db),
		Logger: logs.Logger,
	}
}

func (s *StoriesService) CreateStories(ctx context.Context,req *pb.CreateStoriesRequest) (*pb.CreateStoriesResponce, error) {
	s.Logger.Info("Create Stories")

	res, err := s.Stories.CreateStories(req)
	if err != nil {
		s.Logger.Error("Failed create stories", "error", err.Error())
		return nil, err
	}
	return res, nil
}

func (s *StoriesService) UpdateStories(ctx context.Context,req *pb.UpdateStoriesRequest) (*pb.UpdateStoriesResponce, error) {
	s.Logger.Info("Update Stories")
	res, err := s.Stories.UpdateStories(req)
	if err != nil {
		s.Logger.Error("Failed update stories", "error", err.Error())
		return nil, err
	}
	return res, nil
}

func (s *StoriesService) DeleteStories(ctx context.Context,req *pb.DeleteStoriesRequest) (*pb.DeleteStoriesResponce, error) {
	s.Logger.Info("Delete Stories")
	res, err := s.Stories.DeleteStories(req)
	if err != nil {
		s.Logger.Error("Failed Deleled stories", "error", err.Error())
		return res, err
	}
	return res, nil
}

func (s *StoriesService) GetStories(ctx context.Context,req *pb.GetStoriesRequest) (*pb.GetStoriesResponce, error) {
	s.Logger.Info("Get Stories")
	res, err := s.Stories.GetStories(req)
	if err != nil {
		s.Logger.Error("Failed get stories", "error", err.Error())
		return nil, err
	}
	return res, nil
}

func (s *StoriesService) GetsStories(ctx context.Context,req *pb.GetsStoriesRequest) (*pb.GetsStoriesResponce, error) {
	s.Logger.Info("Gets Stories")
	res, err := s.Stories.GetsStories(req)
	if err != nil {
		s.Logger.Error("Failed gets stories", "error", err.Error())
		return nil, err
	}
	return res, nil
}

func (s *StoriesService) CommentStories(ctx context.Context,req *pb.CommentStoriesRequest) (*pb.CommentStoriesResponce, error) {
	s.Logger.Info("Comment Stories")
	res, err := s.Stories.CommentStories(req)
	if err != nil {
		s.Logger.Error("Failed comment stories", "error", err.Error())
		return nil, err
	}
	return res, nil
}

func (s *StoriesService) GetCommentStories(ctx context.Context,req *pb.GetCommentStoriesRequest) (*pb.GetCommentStoriesResponce, error) {
	s.Logger.Info("Get Comment Stories")
	res, err := s.Stories.GetCommentStories(req)
	if err != nil {
		s.Logger.Error("Failed get comment stories", "error", err.Error())
		return nil, err
	}
	return res, nil
}

func (s *StoriesService) LikeStories(ctx context.Context,req *pb.LikeStoriesRequest) (*pb.LikeStoriesResponce, error) {
	s.Logger.Info("Likes Stories")
	res, err := s.Stories.LikeStories(req)
	if err != nil {
		s.Logger.Error("Failed likes stories", "error", err.Error())
		return nil, err
	}
	return res, nil
}
