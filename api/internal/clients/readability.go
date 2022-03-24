package clients

import (
	"context"
	"os"
	"time"

	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/grpc/readabilityclient"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/grpc/readabilityserver"
	"github.com/pkg/errors"
)

// ReadabilityClient returns a grpc client for the readability service.
func ReadabilityClient() readabilityclient.ReadabilityClient {
	var addr = os.Getenv("READABILITY_SERVER_ADDR")
	if addr == "" {
		addr = readabilityserver.DefaultServerAddress
	}
	conn, err := GRPC(addr)
	if err != nil {
		panic(errors.Wrap(
			err, "clients: could not open grpc connection for readability client",
		))
	}
	return readabilityclient.NewReadabilityClient(conn)
}

// IsReadabilityClientReady returns true if the readability client is ready.
// It waits for the client to be ready for up to timeout.
func IsReadabilityClientReady(ctx context.Context, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	var ready = make(chan struct{})
	go readinessCheck(ctx, ready)

	select {
	case <-ctx.Done():
		return false
	case <-ready:
		return true
	}
}

// readinessCheck retries and checks if the readability client is ready.
func readinessCheck(ctx context.Context, resultChan chan struct{}) {
	rc := ReadabilityClient()
	for {
		deadline, ok := ctx.Deadline()
		if ok {
			if time.Now().After(deadline) {
				break
			}
		}
		info, err := rc.GetReadinessInfo(
			ctx, &readabilityclient.GetReadinessInfoRequest{
				Identifier: "readability-client",
			},
		)
		if err != nil || !info.GetReady() {
			time.Sleep(time.Millisecond * 10)
			continue
		}
		close(resultChan)
		break
	}
}
