package service

import (
	"context"

	"sonic/model/entity"
)

type AuthenticateService interface {
	PostAuthenticate(ctx context.Context, post *entity.Post, password string) (bool, error)
	CategoryAuthenticate(ctx context.Context, categoryID int32, password string) (bool, error)
}
