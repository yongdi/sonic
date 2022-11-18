package service

import (
	"context"

	"sonic/model/dto"
)

type StatisticService interface {
	Statistic(ctx context.Context) (*dto.Statistic, error)
	StatisticWithUser(ctx context.Context) (*dto.StatisticWithUser, error)
}
