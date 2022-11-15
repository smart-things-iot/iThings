package deviceMsgEvent

import (
	"context"
	"encoding/json"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceMsg"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceMsg/msgGateway"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceMsg/msgHubLog"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceStatus"
	"github.com/i-Things/things/src/disvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"time"
)

type GatewayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	dreq   msgGateway.Msg
	topics []string
}

func NewGatewayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GatewayLogic {
	return &GatewayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *GatewayLogic) initMsg(msg *deviceMsg.PublishMsg) (err error) {
	err = utils.Unmarshal(msg.Payload, &l.dreq)
	if err != nil {
		return errors.Parameter.AddDetail("things topic is err:" + msg.Topic)
	}
	l.topics = strings.Split(msg.Topic, "/")
	if len(l.topics) < 5 || l.topics[1] != "up" {
		return errors.Parameter.AddDetail("initMsg topic is err:" + msg.Topic)
	}
	return nil
}

func (l *GatewayLogic) Handle(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	l.Infof("%s req=%+v", utils.FuncName(), msg)
	err = l.initMsg(msg)
	if err != nil {
		return nil, err
	}
	var (
		resp *msgGateway.Msg
	)

	switch l.topics[2] {
	case msgGateway.TypeOperation:
		resp, err = l.HandleOperation(msg)
	case msgGateway.TypeStatus:
		resp, err = l.HandleStatus(msg)
	}
	respStr, _ := json.Marshal(resp)
	return &deviceMsg.PublishMsg{
		Topic:      deviceMsg.GenRespTopic(msg.Topic),
		Payload:    respStr,
		Timestamp:  time.Now(),
		ProductID:  msg.ProductID,
		DeviceName: msg.DeviceName,
	}, nil
}

func (l *GatewayLogic) HandleOperation(msg *deviceMsg.PublishMsg) (respMsg *msgGateway.Msg, err error) {
	l.Debugf("%s", utils.FuncName())
	var resp = msgGateway.Msg{
		CommonMsg: deviceMsg.NewRespCommonMsg(l.dreq.Method, l.dreq.ClientToken),
	}
	resp.AddStatus(errors.OK)
	switch l.dreq.Method {
	case deviceMsg.Bind:
		_, err := l.svcCtx.DeviceM.DeviceGatewayMultiCreate(l.ctx, &dm.DeviceGatewayMultiCreateReq{
			GatewayProductID:  msg.ProductID,
			GatewayDeviceName: msg.DeviceName,
			List:              ToDmDevicesCore(l.dreq.Payload.Devices),
		})
		if err != nil {
			resp.AddStatus(err)
			return &resp, err
		}
		resp.Payload = &msgGateway.GatewayPayload{Devices: l.dreq.Payload.Devices}
	case deviceMsg.Unbind:
		_, err := l.svcCtx.DeviceM.DeviceGatewayMultiDelete(l.ctx, &dm.DeviceGatewayMultiDeleteReq{
			GatewayProductID:  msg.ProductID,
			GatewayDeviceName: msg.DeviceName,
			List:              ToDmDevicesCore(l.dreq.Payload.Devices),
		})
		if err != nil {
			resp.AddStatus(err)
			return &resp, err
		}
		resp.Payload = &msgGateway.GatewayPayload{Devices: l.dreq.Payload.Devices}
	case deviceMsg.DescribeSubDevices:
		deviceList, err := l.svcCtx.DeviceM.DeviceGatewayIndex(l.ctx, &dm.DeviceGatewayIndexReq{
			GatewayProductID:  msg.ProductID,
			GatewayDeviceName: msg.DeviceName,
		})
		if err != nil {
			resp.AddStatus(err)
			return &resp, err
		}
		var payload msgGateway.GatewayPayload
		for _, device := range deviceList.List {
			payload.Devices = append(payload.Devices, &msgGateway.Device{
				ProductID:  device.ProductID,
				DeviceName: device.DeviceName,
				Result:     errors.OK.Code,
			})
		}
		resp.Payload = &payload
	}
	return &resp, err
}

func (l *GatewayLogic) HandleStatus(msg *deviceMsg.PublishMsg) (respMsg *msgGateway.Msg, err error) {
	l.Debugf("%s", utils.FuncName())
	l.Debugf("%s", utils.FuncName())
	var resp = msgGateway.Msg{
		CommonMsg: deviceMsg.NewRespCommonMsg(l.dreq.Method, l.dreq.ClientToken),
	}
	resp.AddStatus(errors.OK)
	var (
		isOnline = int64(def.False)
		action   = deviceStatus.DisConnectStatus
		payload  msgGateway.GatewayPayload
	)

	switch l.dreq.Method {
	case deviceMsg.Online:
		isOnline = def.True
		action = deviceStatus.ConnectStatus
	case deviceMsg.Offline:
	default:
		err := errors.Parameter.AddDetailf("not support method :%s", l.dreq.Method)
		resp.AddStatus(err)
		return &resp, err
	}
	for _, v := range l.dreq.Payload.Devices {
		err = l.svcCtx.HubLogRepo.Insert(l.ctx, &msgHubLog.HubLog{
			ProductID:  v.ProductID,
			Action:     action,
			Timestamp:  msg.Timestamp, // 操作时间
			DeviceName: v.DeviceName,
			TranceID:   utils.TraceIdFromContext(l.ctx),
			ResultType: errors.Fmt(err).GetCode(),
		})
		if err != nil {
			l.Errorf("%s.LogRepo.insert productID:%v deviceName:%v err:%v",
				utils.FuncName(), v.ProductID, v.DeviceName, err)
		}
		//更新对应设备的online状态
		_, err := l.svcCtx.DeviceM.DeviceInfoUpdate(l.ctx, &dm.DeviceInfo{
			ProductID:  v.ProductID,
			DeviceName: v.DeviceName,
			IsOnline:   isOnline,
		})
		if err != nil {
			l.Errorf("%s.LogRepo.DeviceInfoUpdate productID:%v deviceName:%v err:%v",
				utils.FuncName(), v.ProductID, v.DeviceName, err)
		}
		payload.Devices = append(payload.Devices, &msgGateway.Device{
			ProductID:  v.ProductID,
			DeviceName: v.DeviceName,
			Result:     errors.Fmt(err).GetCode(),
		})
	}
	return &resp, err
}
