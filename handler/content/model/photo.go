package model

import (
	"context"

	"sonic/model/dto"
	"sonic/model/param"
	"sonic/model/property"
	"sonic/service"
	"sonic/template"
)

func NewPhotoModel(optionService service.OptionService,
	themeService service.ThemeService,
	photoService service.PhotoService,
	PhotoService service.PhotoService,
) *PhotoModel {
	return &PhotoModel{
		OptionService: optionService,
		ThemeService:  themeService,
		PhotoService:  photoService,
	}
}

type PhotoModel struct {
	PhotoService  service.PhotoService
	OptionService service.OptionService
	ThemeService  service.ThemeService
}

func (p *PhotoModel) Photos(ctx context.Context, model template.Model, page int) (string, error) {
	pageSize := p.OptionService.GetOrByDefault(ctx, property.PhotoPageSize).(int)
	photos, total, err := p.PhotoService.Page(ctx,
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
	photoDTOs := p.PhotoService.ConvertToDTOs(ctx, photos)
	photoPage := dto.NewPage(photoDTOs, total, param.Page{PageNum: page, PageSize: pageSize})
	model["is_photos"] = true
	model["photos"] = photoPage
	model["meta_keywords"] = p.OptionService.GetOrByDefault(ctx, property.SeoKeywords)
	model["meta_description"] = p.OptionService.GetOrByDefault(ctx, property.SeoDescription)
	return p.ThemeService.Render(ctx, "photos")
}
