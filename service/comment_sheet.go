package service

import (
	"context"

	"sonic/consts"
)

type SheetCommentService interface {
	BaseCommentService
	CountByStatus(ctx context.Context, status consts.CommentStatus) (int64, error)
}
