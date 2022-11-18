package service

import (
	"context"

	"sonic/model/param"
)

type InstallService interface {
	InstallBlog(ctx context.Context, installParam param.Install) error
}
