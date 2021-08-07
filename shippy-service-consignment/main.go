package main

import (
	"context"
	"errors"
	pb "github.com/kcwong395/go-micro/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/kcwong395/go-micro/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
	"log"
	"sync"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type consignmentService struct {
	repo         repository
	vesselClient vesselProto.VesselService
}

func (s *consignmentService) GetConsignments(ctx context.Context, request *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func (s *consignmentService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	vesselResponse, err := s.vesselClient.FindAvailable(ctx, &vesselProto.Specification{MaxWeight: req.Weight, Capacity: int32(len(req.Containers))})
	if vesselResponse == nil {
		return errors.New("error fetching vessel, returned nil")
	}
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = consignment
	return nil
}

func main() {
	repo := &Repository{}

	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	service.Init()
	vesselClient := vesselProto.NewVesselService("shippy.service.vessel", service.Client())

	err := pb.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo, vesselClient})
	if err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}

}
