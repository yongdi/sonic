package impl

import (
	"sonic/injection"
	"sonic/service/file_storage"
)

func init() {
	injection.Provide(
		NewAdminService,
		NewAttachmentService,
		NewAuthenticateService,
		NewBackUpService,
		NewBaseCommentService,
		NewBasePostService,
		NewCategoryService,
		NewEmailService,
		NewInstallService,
		NewJournalService,
		NewLinkService,
		NewJournalCommentService,
		NewLogService,
		NewMenuService,
		NewMetaService,
		NewBaseMFAService,
		NewTwoFactorTOTPMFAService,
		NewOneTimeTokenService,
		NewOptionService,
		NewClientOptionService,
		NewPhotoService,
		NewPostService,
		NewPostCategoryService,
		NewPostCommentService,
		NewPostTagService,
		NewSheetService,
		NewSheetCommentService,
		NewStatisticService,
		NewTagService,
		NewThemeService,
		NewUserService,
		file_storage.NewFileStorageComposite,
	)
}
