package extension

import (
	"context"

	"sonic/consts"
	"sonic/model/param"
	"sonic/model/vo"
	"sonic/service"
	"sonic/service/assembler"
	"sonic/template"
)

type commentExtension struct {
	PostCommentService   service.PostCommentService
	PostCommentAssembler assembler.PostCommentAssembler
	Template             *template.Template
}

func RegisterCommentFunc(template *template.Template, postCommentService service.PostCommentService, postCommentAssembler assembler.PostCommentAssembler) {
	ce := commentExtension{
		PostCommentService:   postCommentService,
		Template:             template,
		PostCommentAssembler: postCommentAssembler,
	}
	ce.addGetLatestComment()
	ce.addGetCommentCount()
}

func (ce *commentExtension) addGetLatestComment() {
	getLatestComment := func(top int) ([]*vo.PostCommentWithPost, error) {
		commentQuery := param.CommentQuery{
			Sort:          &param.Sort{Fields: []string{"createTime,desc"}},
			Page:          param.Page{PageNum: 0, PageSize: top},
			CommentStatus: consts.CommentStatusPublished.Ptr(),
		}
		comments, _, err := ce.PostCommentService.Page(context.Background(), commentQuery, consts.CommentTypePost)
		if err != nil {
			return nil, err
		}
		return ce.PostCommentAssembler.ConvertToWithPost(context.Background(), comments)
	}
	ce.Template.AddFunc("getLatestComment", getLatestComment)
}

func (ce *commentExtension) addGetCommentCount() {
	getCommentCount := func() (int64, error) {
		count, err := ce.PostCommentService.CountByStatus(context.Background(), consts.CommentStatusPublished)
		return count, err
	}
	ce.Template.AddFunc("getCommentCount", getCommentCount)
}
