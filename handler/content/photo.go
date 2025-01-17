package content

import (
	"github.com/gin-gonic/gin"

	"sonic/handler/content/model"
	"sonic/service"
	"sonic/template"
	"sonic/util"
)

type PhotoHandler struct {
	OptionService service.OptionService
	PhotoService  service.PhotoService
	PhotoModel    *model.PhotoModel
}

func NewPhotoHandler(
	optionService service.OptionService,
	photoService service.PhotoService,
	photoModel *model.PhotoModel,

) *PhotoHandler {
	return &PhotoHandler{
		OptionService: optionService,
		PhotoService:  photoService,
		PhotoModel:    photoModel,
	}
}

func (p *PhotoHandler) PhotosPage(ctx *gin.Context, model template.Model) (string, error) {
	page, err := util.ParamInt32(ctx, "page")
	if err != nil {
		return "", err
	}
	return p.PhotoModel.Photos(ctx, model, int(page-1))
}

func (p *PhotoHandler) Phtotos(ctx *gin.Context, model template.Model) (string, error) {
	return p.PhotoModel.Photos(ctx, model, 0)
}
