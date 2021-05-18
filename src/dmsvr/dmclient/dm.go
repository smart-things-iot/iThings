// Code generated by goctl. DO NOT EDIT!
// Source: dm.proto

//go:generate mockgen -destination ./dm_mock.go -package dmclient -source $GOFILE

package dmclient

import (
	"context"

	"gitee.com/godLei6/things/src/dmsvr/dm"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	LoginAuthReq  = dm.LoginAuthReq
	AccessAuthReq = dm.AccessAuthReq
	Response      = dm.Response

	Dm interface {
		LoginAuth(ctx context.Context, in *LoginAuthReq) (*Response, error)
		AccessAuth(ctx context.Context, in *AccessAuthReq) (*Response, error)
	}

	defaultDm struct {
		cli zrpc.Client
	}
)

func NewDm(cli zrpc.Client) Dm {
	return &defaultDm{
		cli: cli,
	}
}

func (m *defaultDm) LoginAuth(ctx context.Context, in *LoginAuthReq) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.LoginAuth(ctx, in)
}

func (m *defaultDm) AccessAuth(ctx context.Context, in *AccessAuthReq) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.AccessAuth(ctx, in)
}