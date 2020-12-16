package service

import (
	"context"
	"errors"
	"homework/internal/dao"

	"github.com/google/wire"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "homework/api"
)

type Service struct {
	dao dao.Dao
}

var Provider = wire.NewSet(NewService, dao.Provider)

func NewService(d dao.Dao) *Service {
	return &Service{dao: d}
}

func (s *Service) GetArticle(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.dao.GetUser(ctx, int(req.Id))
	if err != nil {
		if errors.Is(err, dao.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "Object Not Found")
		}
		return nil, status.Errorf(codes.Internal, "Error:%v", err)
	}
	return &pb.GetUserResponse{Name: user.Name}, nil
}


