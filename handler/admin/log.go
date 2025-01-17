package admin

import (
	"github.com/gin-gonic/gin"

	"sonic/model/dto"
	"sonic/model/param"
	"sonic/service"
	"sonic/util"
	"sonic/util/xerr"
)

type LogHandler struct {
	LogService service.LogService
}

func NewLogHandler(logService service.LogService) *LogHandler {
	return &LogHandler{
		LogService: logService,
	}
}

func (l *LogHandler) PageLatestLog(ctx *gin.Context) (interface{}, error) {
	top, err := util.MustGetQueryInt32(ctx, "top")
	if err != nil {
		top = 10
	}
	logs, _, err := l.LogService.PageLog(ctx, param.Page{PageSize: int(top)}, &param.Sort{Fields: []string{"createTime,desc"}})
	if err != nil {
		return nil, err
	}
	logDTOs := make([]*dto.Log, 0, len(logs))
	for _, log := range logs {
		logDTOs = append(logDTOs, l.LogService.ConvertToDTO(log))
	}
	return logDTOs, nil
}

func (l *LogHandler) PageLog(ctx *gin.Context) (interface{}, error) {
	type LogParam struct {
		param.Page
		*param.Sort
	}
	var logParam LogParam
	err := ctx.ShouldBindQuery(&logParam)
	if err != nil {
		return nil, xerr.WithMsg(err, "parameter error").WithStatus(xerr.StatusBadRequest)
	}
	if logParam.Sort == nil {
		logParam.Sort = &param.Sort{
			Fields: []string{"createTime,desc"},
		}
	}
	logs, totalCount, err := l.LogService.PageLog(ctx, logParam.Page, logParam.Sort)
	if err != nil {
		return nil, err
	}
	logDTOs := make([]*dto.Log, 0, len(logs))
	for _, log := range logs {
		logDTOs = append(logDTOs, l.LogService.ConvertToDTO(log))
	}
	return dto.NewPage(logDTOs, totalCount, logParam.Page), nil
}

func (l *LogHandler) ClearLog(ctx *gin.Context) (interface{}, error) {
	return nil, l.LogService.Clear(ctx)
}
