package part

import (
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/lib/pq"
	prt "github.com/pprisn/grpc_rest_mnf/api/mnf/v1"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"time"
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
	p, _ := peer.FromContext(ctx)
	u := &prt.Part{}
	var tc time.Time
	tc = time.Now()
	//	td = time.Now()
	if err := s.DB.QueryRow(
		"SELECT id,mnf_id,vendor_code, created_at FROM part WHERE id=$1 and deleted_at IS NULL;", req.Id,
	).Scan(
		&u.Id,
		&u.MnfId,
		&u.VendorCode,
		&tc,
	); err != nil {
		s.LOG.Infof("Peer:%s ERROR SELECT id = %s  from  part : %s", p.Addr.String(), req.Id, err)
		return nil, grpc.Errorf(codes.NotFound, "Could not retrieve item from part: %s", err)
	}
	u.CreatedAt, _ = ptypes.TimestampProto(tc)
	//	u.DeletedAt, _ = ptypes.TimestampProto(td)
	s.LOG.Infof("Peer:%s SELECT id, mnf_id, vendor_code, created_at FROM part_manufacturer => %v\n", p.Addr.String(), u)
	return &prt.GetPartResponse{Item: u}, nil
}

// ListPart retrieves a part item from its ID
func (s Service) ListPart(ctx context.Context, req *prt.ListPartRequest) (*prt.ListPartResponse, error) {
	var items []*prt.Part
	p, _ := peer.FromContext(ctx)
	rows, err := s.DB.Query("select id, mnf_id, vendor_code, created_at FROM part WHERE mnf_id=$1 and deleted_at IS NULL", req.Mnfid)
	if err != nil {
		s.LOG.Infof("Peer:%s ERROR Query SELECT from  part : %s", p.Addr.String(), err)
		return nil, grpc.Errorf(codes.NotFound, "Could not make select from part: %s", err)
	}
	var tt time.Time
	tt = time.Now()
	for rows.Next() {
		u := &prt.Part{}
		err := rows.Scan(&u.Id, &u.MnfId, &u.VendorCode, &tt)
		if err != nil {
			// log
			continue
		}
		u.CreatedAt, _ = ptypes.TimestampProto(tt)
		items = append(items, u)
		s.LOG.Infof("Peer:%s SELECT id, mnf_id, vendor_code, created_at  FROM part => %v\n", p.Addr.String(), u)
	}
	return &prt.ListPartResponse{Items: items}, nil
}

// DeletePart deletes a part given an ID
func (s Service) DeletePart(ctx context.Context, req *prt.DeletePartRequest) (*prt.DeletePartResponse, error) {
	p, _ := peer.FromContext(ctx)
	sqlStmt := `                                                                                            		
DELETE FROM part                                                                               		
WHERE id = $1 and deleted_at IS NULL;`
	_, err := s.DB.Exec(sqlStmt, req.Id)
	if err != nil {
		s.LOG.Infof("Peer:%s ERROR DELETE FROM part WHERE id => %s\n", p.Addr.String(), req.Id)
		return nil, grpc.Errorf(codes.Internal, "Could not delete item from the part: %s", err)
	}
	s.LOG.Infof("Peer:%s DELETE FROM part WHERE id => %s\n", p.Addr.String(), req.Id)
	return &prt.DeletePartResponse{}, nil
}

// UpdatePart updates a part item
func (s Service) UpdatePart(ctx context.Context, req *prt.UpdatePartRequest) (*prt.UpdatePartResponse, error) {
	p, _ := peer.FromContext(ctx)
	sqlStmt := `                                                                                                              		
UPDATE part                                                                                                      		
SET vendor_code = $2                                                                                                                 		
WHERE id = $1 and deleted_at IS NULL;`
	res, err := s.DB.Exec(sqlStmt, req.Item.Id, req.Item.VendorCode)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "DB.Exec could not update part: %s", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "RowsAffected could not do UPDATE from part: %s", err)
	}
	if count == 0 {
		return nil, grpc.Errorf(codes.NotFound, "Could not update item: not found")
	}
	s.LOG.Infof("Peer:%s UPDATE part SET vendor_code = %s WHERE id = %s\n", p.Addr.String(), req.Item.VendorCode, req.Item.Id)
	return &prt.UpdatePartResponse{}, nil
}

// UpdateTodos updates todo items given their respective title and description.
func (s Service) UpdateParts(ctx context.Context, req *prt.UpdatePartsRequest) (*prt.UpdatePartsResponse, error) {
	p, _ := peer.FromContext(ctx)
	sqlStmt := `                                                                                                          		
UPDATE part                                                                                                  		
SET vendor_code = $2                                                                                                             		
WHERE id = $1 and deleted_at IS NULL;`
	for _, item := range req.Items {
		res, err := s.DB.Exec(sqlStmt, item.Id, item.VendorCode)
		if err != nil {
			return nil, grpc.Errorf(codes.Internal, "DB.Exec could not UPDATE part: %s", err)
		}
		count, err := res.RowsAffected()
		if err != nil {
			return nil, grpc.Errorf(codes.Internal, "RowsAffected could not do UPDATE part: %s", err)
		}
		if count == 0 {
			return nil, grpc.Errorf(codes.NotFound, "Could not update item: not found")
		}
		s.LOG.Infof("Peer:%s UPDATE part SET vendor_code = %s WHERE id = %s\n", p.Addr.String(), item.VendorCode, item.Id)
	}
	return &prt.UpdatePartsResponse{}, nil
}
