package clientgrpc

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"publisher_service/internal/domain/entity"
	pb "publisher_service/internal/domain/proto"
)

type GrpcClient struct{}

// pb.LoggsProducerServiceClient
// c.StreamProducer(context.Background())
// MainLoggerRequest
func StreamSendLogs(c pb.StoreServiceClient, store entity.Order) {
	log.Println("gRPC LOG Client invoked ...")

	stream, err := c.StreamStore(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.StoreRequestMessage{
		{
			Id:        store.Store.ID.String(),
			Date:      store.Store.Sale.Date,
			Product:   store.Store.Sale.Product,
			Price:     float32(store.Store.Sale.Price),
			Cost:      float32(store.Store.Sale.Cost),
			UnitSold:  int32(store.Store.Sale.UnitsSold),
			Region:    store.Store.Sale.Region,
			SubRegion: store.Store.Sale.SubRegion,
		},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send Request: %v\n", req)
			stream.Send(req)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while reciving from Server: %v\n", err)
				break
			}

			fmt.Println("Recived: ", res)
		}

		close(waitc)
	}()

	<-waitc
}

var addr string = "finance_server_backend:50051"

func (g GrpcClient) GrcpClient(store entity.Order) {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect Client: %v\n", err)
	}
	defer conn.Close()
	log.Printf("GRPC Client Listening on %s\n: ", addr)

	c := pb.NewStoreServiceClient(conn)
	StreamSendLogs(c, store)
}
