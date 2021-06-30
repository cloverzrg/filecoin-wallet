package filecoin

import (
	"context"
	"github.com/cloverzrg/filecoin-wallet/config"
	"github.com/cloverzrg/filecoin-wallet/logger"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"net/http"
)

var Client api.FullNodeStruct

func init() {
	authToken := ""
	headers := http.Header{"Authorization": []string{"Bearer " + authToken}}

	_, err := jsonrpc.NewMergeClient(context.Background(), config.Endpoint, "Filecoin", []interface{}{&Client.Internal, &Client.CommonStruct.Internal}, headers)
	if err != nil {
		logger.Panicf("connecting with lotus failed: %s", err)
	}
}