package grpc

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	userv1 "github.com/kirigaikabuto/bike-main-api/gen/proto"

	"github.com/kirigaikabuto/bike-main-api/internal/db"

	_ "github.com/jackc/pgx/v5/pgxpool"
)

type UserServer struct {
	userv1.UnimplementedUserServiceServer
	Queries *db.Queries
}

func (s *UserServer) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	user, err := s.Queries.CreateUser(ctx, db.CreateUserParams{
		Name: req.GetName(),
		Email: pgtype.Text{
			String: req.GetEmail(),
			Valid:  true,
		},
	})
	if err != nil {
		return nil, err
	}
	return &userv1.CreateUserResponse{
		User: &userv1.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email.String,
		},
	}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	user, err := s.Queries.GetUserByID(ctx, int32(req.Id))
	if err != nil {
		return nil, err
	}
	return &userv1.GetUserResponse{
		User: &userv1.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email.String,
		},
	}, nil
}
