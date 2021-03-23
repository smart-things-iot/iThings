// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

//go:generate mockgen -destination ./user_mock.go -package userclient -source $GOFILE

package userclient

import (
	"context"

	"yl/src/user/rpc/user"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Register2Req     = user.Register2Req
	GetUserInfoReq   = user.GetUserInfoReq
	LoginResp        = user.LoginResp
	JwtToken         = user.JwtToken
	RegisterCoreReq  = user.RegisterCoreReq
	RegisterCoreResp = user.RegisterCoreResp
	LoginReq         = user.LoginReq
	CheckTokenReq    = user.CheckTokenReq
	CheckTokenResp   = user.CheckTokenResp
	UserInfo         = user.UserInfo
	Register2Resp    = user.Register2Resp
	GetUserInfoResp  = user.GetUserInfoResp

	User interface {
		Login(ctx context.Context, in *LoginReq) (*LoginResp, error)
		RegisterCore(ctx context.Context, in *RegisterCoreReq) (*RegisterCoreResp, error)
		Register2(ctx context.Context, in *Register2Req) (*Register2Resp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq) (*GetUserInfoResp, error)
		CheckToken(ctx context.Context, in *CheckTokenReq) (*CheckTokenResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq) (*LoginResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in)
}

func (m *defaultUser) RegisterCore(ctx context.Context, in *RegisterCoreReq) (*RegisterCoreResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.RegisterCore(ctx, in)
}

func (m *defaultUser) Register2(ctx context.Context, in *Register2Req) (*Register2Resp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register2(ctx, in)
}

func (m *defaultUser) GetUserInfo(ctx context.Context, in *GetUserInfoReq) (*GetUserInfoResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in)
}

func (m *defaultUser) CheckToken(ctx context.Context, in *CheckTokenReq) (*CheckTokenResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CheckToken(ctx, in)
}