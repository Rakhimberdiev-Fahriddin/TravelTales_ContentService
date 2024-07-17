package service

import (
	pb "my_module/generated/content_service"

	"golang.org/x/net/context"
)

func (m *StoriesService) Message(ctx context.Context,req *pb.MessageRequest) (*pb.MessageResponce, error) {
	m.Logger.Info("Message")
	res, err := m.Messag.Message(req)
	if err != nil {
		m.Logger.Error("Failed message", "error", err.Error())
		return nil, err
	}
	return res, err
}

func (m *StoriesService) GetMessage(ctx context.Context,req *pb.GetMessageRequest) (*pb.GetMessageResponce, error) {
	m.Logger.Info("Get Message")
	res, err := m.Messag.GetMessage(req)
	if err != nil {
		m.Logger.Error("Failed get message", "error", err.Error())
		return nil, err
	}
	return res, nil
}
