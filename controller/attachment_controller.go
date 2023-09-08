package controller

import (
	"attachment/pb"
	"attachment/service"
	"context"

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

func (a AttachmentServer) FetchSingle(ctx context.Context, req *pb.FindOneRequestParams) (*pb.ResponseAttachment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchSingle not implemented")
}

func (a AttachmentServer) FetchAll(ctx context.Context, req *pb.FindAllRequestParams) (*pb.ResponseAttachments, error) {
	return a.Svc.FetchAll(req.SourceType, req.SourceId)
}

func (a AttachmentServer) Delete(ctx context.Context, req *pb.AttachmentIDs) (*pb.DeleteResponse, error) {
	return a.Svc.Delete(req)
}
