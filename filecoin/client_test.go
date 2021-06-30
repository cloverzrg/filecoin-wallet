package filecoin

import (
	"context"
	"testing"
)

func TestChainHead(t *testing.T) {
	chainHead, err := Client.ChainHead(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(chainHead.Height())
}
