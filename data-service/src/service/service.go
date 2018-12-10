package service

import (
	"log"
	"net"
	"time"

	"github.com/atyagi9006/certificationapp/data-service/src/config"
	"github.com/atyagi9006/certificationapp/data-service/src/dao"
	"github.com/atyagi9006/certificationapp/data-service/src/handlers"
	"github.com/atyagi9006/certificationapp/grpcproto"
	"google.golang.org/grpc"
)

func DbInit() {
	db := config.Init()
	log.Println("Connted DB : ", db.DBConfig.DatabaseName)
	dao.Init(db)
}

func StartServer(port string) {
	log.Println("Starting DATA-Service at : ", port)
	time.Sleep(5 * time.Second)

	log.Println("Serving...")
	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	registerServices(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func registerServices(s *grpc.Server) {
	grpcproto.RegisterCandidateServiceServer(s, &handlers.Server{})
	grpcproto.RegisterUserServiceServer(s, &handlers.Server{})
	grpcproto.RegisterTestServiceServer(s, &handlers.Server{})
}
