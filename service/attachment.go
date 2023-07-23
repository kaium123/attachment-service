package service

import (
	"context"
	"newsfeed/common/logger"
	"newsfeed/common/utils"
	"newsfeed/ent"
	fileUploader "newsfeed/modules/attachment/file_uploader"
	"newsfeed/pb"

	"github.com/gin-gonic/gin"
)

type AttachmentUploaderServiceInterface interface {
	UploadAttachments(attachments *pb.RequestAttachments) (*pb.ResponseAttachments, error)
	GetSingleAttachment(context *gin.Context, attachmentPath string)
	Delete(attachmentPath string) error
}

type AttachmentUploaderService struct {
	uploader  fileUploader.FileUploaderInterface
	entCLient *ent.Client
}

func NewAttachmentService(entClient *ent.Client, uploader fileUploader.FileUploaderInterface) AttachmentUploaderServiceInterface {
	service := &AttachmentUploaderService{
		uploader:  uploader,
		entCLient: entClient,
	}
	return service
}

func (aus AttachmentUploaderService) Delete(attachmentPath string) error {
	return nil
}

func (aus AttachmentUploaderService) GetSingleAttachment(context *gin.Context, attachmentPath string) {

}

func (aus AttachmentUploaderService) UploadAttachments(attachments *pb.RequestAttachments) (*pb.ResponseAttachments, error) {
	requestAttachments := &pb.ResponseAttachments{}
	logger.LogError(attachments)

	createdAttachments := []*ent.AttachmentCreate{}
	for _, attach := range attachments.Attachments {
		attachment := aus.entCLient.Attachment.Create().
			SetName(attach.Name).
			SetPath(attach.Path)
		createdAttachments = append(createdAttachments, attachment)
	}

	resp, err := aus.entCLient.Attachment.CreateBulk(createdAttachments...).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	responseAttachments := &pb.ResponseAttachments{}
	err = utils.CopyStructToStruct(resp, requestAttachments.Attachments)
	if err != nil {
		return nil, err
	}

	return responseAttachments, nil
}
