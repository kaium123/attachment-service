package repository

import (
	"fmt"
	"newsfeed/ent"
	"newsfeed/ent/attachment"

	"golang.org/x/net/context"
)

type AttachmentRepositoryInterface interface {
	Store(ctx context.Context, attachmentPath string, name string) (*ent.Attachment, error)
	GetByID(ctx context.Context, id int) (*ent.Attachment, error)
	GetByPath(ctx context.Context, path string) (*ent.Attachment, error)
}

type AttachmentRepository struct {
	Db *ent.Client
}

func NewAttachmentRepository(db *ent.Client) AttachmentRepositoryInterface {
	return &AttachmentRepository{db}
}

func (ar *AttachmentRepository) Store(ctx context.Context, attachmentPath string, name string) (*ent.Attachment, error) {
	userData, err := ar.Db.Attachment.Create().SetPath(attachmentPath).SetName(name).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating attachment: %w", err)
	}
	return userData, nil
}

func (ar *AttachmentRepository) GetByID(ctx context.Context, id int) (*ent.Attachment, error) {
	attachmentData, err := ar.Db.Attachment.Query().Where(attachment.ID(id)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating attachment: %w", err)
	}
	return attachmentData, nil
}

func (ar *AttachmentRepository) GetByPath(ctx context.Context, path string) (*ent.Attachment, error) {
	attachmentData, err := ar.Db.Attachment.Query().Where(attachment.Path(path)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating attachment: %w", err)
	}
	return attachmentData, nil
}
