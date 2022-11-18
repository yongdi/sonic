package service

import (
	"context"

	"sonic/consts"
	"sonic/model/entity"
	"sonic/model/param"
)

type JournalCommentService interface {
	BaseCommentService
	CountByStatusAndJournalID(ctx context.Context, status consts.CommentStatus, journalIDs []int32) (map[int32]int64, error)
	UpdateBy(ctx context.Context, commentID int32, commentParam *param.Comment) (*entity.Comment, error)
	CountByStatus(ctx context.Context, status consts.CommentStatus) (int64, error)
}
