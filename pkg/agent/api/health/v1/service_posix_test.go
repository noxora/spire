//go:build !windows
// +build !windows

package health_test

import (
	"net"
	"testing"

	"github.com/spiffe/go-spiffe/v2/proto/spiffe/workload"
	"github.com/spiffe/spire/test/spiretest"
)

func startWorkloadAPI(t *testing.T, server workload.SpiffeWorkloadAPIServer) net.Addr {
	return spiretest.StartWorkloadAPIOnTempUDSSocket(t, server)
}
