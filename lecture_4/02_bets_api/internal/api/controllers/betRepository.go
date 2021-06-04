package controllers

import (
	"context"

	controllermodels "github.com/superbet-group/code-cadets-2021/lecture_4/02_bets_api/internal/api/controllers/models"
)

// BetRepository implements bet related functions.
type BetRepository interface {
	GetBetByID(ctx context.Context, id string) (controllermodels.BetDto, bool, error)
	GetBetsByUserID(ctx context.Context, userId string) ([]controllermodels.BetDto, bool, error)
	GetBetByStatus(ctx context.Context, status string) ([]controllermodels.BetDto, bool, error)
}
