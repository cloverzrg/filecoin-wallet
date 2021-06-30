package filecoin

import (
	"context"
	"github.com/cloverzrg/filecoin-wallet/logger"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"net/http"
)

var Client api.FullNodeStruct

func init() {
	authToken := ""
	headers := http.Header{"Authorization": []string{"Bearer " + authToken}}
	addr := "calibration.node.glif.io"

	_, err := jsonrpc.NewMergeClient(context.Background(), "https://"+addr+"/rpc/v0", "Filecoin", []interface{}{&Client.Internal, &Client.CommonStruct.Internal}, headers)
	if err != nil {
		logger.Panicf("connecting with lotus failed: %s", err)
	}
}