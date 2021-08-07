package main

import (
	"context"
	"encoding/json"
	pb "github.com/kcwong395/go-micro/shippy-service-consignment/proto/consignment"
	"io/ioutil"
	"log"
	"os"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	service := micro.NewService(micro.Name("shippy.consignment.cli"))
	service.Init()

	// The original tutorial should have a typo here,
	// it should be shippy.service.consignment rather than shippy.consignment.service
	client := pb.NewShippingService("shippy.service.consignment", service.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	resp, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range resp.Consignments {
		log.Println(v)
	}
}
