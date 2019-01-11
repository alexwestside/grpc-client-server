package main

import (
	"flag"
	"fmt"
	pb "github.com/grpc-client-server/server/data"
	"io"
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"strconv"
)

var (
	port = flag.String("p", "", "The server port")
)

type DataStreamServer struct {}

func (s *DataStreamServer) StreamData(stream pb.DataStream_StreamDataServer) (error) {
	var count int
	for {
		contact, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.ContactSummary{Summary: strconv.Itoa(count)})
		}
		if err != nil {
			return err
		}
		fmt.Println(contact)
		count++
		//for _, feature := range s.savedFeatures {
		//	if proto.Equal(feature.Location, point) {
		//		featureCount++
		//	}
		//}
	}
	return nil
}


func main() {
	flag.Parse()
	fmt.Println(*port)

	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDataStreamServer(s, &DataStreamServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}