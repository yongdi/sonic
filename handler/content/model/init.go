package model

import (
	"sonic/injection"
)

func init() {
	injection.Provide(NewPostModel)
	injection.Provide(NewCategoryModel)
	injection.Provide(NewSheetModel)
	injection.Provide(NewTagModel)
	injection.Provide(NewLinkModel)
	injection.Provide(NewPhotoModel)
	injection.Provide(NewJournalModel)
}
