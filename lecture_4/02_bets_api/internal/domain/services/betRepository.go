package services

import (
	"context"
	storagemodels "github.com/superbet-group/code-cadets-2021/lecture_4/02_bets_api/internal/infrastructure/sqlite/models"
)

// BetRepository implements bet related functions.
type BetRepository interface {
	GetBetByID(ctx context.Context, id string) (storagemodels.Bet, bool, error)
	GetBetsByUserID(ctx context.Context, userId string) ([]storagemodels.Bet, bool, error)
	GetBetByStatus(ctx context.Context, status string) ([]storagemodels.Bet, bool, error)
}
