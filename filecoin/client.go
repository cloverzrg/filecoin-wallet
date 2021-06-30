package filecoin

import (
	"context"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"log"
	"net/http"
)

var Client api.FullNodeStruct

func init() {
	authToken := ""
	headers := http.Header{"Authorization": []string{"Bearer " + authToken}}
	addr := "calibration.node.glif.io"

	_, err := jsonrpc.NewMergeClient(context.Background(), "https://"+addr+"/rpc/v0", "Filecoin", []interface{}{&Client.Internal, &Client.CommonStruct.Internal}, headers)
	if err != nil {
		log.Fatalf("connecting with lotus failed: %s", err)
	}
}