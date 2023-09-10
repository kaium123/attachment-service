package service

import (
	"attachment/common/logger"
	"attachment/common/utils"
	fileUploader "attachment/file_uploader"
	"attachment/models"
	"attachment/pb"
	"context"
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

type AttachmentUploaderServiceInterface interface {
	UploadAttachments(attachments *pb.RequestAttachments) (*pb.ResponseAttachments, error)
	GetSingleAttachment(context *gin.Context, attachmentPath string)
	Delete(req *pb.AttachmentIDs) (*pb.DeleteResponse, error)
	FetchAll(sourceType string, sourceID int64) (*pb.ResponseAttachments, error)
}

type AttachmentUploaderService struct {
	uploader fileUploader.FileUploaderInterface
	Db       *sql.DB
}

func NewAttachmentService(Db *sql.DB, uploader fileUploader.FileUploaderInterface) AttachmentUploaderServiceInterface {
	service := &AttachmentUploaderService{
		uploader: uploader,
		Db:       Db,
	}
	return service
}

func (aus AttachmentUploaderService) Delete(req *pb.AttachmentIDs) (*pb.DeleteResponse, error) {
	query := "DELETE FROM attachments WHERE sourceId = $1 AND sourceType = $2"
	logger.LogError(req.SourceId, " ", req.SourceType)

	result, err := aus.Db.Exec(query, req.SourceId, req.SourceType)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	rows, err := result.RowsAffected()
	logger.LogError(rows)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{}, nil
}

func (aus AttachmentUploaderService) GetSingleAttachment(context *gin.Context, attachmentPath string) {

}

func (aus AttachmentUploaderService) UploadAttachments(attachments *pb.RequestAttachments) (*pb.ResponseAttachments, error) {
	requestAttachments := &pb.ResponseAttachments{}
	logger.LogError(attachments)

	tx, err := aus.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	for _, attach := range attachments.Attachments {
		sql := "INSERT INTO attachments (name, path,sourceType,sourceId) VALUES ($1, $2, $3, $4)"
		_, err = tx.ExecContext(context.Background(), sql, attach.Name, attach.Path, attach.SourceType, attach.SourceId)
		if err != nil {
			logger.LogError(err)
			tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return requestAttachments, nil
}

func getLastInsertedIDs(tx *sql.Tx) ([]int64, error) {
	return nil, errors.New("Not implemented")
}

func (aus AttachmentUploaderService) FetchAll(sourceType string, sourceID int64) (*pb.ResponseAttachments, error) {
	// SQL query with placeholders
	query := `
	 SELECT name, path 
	 FROM attachments
	 WHERE sourceId = $1 AND sourceType = $2
 `
	rows, err := aus.Db.QueryContext(context.Background(), query, sourceID, sourceType)
	if err != nil {
		logger.LogError(err)
		return nil, err

	}
	defer rows.Close()
	var attachments []models.Attachment
	for rows.Next() {
		var attachment models.Attachment
		if err := rows.Scan(&attachment.Name, &attachment.Path); err != nil {
			logger.LogError(err)
			return nil, err
		}
		attachments = append(attachments, attachment)
	}
	if err := rows.Err(); err != nil {
		logger.LogError(err)
		return nil, err
	}
	pbResponse := &pb.ResponseAttachments{}
	for _, attachment := range attachments {
		tmpAttachment := &pb.ResponseAttachment{}
		utils.CopyStructToStruct(attachment, tmpAttachment)
		tmpAttachment.SourceId = uint64(sourceID)
		pbResponse.Attachments = append(pbResponse.Attachments, tmpAttachment)
	}

	return pbResponse, nil
}
