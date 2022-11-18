package vo

import "sonic/model/dto"

type LinkTeamVO struct {
	Team  string
	Links []*dto.Link
}
