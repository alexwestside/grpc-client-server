package main

import (
	pb "github.com/grpc-client-server/server/data"
)

func writer() {

	for record := range ReadChann {
		WriteChann <- &pb.ContactRequest{
			Name:         record[1],
			Email:        record[2],
			MobileNumber: record[3],
		}
	}

	close(WriteChann)
}
