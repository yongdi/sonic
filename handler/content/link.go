package content

import (
	"github.com/gin-gonic/gin"

	"sonic/handler/content/model"
	"sonic/template"
)

type LinkHandler struct {
	LinkModel *model.LinkModel
}

func NewLinkHandler(
	linkModel *model.LinkModel,
) *LinkHandler {
	return &LinkHandler{
		LinkModel: linkModel,
	}
}

func (t *LinkHandler) Link(ctx *gin.Context, model template.Model) (string, error) {
	return t.LinkModel.Links(ctx, model)
}
