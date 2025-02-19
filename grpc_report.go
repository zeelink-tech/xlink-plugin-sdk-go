package plugin

import (
	"context"
	"github.com/zeelink-tech/xlink-plugin-sdk-go/proto"
)

var _ Report = &gRPCReportClient{}

type gRPCReportClient struct {
	client proto.ReportClient
}

func (m *gRPCReportClient) Post(req *Request) (*Response, error) {
	res, err := m.client.Post(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return &Response{Data: res.Data}, nil
}

func (m *gRPCReportClient) State(req *Request) (*Response, error) {
	res, err := m.client.State(context.Background(), &proto.RequestArgs{
		Request: req.Req,
	})
	if err != nil {
		return nil, err
	}
	return &Response{Data: res.Data}, nil
}

type gRPCReportServer struct {
	proto.UnimplementedReportServer
	Impl Report
}

func (m *gRPCReportServer) Post(_ context.Context, req *proto.RequestArgs) (*proto.ResponseResult, error) {
	res, err := m.Impl.Post(&Request{
		Req: req.Request,
	})
	if err != nil {
		return &proto.ResponseResult{}, err
	}
	return &proto.ResponseResult{Data: res.Data}, nil
}

func (m *gRPCReportServer) State(_ context.Context, req *proto.RequestArgs) (resp *proto.ResponseResult, err error) {
	res, err := m.Impl.State(&Request{
		Req: req.Request,
	})
	if err != nil {
		return &proto.ResponseResult{}, err
	}
	return &proto.ResponseResult{Data: res.Data}, nil
}
