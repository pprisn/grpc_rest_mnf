package part

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	prt "github.com/pprisn/grpc_rest_mnf/api/mnf/v1"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type Service struct {
	DB  *sql.DB
	LOG *logrus.Logger
}

// CreatePart creates
func (s Service) CreatePart(ctx context.Context, req *prt.CreatePartRequest) (*prt.CreatePartResponse, error) {
	req.Item.Id = uuid.NewV4().String()
	return &prt.CreatePartResponse{Id: req.Item.Id}, nil
}

// CreateParts create todo items from a list of todo descriptions
func (s Service) CreateParts(ctx context.Context, req *prt.CreatePartsRequest) (*prt.CreatePartsResponse, error) {
	var ids []string
	for _, item := range req.Items {
		item.Id = uuid.NewV4().String()
		ids = append(ids, item.Id)
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
