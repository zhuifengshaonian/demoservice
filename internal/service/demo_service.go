package service

import (
	"context"
	v1 "demoservice/api/demoservice/v1"
	"demoservice/internal/pkg/zlog"
	"sync"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Sn struct {
	mu sync.Mutex
	SN int64
}

type DemoService struct {
	*v1.UnimplementedDemoServiceServer
}

var sn = Sn{SN: 0}

// Next 生成新的序列号 同步锁 不支持并发获取
func (sn *Sn) Next() int64 {
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.SN = sn.SN + 1
	return sn.SN
}

func NewDemoService() *DemoService {
	return &DemoService{}
}

func (d *DemoService) Talk(c context.Context, r *v1.TalkRequest) (reply *v1.TalkReply, err error) {
	zlog.Sugar.Infof("recive msg: %v", r)
	reply = &v1.TalkReply{
		Sn: sn.Next(),
		TalkMessage: &v1.TalkMessage{
			Message:  r.TalkMessage.ToName + " reply: " + r.TalkMessage.Message,
			FromName: r.TalkMessage.ToName,
			ToName:   r.TalkMessage.FromName,
			Time:     timestamppb.Now(),
		}}
	zlog.Sugar.Infof("reply msg: %v", r)
	return
}
