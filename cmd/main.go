package main

import (
	"log"
	"my_module/config"
	"my_module/generated/content_service"
	"my_module/logs"
	"my_module/service"
	"my_module/storage/postgres"
	"net"

	"google.golang.org/grpc"
)


func main(){
	logs.InitLogger()
	logs.Logger.Info("Starting the Server")
	db,err := postgres.ConnectDB()
	
	if err != nil{
		logs.Logger.Error("Failed connect to DATA BASE","error",err.Error())
		panic(err)
	}
	defer db.Close()

	config := config.Load()

	listener,err := net.Listen("tcp",":"+config.GRPC_PORT)
	if err != nil{
		logs.Logger.Error("Failed listen","error",err.Error())
		panic(err)
	}

	contentService := service.NewStoriesService(db)

	s := grpc.NewServer()

	content_service.RegisterContentServer(s,contentService)

	logs.Logger.Info("Server is Running","PORT",config.GRPC_PORT)
	log.Println("Server is running ",listener.Addr())

	if err := s.Serve(listener);err != nil{
		logs.Logger.Error("Failed server is runing ","error",err.Error())
		panic(err)
	}
}