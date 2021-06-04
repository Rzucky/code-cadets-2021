package sqlite

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	controllermodel "github.com/superbet-group/code-cadets-2021/lecture_4/02_bets_api/internal/api/controllers/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/lecture_4/02_bets_api/internal/infrastructure/sqlite/models"
)

// BetRepository provides methods that operate on bets SQLite database.
type BetRepository struct {
	dbExecutor DatabaseExecutor
	betMapper  BetMapper
}

// NewBetRepository creates and returns a new BetRepository.
func NewBetRepository(dbExecutor DatabaseExecutor, betMapper BetMapper) *BetRepository {
	return &BetRepository{
		dbExecutor: dbExecutor,
		betMapper:  betMapper,
	}
}

// GetBetByID fetches a bet from the database and returns it. The second returned value indicates
// whether the bet exists in DB. If the bet does not exist, an error will not be returned.
func (r *BetRepository) GetBetByID(ctx context.Context, id string) (controllermodel.BetDto, bool, error) {

	storageBet, err := r.queryGetBetByID(ctx, id)
	if err == sql.ErrNoRows {
		return controllermodel.BetDto{}, false, nil
	}
	if err != nil {
		return controllermodel.BetDto{}, false, errors.Wrap(err, "bet repository failed to get a bet with id "+id)
	}

	domainBet := r.betMapper.MapStorageBetToBetDto(storageBet)
	return domainBet, true, nil
}

func (r *BetRepository) queryGetBetByID(ctx context.Context, id string) (storagemodels.Bet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE id=?;", id)
	if err != nil {
		return storagemodels.Bet{}, err
	}
	defer row.Close()

	// This will move to the "next" result (which is the only result, because a single bet is fetched).
	// row.Next()

	// This will move to the "next" result (which is the only result, because a single bet is fetched).
	if !row.Next() {
		return storagemodels.Bet{}, sql.ErrNoRows
	}

	var customerId string
	var status string
	var selectionId string
	var selectionCoefficient int
	var payment int
	var payoutSql sql.NullInt64

	err = row.Scan(&id, &customerId, &status, &selectionId, &selectionCoefficient, &payment, &payoutSql)
	if err != nil {
		return storagemodels.Bet{}, err
	}

	var payout int
	if payoutSql.Valid {
		payout = int(payoutSql.Int64)
	}

	return storagemodels.Bet{
		Id:                   id,
		CustomerId:           customerId,
		Status:               status,
		SelectionId:          selectionId,
		SelectionCoefficient: selectionCoefficient,
		Payment:              payment,
		Payout:               payout,
	}, nil
}

func (r *BetRepository) GetBetsByUserID(ctx context.Context, userId string) ([]controllermodel.BetDto, bool, error) {

	storageBets, err := r.queryGetBetsByUserID(ctx, userId)
	if err == sql.ErrNoRows {
		return []controllermodel.BetDto{}, false, nil
	}
	if err != nil {
		return []controllermodel.BetDto{}, false, errors.Wrap(err, "bet repository failed to get a bets with user id "+userId)
	}
	// iterating over all storage bets and mapping them to BetDto
	var usersBets []controllermodel.BetDto
	for _, userStorageBet := range storageBets {
		usersBets = append(usersBets, r.betMapper.MapStorageBetToBetDto(userStorageBet))
	}

	return usersBets, true, nil
}

func (r *BetRepository) queryGetBetsByUserID(ctx context.Context, userId string) ([]storagemodels.Bet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE customer_id=?;", userId)
	if err != nil {
		return []storagemodels.Bet{}, err
	}
	defer row.Close()

	var usersBets []storagemodels.Bet

	for row.Next() {
		var id string
		var status string
		var selectionId string
		var selectionCoefficient int
		var payment int
		var payoutSql sql.NullInt64

		err = row.Scan(&id, &userId, &status, &selectionId, &selectionCoefficient, &payment, &payoutSql)
		if err != nil {
			return []storagemodels.Bet{}, err
		}

		var payout int
		if payoutSql.Valid {
			payout = int(payoutSql.Int64)
		}
		usersBets = append(usersBets, storagemodels.Bet{
			Id:                   id,
			CustomerId:           userId,
			Status:               status,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
			Payout:               payout,
		})
	}

	return usersBets, nil
}

func (r *BetRepository) GetBetByStatus(ctx context.Context, status string) ([]controllermodel.BetDto, bool, error) {
	storageBets, err := r.queryGetBetByStatus(ctx, status)
	if err == sql.ErrNoRows {
		return []controllermodel.BetDto{}, false, nil
	}
	if err != nil {
		return []controllermodel.BetDto{}, false, errors.Wrap(err, "bet repository failed to get active bets")
	}
	// iterating over all storage bets and mapping them to BetDto
	var activeBets []controllermodel.BetDto
	for _, userStorageBet := range storageBets {
		activeBets = append(activeBets, r.betMapper.MapStorageBetToBetDto(userStorageBet))
	}

	return activeBets, true, nil
}

func (r *BetRepository) queryGetBetByStatus(ctx context.Context, status string) ([]storagemodels.Bet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE status=?;", status)
	if err != nil {
		return []storagemodels.Bet{}, err
	}
	defer row.Close()

	var activeBets []storagemodels.Bet

	for row.Next() {
		var id string
		var customerId string
		var selectionId string
		var selectionCoefficient int
		var payment int
		var payoutSql sql.NullInt64

		err = row.Scan(&id, &customerId, &status, &selectionId, &selectionCoefficient, &payment, &payoutSql)
		if err != nil {
			return []storagemodels.Bet{}, err
		}

		var payout int
		if payoutSql.Valid {
			payout = int(payoutSql.Int64)
		}
		activeBets = append(activeBets, storagemodels.Bet{
			Id:                   id,
			CustomerId:           customerId,
			Status:               status,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
			Payout:               payout,
		})
	}

	return activeBets, nil
}
