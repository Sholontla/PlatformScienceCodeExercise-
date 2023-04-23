package logger

// import (
// 	"context"
// 	"fmt"
// 	"io"
// 	"log"

// 	"github.com/google/uuid"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"

// 	pb "service/producer/demo/internal/proto/loggs"
// )

// func StreamSendLogs(c pb.LoggsProducerServiceClient, logServ string) {
// 	log.Println("gRPC LOG Client invoked ...")

// 	stream, err := c.StreamProducer(context.Background())

// 	if err != nil {
// 		log.Fatalf("Error while creating stream: %v\n", err)
// 	}

// 	reqs := []*pb.MainLoggerRequest{
// 		{LoggerId: uuid.New().String(), LoggerMessage: logServ},
// 	}

// 	waitc := make(chan struct{})

// 	go func() {
// 		for _, req := range reqs {
// 			log.Printf("Send Request: %v\n", req)
// 			stream.Send(req)
// 		}

// 		stream.CloseSend()
// 	}()

// 	go func() {
// 		for {
// 			res, err := stream.Recv()
// 			if err == io.EOF {
// 				break
// 			}

// 			if err != nil {
// 				log.Printf("Error while reciving from Server: %v\n", err)
// 				break
// 			}

// 			fmt.Println("Recived: ", res)
// 		}

// 		close(waitc)
// 	}()

// 	<-waitc
// }

// var addr string = "logger_backend:50010"

// func GrcpClient(logg string) {

// 	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

// 	if err != nil {
// 		log.Fatalf("Failed to connect Client: %v\n", err)
// 	}
// 	defer conn.Close()
// 	log.Printf("GRPC Client Listening on %s\n: ", addr)

// 	fmt.Println("LOG .........................: ", logg)
// 	c := pb.NewLoggsProducerServiceClient(conn)

// 	StreamSendLogs(c, logg)
// }
