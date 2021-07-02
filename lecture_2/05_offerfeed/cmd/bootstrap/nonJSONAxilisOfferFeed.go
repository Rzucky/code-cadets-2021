package bootstrap

import (
	stdhttp "net/http"

	"code-cadets-2021/lecture_2/05_offerfeed/internal/infrastructure/http"
)

func NewNonJSONAxilisOfferFeed() *http.NonJSONAxilisOfferFeed {
	return http.NewNonJSONAxilisOfferFeed(stdhttp.Client{})
}
