package mnf

import (
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/lib/pq"
	api "github.com/pprisn/grpc_rest_mnf/api/mnf/v1"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"time"
	//	"fmt"
)

//Service is the service dealing with storing
//and retrieving  items from the database.

// Service ...
type Service struct {
	DB  *sql.DB
	LOG *logrus.Logger
}

// CreateMnf creates a manufacturer given a description
func (s Service) CreateMnf(ctx context.Context, req *api.CreateMnfRequest) (*api.CreateMnfResponse, error) {
	p, _ := peer.FromContext(ctx)
	req.Item.Id = uuid.NewV4().String()
	u := &api.Mnf{}
	if err := s.DB.QueryRow(
		"INSERT INTO part_manufacturer (id, name, created_at) VALUES ($1, $2, NOW()) RETURNING id",
		req.Item.Id,
		req.Item.Name,
	).Scan(&u.Id); err != nil {
		s.LOG.Fatalf("Peer:%s ERROR INSERT part_manufacturer (name) values (%s) err: %s", p.Addr.String(), req.Item.Name, err)
		return nil, grpc.Errorf(codes.Internal, "Could not insert item into part_manufacturer : %s", err)
	}
	return &api.CreateMnfResponse{Id: u.Id}, nil
}

// CreateParts create Mnf items from a list of Mnf descriptions
func (s Service) CreateMnfs(ctx context.Context, req *api.CreateMnfsRequest) (*api.CreateMnfsResponse, error) {
	var ids []string
	p, _ := peer.FromContext(ctx)
	for _, item := range req.Items {
		item.Id = uuid.NewV4().String()
		u := &api.Mnf{}
		if err := s.DB.QueryRow(
			"INSERT INTO part_manufacturer (id, name, created_at) VALUES ($1, $2, NOW()) RETURNING id",
			item.Id,
			item.Name,
		).Scan(&u.Id); err != nil {
			s.LOG.Infof("%s ERROR INSERT part_manufacturer ( name) values ( %s )  err:%s", p.Addr.String(), item.Name, err)
			//return nil, grpc.Errorf(codes.Internal, "Could not insert name =\"%s\" into part_manufacturer: %s", item.Name, err)
			continue
		}
		ids = append(ids, item.Id)
		s.LOG.Infof("Peer:%s INSERT (id,name) values (%s , %s)\n", p.Addr.String(), item.Id, item.Name)
	}

	return &api.CreateMnfsResponse{Ids: ids}, nil
}

// GetMnf retrieves a part item from its ID
func (s Service) GetMnf(ctx context.Context, req *api.GetMnfRequest) (*api.GetMnfResponse, error) {
	p, _ := peer.FromContext(ctx)
	u := &api.Mnf{}
	var tt time.Time
	tt = time.Now()
	if err := s.DB.QueryRow(
		"SELECT id, name, created_at FROM part_manufacturer WHERE id=$1", req.Id,
	).Scan(
		&u.Id,
		&u.Name,
		&tt,
	); err != nil {
		s.LOG.Infof("Peer:%s ERROR Could not retrieve  id = %s  from  part_manufacturer : %s", p.Addr.String(), req.Id, err)
		return nil, grpc.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", err)
	}
	u.CreatedAt, _ = ptypes.TimestampProto(tt)

	s.LOG.Infof("Peer:%s SELECT id, name FROM part_manufacturer => %v\n", p.Addr.String(), u)
	return &api.GetMnfResponse{Item: u}, nil
}

// ListMnf item from its ID
func (s Service) ListMnf(ctx context.Context, req *api.ListMnfRequest) (*api.ListMnfResponse, error) {
	var items []*api.Mnf
	p, _ := peer.FromContext(ctx)
	rows, err := s.DB.Query("select id, name, created_at from part_manufacturer LIMIT $1", req.Limit)
	if err != nil {
		return nil, grpc.Errorf(codes.NotFound, "Could not make select from part_manufacturr : %s", err)
	}
	var tt time.Time
	tt = time.Now()
	for rows.Next() {
		u := &api.Mnf{}
		err := rows.Scan(&u.Id, &u.Name, &tt)
		if err != nil {
			// log
			continue
		}
		u.CreatedAt, _ = ptypes.TimestampProto(tt)
		items = append(items, u)
		s.LOG.Infof("Peer:%s SELECT id, name, created_at  FROM part_manufacturer => %v\n", p.Addr.String(), u)
	}
	return &api.ListMnfResponse{Items: items}, nil
}

// DeleteMnf deletes a part given an ID
func (s Service) DeleteMnf(ctx context.Context, req *api.DeleteMnfRequest) (*api.DeleteMnfResponse, error) {
	p, _ := peer.FromContext(ctx)
	sqlStmt := `
DELETE FROM part_manufacturer
WHERE id = $1;`
	_, err := s.DB.Exec(sqlStmt, req.Id)
	if err != nil {
		s.LOG.Infof("Peer:%s ERROR DELETE FROM part_manufacturer WHERE id => %s\n", p.Addr.String(), req.Id)
		return nil, grpc.Errorf(codes.Internal, "Could not delete item from the database: %s", err)
	}
	s.LOG.Infof("Peer:%s DELETE FROM part_manufacturer WHERE id => %s\n", p.Addr.String(), req.Id)
	return &api.DeleteMnfResponse{}, nil
}

// UpdateMnf updates a manufactur item
func (s Service) UpdateMnf(ctx context.Context, req *api.UpdateMnfRequest) (*api.UpdateMnfResponse, error) {
	p, _ := peer.FromContext(ctx)
	//u := &api.Mnf{}
	//sqlStmt := `
	//UPDATE part_manufacturer
	//SET name = $2
	//WHERE id = $1
	//RETURNING id, name;`

	sqlStmt := `
UPDATE part_manufacturer
SET name = $2
WHERE id = $1;`
	//	QueryRow(sqlStmt, req.Item.Id, req.Item.Name).Scan(&u.Id, &u.Name)
	res, err := s.DB.Exec(sqlStmt, req.Item.Id, req.Item.Name)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "DB.Exec could not update item from the database: %s", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "RowsAffected could not do UPDATE from database: %s", err)
	}
	if count == 0 {
		return nil, grpc.Errorf(codes.NotFound, "Could not update item: not found")
	}
	s.LOG.Infof("Peer:%s UPDATE part_manufacturer SET name = %s WHERE id = %s\n", p.Addr.String(), req.Item.Name, req.Item.Id)
	return &api.UpdateMnfResponse{}, nil
}

// UpdateMnf updates part_manufactur items.
func (s Service) UpdateMnfs(ctx context.Context, req *api.UpdateMnfsRequest) (*api.UpdateMnfsResponse, error) {
	p, _ := peer.FromContext(ctx)
	sqlStmt := `
UPDATE part_manufacturer
SET name = $2
WHERE id = $1;`

	for _, item := range req.Items {
		res, err := s.DB.Exec(sqlStmt, item.Id, item.Name)
		if err != nil {
			return nil, grpc.Errorf(codes.Internal, "DB.Exec could not update item from the database: %s", err)
		}
		count, err := res.RowsAffected()
		if err != nil {
			return nil, grpc.Errorf(codes.Internal, "RowsAffected could not do UPDATE from database: %s", err)
		}
		if count == 0 {
			return nil, grpc.Errorf(codes.NotFound, "Could not update item: not found")
		}
		s.LOG.Infof("Peer:%s UPDATE part_manufacturer SET name = %s WHERE id = %s\n", p.Addr.String(), item.Name, item.Id)
	}

	return &api.UpdateMnfsResponse{}, nil
}
