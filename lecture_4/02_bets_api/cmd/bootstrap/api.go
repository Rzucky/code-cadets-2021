package bootstrap

import (
	"github.com/superbet-group/code-cadets-2021/lecture_4/02_bets_api/cmd/config"
	"github.com/superbet-group/code-cadets-2021/lecture_4/02_bets_api/internal/api"
	"github.com/superbet-group/code-cadets-2021/lecture_4/02_bets_api/internal/api/controllers"
	"github.com/superbet-group/code-cadets-2021/lecture_4/02_bets_api/internal/domain/mappers"
	sqlite "github.com/superbet-group/code-cadets-2021/lecture_4/02_bets_api/internal/infrastructure"
)

func newBetRepository(dbExecutor sqlite.DatabaseExecutor, mapper *mappers.BetMapper) *sqlite.BetRepository {
	return sqlite.NewBetRepository(dbExecutor, mapper)
}

func newController(betRepository controllers.BetRepository) *controllers.Controller {
	return controllers.NewController(betRepository)
}

func newBetMapper() *mappers.BetMapper {
	return mappers.NewBetMapper()
}

// Api bootstraps the http server.
func Api(dbExecutor sqlite.DatabaseExecutor) *api.WebServer {

	betMapper := newBetMapper()
	betRepository := newBetRepository(dbExecutor, betMapper)
	controller := newController(betRepository)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
