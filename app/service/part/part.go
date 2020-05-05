package part

import (
	"context"

	"database/sql"
	_ "github.com/lib/pq"

	prt "github.com/pprisn/grpc_rest_mnf/api/mnf/v1"
	//	"github.com/pprisn/grpc_rest_mnf/app/store"
	uuid "github.com/satori/go.uuid"
	//	"google.golang.org/grpc"
	//	"google.golang.org/grpc/codes"
)

// Service is the service dealing with storing
// and retrieving part items from the database.

type Service struct {
	DB *sql.DB
}

// CreateTodo creates a todo given a description
func (s Service) CreatePart(ctx context.Context, req *prt.CreatePartRequest) (*prt.CreatePartResponse, error) {
	req.Item.Id = uuid.NewV4().String()
	//Проверить наличие MnfId в базе
	//	var mnfItem mnf.Mnf

	//	err := s.DB.Model(&mnfItem).Where("id = ?", req.Item.MnfId).First()
	//	if err != nil {
	//		return nil, grpc.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
	//	}
	//var v mnf.Part
	//	_, err := s.ST.Part().Create(req.Item)
	//err := s.DB.Insert(req.Item)
	//	if err != nil {
	//		return nil, grpc.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
	//	}
	return &prt.CreatePartResponse{Id: req.Item.Id}, nil
}

// CreateParts create todo items from a list of todo descriptions
func (s Service) CreateParts(ctx context.Context, req *prt.CreatePartsRequest) (*prt.CreatePartsResponse, error) {
	var ids []string
	for _, item := range req.Items {
		item.Id = uuid.NewV4().String()
		ids = append(ids, item.Id)
	}
	/*
		err := s.DB.Insert(&req.Items)
		if err != nil {
			return nil, grpc.Errorf(codes.Internal, "Could not insert items into the database: %s", err)
		}
	*/
	return &prt.CreatePartsResponse{Ids: ids}, nil
}

// GetPart retrieves a part item from its ID
func (s Service) GetPart(ctx context.Context, req *prt.GetPartRequest) (*prt.GetPartResponse, error) {
	var item prt.Part
	/*
		err := s.DB.Model(&item).Where("id = ?", req.Id).First()
		if err != nil {
			return nil, grpc.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", err)
		}
	*/
	return &prt.GetPartResponse{Item: &item}, nil
}

// ListPart retrieves a part item from its ID
func (s Service) ListPart(ctx context.Context, req *prt.ListPartRequest) (*prt.ListPartResponse, error) {
	var items []*prt.Part
	/*
		query := s.DB.Model(&items).Order("created_at ASC")
		if req.Limit > 0 {
			query.Limit(int(req.Limit))
		}
		if req.NotCompleted {
			query.Where("completed = false")
		}
		err := query.Select()
		if err != nil {
			return nil, grpc.Errorf(codes.NotFound, "Could not list items from the database: %s", err)
		}
	*/
	return &prt.ListPartResponse{Items: items}, nil
}

// DeletePart deletes a part given an ID
func (s Service) DeletePart(ctx context.Context, req *prt.DeletePartRequest) (*prt.DeletePartResponse, error) {
	/*
		err := s.DB.Delete(&todo.Todo{Id: req.Id})
		if err != nil {
			return nil, grpc.Errorf(codes.Internal, "Could not delete item from the database: %s", err)
		}
	*/
	return &prt.DeletePartResponse{}, nil
}

// UpdatePart updates a part item
func (s Service) UpdatePart(ctx context.Context, req *prt.UpdatePartRequest) (*prt.UpdatePartResponse, error) {
	/*
		req.Item.UpdatedAt = types.TimestampNow()
		res, err := s.DB.Model(req.Item).Column("title", "description", "completed", "updated_at").Update()
		if res.RowsAffected() == 0 {
			return nil, grpc.Errorf(codes.NotFound, "Could not update item: not found")
		}
		if err != nil {
			return nil, grpc.Errorf(codes.Internal, "Could not update item from the database: %s", err)
		}
	*/
	return &prt.UpdatePartResponse{}, nil
}

// UpdateTodos updates todo items given their respective title and description.
func (s Service) UpdateParts(ctx context.Context, req *prt.UpdatePartsRequest) (*prt.UpdatePartsResponse, error) {
	/*
		time := types.TimestampNow()
		for _, item := range req.Items {
			item.UpdatedAt = time
		}
		res, err := s.DB.Model(&req.Items).Column("title", "description", "completed", "updated_at").Update()
		if res.RowsAffected() == 0 {
			return nil, grpc.Errorf(codes.NotFound, "Could not update items: not found")
		}
		if err != nil {
			return nil, grpc.Errorf(codes.Internal, "Could not update items from the database: %s", err)
		}
	*/
	return &prt.UpdatePartsResponse{}, nil
}
