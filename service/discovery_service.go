package service

import (
	context "context"

	"github.com/MohammadTaheri304/discovery/database"
	"github.com/MohammadTaheri304/discovery/rpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type DiscoveryService struct {
	database *database.Database
}

func NewDiscoveryService(db *database.Database) *DiscoveryService {
	return &DiscoveryService{
		database: db,
	}
}

func (m *DiscoveryService) Register(ctx context.Context, req *rpc.Message) (*rpc.Message, error) {
	//TODO Ensure validity before save it!
	m.database.Set(req.GetKey(), req.GetValue())
	return req, nil
}

func (m *DiscoveryService) Get(ctx context.Context, req *rpc.Message) (*rpc.Message, error) {
	val, ok := m.database.Get(req.Key)
	if !ok {
		return nil, status.Error(codes.NotFound, "key.not.found")
	}
	req.Value = val
	return req, nil
}
