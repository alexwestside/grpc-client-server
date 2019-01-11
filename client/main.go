package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "github.com/grpc-client-server/server/data"
	"context"
)

var (
	File   = flag.String("f", "", "The file containning the CA root cert file")
	Server = flag.String("s", "", "The server address in the format of host:port")
)

var ReadChann chan []string
var WriteChann chan *pb.ContactRequest
var Done chan bool
var Exit bool

func main() {
	fmt.Println(Exit)
	flag.Parse()
	fmt.Println(*File, *Server)

	WriteChann = make(chan *pb.ContactRequest)
	ReadChann = make(chan []string)

	go reader()
	go writer()

	conn, err := grpc.Dial(*Server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDataStreamClient(conn)

	stream, err := client.StreamData(context.Background())
	if err != nil {
		log.Fatalf("%v.RecordRoute(_) = _, %v", client, err)
	}

	for contact := range WriteChann {
		if err := stream.Send(contact); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, *contact, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v", reply)

}
