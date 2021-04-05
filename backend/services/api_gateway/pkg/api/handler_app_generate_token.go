package api

import (
	"net/http"

	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (api API) HandlerAppGenerateToken(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerAppGenerateToken] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	// resp, err := api.AppServiceClient.GenerateToken()
}
