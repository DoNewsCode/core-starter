package handlers

import (
	"context"
	pb "github.com/DoNewsCode/core-starter/app/proto"
	"github.com/DoNewsCode/core-starter/app/repositories"
)

func NewUserHandler(repository repositories.UserRepository) *UserHandler {
	return &UserHandler{
		repository: repository,
	}
}

type UserHandler struct {
	repository repositories.UserRepository
}

func (h *UserHandler) GetOne(ctx context.Context, request *pb.GetOneUserRequest) (*pb.UserInfoReply, error) {
	user := h.repository.GetOne(request.GetId())

	rep := &pb.UserInfoReply{
		Id:        int64(user.ID),
		Name:      user.Username,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return rep, nil
}
