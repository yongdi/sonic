package service

import (
	"context"
	"mime/multipart"

	"sonic/consts"
	"sonic/model/dto"
	"sonic/model/entity"
	"sonic/model/param"
)

type AttachmentService interface {
	Page(ctx context.Context, queryParam *param.AttachmentQuery) ([]*entity.Attachment, int64, error)
	GetAttachment(ctx context.Context, attachmentID int32) (*entity.Attachment, error)
	Upload(ctx context.Context, fileHeader *multipart.FileHeader) (*dto.AttachmentDTO, error)
	Delete(ctx context.Context, attachmentID int32) (*entity.Attachment, error)
	DeleteBatch(ctx context.Context, ids []int32) ([]*entity.Attachment, error)
	Update(ctx context.Context, id int32, updateParam *param.AttachmentUpdate) (*entity.Attachment, error)
	GetAllMediaTypes(ctx context.Context) ([]string, error)
	GetAllTypes(ctx context.Context) ([]consts.AttachmentType, error)
	ConvertToDTO(ctx context.Context, attachment *entity.Attachment) (*dto.AttachmentDTO, error)
	ConvertToDTOs(ctx context.Context, attachments []*entity.Attachment) ([]*dto.AttachmentDTO, error)
}
