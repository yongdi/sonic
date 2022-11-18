package content

import (
	"github.com/gin-gonic/gin"

	"sonic/handler/content/model"
	"sonic/service"
	"sonic/template"
	"sonic/util"
)

type JournalHandler struct {
	OptionService  service.OptionService
	JournalService service.JournalService
	JournalModel   *model.JournalModel
}

func NewJournalHandler(
	optionService service.OptionService,
	journalService service.JournalService,
	journalModel *model.JournalModel,

) *JournalHandler {
	return &JournalHandler{
		OptionService:  optionService,
		JournalService: journalService,
		JournalModel:   journalModel,
	}
}

func (p *JournalHandler) JournalsPage(ctx *gin.Context, model template.Model) (string, error) {
	page, err := util.ParamInt32(ctx, "page")
	if err != nil {
		return "", err
	}
	return p.JournalModel.Journals(ctx, model, int(page-1))
}

func (p *JournalHandler) Journals(ctx *gin.Context, model template.Model) (string, error) {
	return p.JournalModel.Journals(ctx, model, 0)
}
