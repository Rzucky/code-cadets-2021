package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller implements handlers for web server requests.
type Controller struct {
	betRepository BetRepository
}

const (
	wonStatus    = "won"
	lostStatus   = "lost"
	activeStatus = "active"
)

// NewController creates a new instance of Controller
func NewController(betRepository BetRepository) *Controller {
	return &Controller{
		betRepository: betRepository,
	}
}

// GetBet handlers bet request.
func (e *Controller) GetBet() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		betId := ctx.Param("id")
		if betId == "" {
			ctx.String(http.StatusBadRequest, "id is not valid.")
			return
		}

		foundBet, found, err := e.betRepository.GetBetByID(ctx, betId)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed,", err)
			return
		}
		if !found {
			ctx.String(http.StatusNotFound, "bet with given id does not exist.")
			return
		}
		ctx.JSON(http.StatusOK, foundBet)
	}
}

func (e *Controller) GetBetsByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("id")
		if userId == "" {
			ctx.String(http.StatusBadRequest, "user id is not valid.")
			return
		}
		foundBets, found, err := e.betRepository.GetBetsByUserID(ctx, userId)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed,", err)
			return
		}
		if !found {
			ctx.String(http.StatusNotFound, "bets with given user id do not exist.")
			return
		}

		ctx.JSON(http.StatusOK, foundBets)

	}
}

func (e *Controller) GetBetByStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := ctx.Query("status")

		if status != lostStatus && status != wonStatus && status != activeStatus {
			ctx.String(http.StatusBadRequest, "status is not valid.")
			return
		}

		foundBets, found, err := e.betRepository.GetBetByStatus(ctx, status)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed,", err)
			return
		}
		if !found {
			ctx.String(http.StatusNotFound, "bets with given status do not exist.")
			return
		}

		ctx.JSON(http.StatusOK, foundBets)

	}
}
