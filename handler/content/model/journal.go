package model

import (
	"context"

	"sonic/model/dto"
	"sonic/model/param"
	"sonic/model/property"
	"sonic/service"
	"sonic/template"
)

func NewJournalModel(optionService service.OptionService,
	themeService service.ThemeService,
	journalService service.JournalService,
	JournalService service.JournalService,
) *JournalModel {
	return &JournalModel{
		OptionService:  optionService,
		ThemeService:   themeService,
		JournalService: journalService,
	}
}

type JournalModel struct {
	JournalService service.JournalService
	OptionService  service.OptionService
	ThemeService   service.ThemeService
}

func (p *JournalModel) Journals(ctx context.Context, model template.Model, page int) (string, error) {
	pageSize := p.OptionService.GetOrByDefault(ctx, property.JournalPageSize).(int)
	journals, total, err := p.JournalService.Page(ctx,
		param.Page{
			PageNum:  page,
			PageSize: pageSize,
		},
		&param.Sort{
			Fields: []string{"createTime,desc"},
		})
	if err != nil {
		return "", err
	}
	journalDTOs, err := p.JournalService.ConvertToWithCommentDTOList(ctx, journals)
	if err != nil {
		return "", err
	}
	journalPage := dto.NewPage(journalDTOs, total, param.Page{PageNum: page, PageSize: pageSize})
	model["is_journals"] = true
	model["journals"] = journalPage
	model["meta_keywords"] = p.OptionService.GetOrByDefault(ctx, property.SeoKeywords)
	model["meta_description"] = p.OptionService.GetOrByDefault(ctx, property.SeoDescription)
	return p.ThemeService.Render(ctx, "journals")
}
