package vo

import "sonic/model/dto"

type Menu struct {
	dto.Menu
	Children []*Menu `json:"children"`
}
