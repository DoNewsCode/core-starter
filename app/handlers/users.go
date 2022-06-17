package handlers

import (
	context "context"
	"github.com/DoNewsCode/core-starter/proto/users"
)

type Users struct {
	users.UnimplementedUsersServer
}

func NewUsers() *Users {
	return &Users{
		UnimplementedUsersServer: users.UnimplementedUsersServer{},
	}
}

func (u *Users) Login(ctx context.Context, req *users.LoginReq) (*users.Rep, error) {
	return &users.Rep{
		Code: 0,
		Msg:  "ok",
	}, nil
}

func (u *Users) GetCode(ctx context.Context, req *users.GetCodeReq) (*users.Rep, error) {
	return &users.Rep{
		Code: 0,
		Msg:  "ok",
	}, nil
}
