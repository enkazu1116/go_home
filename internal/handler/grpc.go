package handler

import (
	"context"
	"time"

	"github.com/enkazu1116/go_home/internal/domain"
	"github.com/enkazu1116/go_home/internal/entity"
	pb "github.com/enkazu1116/go_home/internal/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// userGRPCServer implements pb.UserServiceServer
type userGRPCServer struct {
	pb.UnimplementedUserServiceServer
	usecase domain.UserUsecase
}

func NewUserGRPCServer(u domain.UserUsecase) pb.UserServiceServer {
	return &userGRPCServer{usecase: u}
}

// helpers: entity <-> pb
func toPB(u *entity.User) *pb.User {
	if u == nil {
		return nil
	}
	var deleted *timestamppb.Timestamp
	if u.DeletedAt.Valid {
		deleted = timestamppb.New(u.DeletedAt.Time)
	}
	return &pb.User{
		Id:        u.ID,
		AuthId:    u.AuthID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
		DeletedAt: deleted,
	}
}

func fromPB(p *pb.User) entity.User {
	var del entity.User
	// map common fields
	u := entity.User{
		ID:     p.GetId(),
		AuthID: p.GetAuthId(),
		Name:   p.GetName(),
		Email:  p.GetEmail(),
		Role:   p.GetRole(),
	}
	// created/updated may be nil
	if p.GetCreatedAt() != nil {
		u.CreatedAt = p.GetCreatedAt().AsTime()
	}
	if p.GetUpdatedAt() != nil {
		u.UpdatedAt = p.GetUpdatedAt().AsTime()
	}
	// deleted
	if p.GetDeletedAt() != nil {
		u.DeletedAt.Time = p.GetDeletedAt().AsTime()
		u.DeletedAt.Valid = true
	}
	return u
}

func (s *userGRPCServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := entity.User{
		ID:     uuid.NewString(),
		AuthID: req.GetAuthId(),
		Name:   req.GetName(),
		Email:  req.GetEmail(),
		Role:   req.GetRole(),
	}
	// set timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	if err := s.usecase.CreateUser(ctx, user); err != nil {
		return nil, status.Errorf(codes.Internal, "create user: %v", err)
	}
	return &pb.CreateUserResponse{User: toPB(&user)}, nil
}

func (s *userGRPCServer) ListUsers(ctx context.Context, _ *emptypb.Empty) (*pb.ListUsersResponse, error) {
	users, err := s.usecase.FindAllUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list users: %v", err)
	}
	out := make([]*pb.User, 0, len(users))
	for i := range users {
		out = append(out, toPB(&users[i]))
	}
	return &pb.ListUsersResponse{Users: out}, nil
}

func (s *userGRPCServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	u, err := s.usecase.FindFirst(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%v", err)
	}
	return &pb.GetUserResponse{User: toPB(u)}, nil
}

func (s *userGRPCServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	if req.GetUser() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "user required")
	}
	ent := fromPB(req.GetUser())
	// update timestamp
	ent.UpdatedAt = time.Now()
	if err := s.usecase.UpdateUser(ctx, ent); err != nil {
		return nil, status.Errorf(codes.Internal, "update user: %v", err)
	}
	return &pb.UpdateUserResponse{User: toPB(&ent)}, nil
}

func (s *userGRPCServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	if req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id required")
	}
	ent := entity.User{ID: req.GetId()}
	if err := s.usecase.DeleteUser(ctx, ent); err != nil {
		return nil, status.Errorf(codes.Internal, "delete user: %v", err)
	}
	return &emptypb.Empty{}, nil
}
