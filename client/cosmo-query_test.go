package client

import (
	"testing"
)

func TestNew(t *testing.T) {
	resp := Auth("cosmos1lntvs8pus3kf4vfn2fhhvmqe02s5ctr020heql")

	if resp.IsError() {
		t.Fatalf("non 200")
	}
}
