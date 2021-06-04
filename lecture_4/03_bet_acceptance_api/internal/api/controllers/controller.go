package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/superbet-group/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/api/controllers/models"
)

// Controller implements handlers for web server requests.
type Controller struct {
	betValidator BetValidator
	betService   BetService
}

// NewController creates a new instance of Controller
func NewController(betValidator BetValidator, betService BetService) *Controller {
	return &Controller{
		betValidator: betValidator,
		betService:   betService,
	}
}

// AddBet handlers adding a bet to queue.
func (e *Controller) AddBet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var betRequestDto models.BetRequestDto
		err := ctx.ShouldBindWith(&betRequestDto, binding.JSON)
		if err != nil {
			ctx.String(http.StatusBadRequest, "bet request is not valid.")
			return
		}

		if !e.betValidator.BetInputIsValid(betRequestDto) {
			ctx.String(http.StatusBadRequest, "bet input is not valid.")
			return
		}

		err = e.betService.SendBetWithId(betRequestDto)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}

		ctx.Status(http.StatusOK)
	}
}
