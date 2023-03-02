// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/logic/log"
	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"
)

type LogServer struct {
	svcCtx *svc.ServiceContext
	sys.UnimplementedLogServer
}

func NewLogServer(svcCtx *svc.ServiceContext) *LogServer {
	return &LogServer{
		svcCtx: svcCtx,
	}
}

func (s *LogServer) LoginLogIndex(ctx context.Context, in *sys.LoginLogIndexReq) (*sys.LoginLogIndexResp, error) {
	l := loglogic.NewLoginLogIndexLogic(ctx, s.svcCtx)
	return l.LoginLogIndex(in)
}

func (s *LogServer) OperLogIndex(ctx context.Context, in *sys.OperLogIndexReq) (*sys.OperLogIndexResp, error) {
	l := loglogic.NewOperLogIndexLogic(ctx, s.svcCtx)
	return l.OperLogIndex(in)
}

func (s *LogServer) LoginLogCreate(ctx context.Context, in *sys.LoginLogCreateReq) (*sys.Response, error) {
	l := loglogic.NewLoginLogCreateLogic(ctx, s.svcCtx)
	return l.LoginLogCreate(in)
}

func (s *LogServer) OperLogCreate(ctx context.Context, in *sys.OperLogCreateReq) (*sys.Response, error) {
	l := loglogic.NewOperLogCreateLogic(ctx, s.svcCtx)
	return l.OperLogCreate(in)
}