package controller

import (
	"context"
	"newsfeed/pb"
	"newsfeed/service"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAttachmentController(service service.AttachmentUploaderServiceInterface) *AttachmentServer {
	return &AttachmentServer{Svc: service}
}

type AttachmentServer struct {
	pb.UnimplementedAttachmentServiceServer
	Svc service.AttachmentUploaderServiceInterface
}

func (a AttachmentServer) CreateMultiple(ctx context.Context, req *pb.RequestAttachments) (*pb.ResponseAttachments, error) {
	return a.Svc.UploadAttachments(req)
}

func (a AttachmentServer) FetchSingle(ctx context.Context, req *pb.FindOneRequest) (*pb.ResponseAttachment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchSingle not implemented")
}

func (a AttachmentServer) FetchAll(ctx context.Context, req *empty.Empty) (*pb.ResponseAttachments, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchAll not implemented")
}

func (a AttachmentServer) Delete(ctx context.Context, req *pb.AttachmentIDs) (*pb.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
