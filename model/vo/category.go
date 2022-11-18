package vo

import "sonic/model/dto"

type CategoryVO struct {
	dto.CategoryDTO
	Children []*CategoryVO `json:"children"`
}
