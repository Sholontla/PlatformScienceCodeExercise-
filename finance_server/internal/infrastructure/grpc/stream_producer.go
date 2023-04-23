package grpc

import (
	"finance_server/internal/domain/entity"
	pb "finance_server/internal/domain/proto"
	"finance_server/internal/domain/service"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/google/uuid"
)

func (s *Server) StreamStore(stream pb.StoreService_StreamStoreServer) error {
	log.Println("Stream server invoked ....")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return err
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		r := entity.Sale{
			Date:      req.Date,
			Product:   req.Product,
			Price:     float64(req.Price),
			Cost:      float64(req.Cost),
			UnitsSold: int(req.UnitSold),
			Region:    req.Region,
			SubRegion: req.SubRegion,
		}

		fmt.Println("Comming from StreamStore req: ", r)

		service.SaveData(&r)

		err = stream.Send(&pb.StoreResponseMessage{
			Id:       uuid.New().String(),
			Date:     "client 123",
			Response: time.Now().Local().String(),
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
