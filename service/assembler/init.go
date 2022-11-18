package assembler

import "sonic/injection"

func init() {
	injection.Provide(
		NewBasePostAssembler,
		NewPostAssembler,
		NewSheetAssembler,
		NewBaseCommentAssembler,
		NewPostCommentAssembler,
		NewJournalCommentAssembler,
		NewSheetCommentAssembler,
	)
}
