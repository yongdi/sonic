package content

import (
	"github.com/gin-gonic/gin"

	"sonic/handler/content/model"
	"sonic/template"
	"sonic/util"
)

type IndexHandler struct {
	PostModel *model.PostModel
}

func NewIndexHandler(postModel *model.PostModel) *IndexHandler {
	return &IndexHandler{
		PostModel: postModel,
	}
}

func (h *IndexHandler) Index(ctx *gin.Context, model template.Model) (string, error) {
	return h.PostModel.List(ctx, 0, model)
}

func (h *IndexHandler) IndexPage(ctx *gin.Context, model template.Model) (string, error) {
	page, err := util.ParamInt32(ctx, "page")
	if err != nil {
		return "", err
	}
	return h.PostModel.List(ctx, int(page)-1, model)
}
