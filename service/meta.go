package service

import (
	"context"

	"sonic/model/dto"
	"sonic/model/entity"
)

type MetaService interface {
	GetPostsMeta(ctx context.Context, postIDs []int32) (map[int32][]*entity.Meta, error)
	GetPostMeta(ctx context.Context, postID int32) ([]*entity.Meta, error)
	ConvertToMetaDTO(meta *entity.Meta) *dto.Meta
	ConvertToMetaDTOs(metas []*entity.Meta) []*dto.Meta
}
