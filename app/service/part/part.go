package part

import (
	"context"
	"database/sql"
	//	"github.com/golang/protobuf/ptypes"
	_ "github.com/lib/pq"
	prt "github.com/pprisn/grpc_rest_mnf/api/mnf/v1"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	//	"time"
)

type Service struct {
	DB  *sql.DB
	LOG *logrus.Logger
}

// CreatePart creates
func (s Service) CreatePart(ctx context.Context, req *prt.CreatePartRequest) (*prt.CreatePartResponse, error) {
	p, _ := peer.FromContext(ctx)
	req.Item.Id = uuid.NewV4().String()
	u := &prt.Part{}
	if err := s.DB.QueryRow(
		"INSERT INTO part (id, mnf_id, vendor_code, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id",
		req.Item.Id,
		req.Item.MnfId,
		req.Item.VendorCode,
	).Scan(&u.Id); err != nil {
		s.LOG.Fatalf("Peer:%s ERROR INSERT part (mnf_id, vendor_code) values (%s, %s) err: %s", p.Addr.String(), req.Item.MnfId, req.Item.VendorCode, err)
		return nil, grpc.Errorf(codes.Internal, "Could not insert item into part_manufacturer : %s", err)
	}
	s.LOG.Infof("Peer:%s INSERT INTO part (mnf_id,vendor_code) values (%s , %s)\n", p.Addr.String(), req.Item.MnfId, req.Item.VendorCode)
	return &prt.CreatePartResponse{Id: u.Id}, nil
}

// CreateParts create todo items from a list of todo descriptions
func (s Service) CreateParts(ctx context.Context, req *prt.CreatePartsRequest) (*prt.CreatePartsResponse, error) {
	var ids []string
	p, _ := peer.FromContext(ctx)
	for _, item := range req.Items {
		item.Id = uuid.NewV4().String()
		u := &prt.Part{}
		if err := s.DB.QueryRow(
			"INSERT INTO part (id, mnf_id, vendor_code, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id",
			item.Id,
			item.MnfId,
			item.VendorCode,
		).Scan(&u.Id); err != nil {
			s.LOG.Fatalf("Peer:%s ERROR INSERT part (mnf_id, vendor_code) values (%s, %s) err: %s", p.Addr.String(), item.MnfId, item.VendorCode, err)
			continue
		}
		ids = append(ids, item.Id)
		s.LOG.Infof("Peer:%s INSERT part (id,mnf_id, vendor_code) values (%s, %s, %s)", p.Addr.String(), item.Id, item.MnfId, item.VendorCode)
	}
	return &prt.CreatePartsResponse{Ids: ids}, nil
}

// GetPart retrieves a part item from its ID
func (s Service) GetPart(ctx context.Context, req *prt.GetPartRequest) (*prt.GetPartResponse, error) {
	var item prt.Part
	return &prt.GetPartResponse{Item: &item}, nil
}

// ListPart retrieves a part item from its ID
func (s Service) ListPart(ctx context.Context, req *prt.ListPartRequest) (*prt.ListPartResponse, error) {
	var items []*prt.Part
	return &prt.ListPartResponse{Items: items}, nil
}

// DeletePart deletes a part given an ID
func (s Service) DeletePart(ctx context.Context, req *prt.DeletePartRequest) (*prt.DeletePartResponse, error) {
	return &prt.DeletePartResponse{}, nil
}

// UpdatePart updates a part item
func (s Service) UpdatePart(ctx context.Context, req *prt.UpdatePartRequest) (*prt.UpdatePartResponse, error) {
	return &prt.UpdatePartResponse{}, nil
}

// UpdateTodos updates todo items given their respective title and description.
func (s Service) UpdateParts(ctx context.Context, req *prt.UpdatePartsRequest) (*prt.UpdatePartsResponse, error) {
	return &prt.UpdatePartsResponse{}, nil
}
