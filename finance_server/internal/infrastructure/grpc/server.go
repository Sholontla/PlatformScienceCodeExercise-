package grpc

import (
	pb "finance_server/internal/domain/proto"
	"finance_server/internal/domain/service"
)

type Server struct {
	pb.StoreServiceServer
	service service.CacheService
}
