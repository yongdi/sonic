package content

import (
	"github.com/gin-gonic/gin"

	"sonic/handler/content/model"
	"sonic/service"
	"sonic/template"
	"sonic/util"
)

type SheetHandler struct {
	OptionService service.OptionService
	SheetService  service.SheetService
	SheetModel    *model.SheetModel
}

func NewSheetHandler(
	optionService service.OptionService,
	sheetService service.SheetService,
	sheetModel *model.SheetModel,
) *SheetHandler {
	return &SheetHandler{
		OptionService: optionService,
		SheetService:  sheetService,
		SheetModel:    sheetModel,
	}
}

func (s *SheetHandler) SheetBySlug(ctx *gin.Context, model template.Model) (string, error) {
	slug, err := util.ParamString(ctx, "slug")
	if err != nil {
		return "", err
	}
	sheet, err := s.SheetService.GetBySlug(ctx, slug)
	if err != nil {
		return "", err
	}
	token, _ := ctx.Cookie("authentication")
	return s.SheetModel.Content(ctx, sheet, token, model)
}
