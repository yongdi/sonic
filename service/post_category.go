package service

import (
	"context"

	"sonic/consts"
	"sonic/model/entity"
)

type PostCategoryService interface {
	ListByCategoryID(ctx context.Context, categoryID int32, status consts.PostStatus) ([]*entity.Post, error)
	ListByPostIDs(ctx context.Context, postIDs []int32) ([]*entity.PostCategory, error)
	ListCategoryMapByPostID(ctx context.Context, postIDs []int32) (map[int32][]*entity.Category, error)
	ListCategoryByPostID(ctx context.Context, postID int32) ([]*entity.Category, error)
}
