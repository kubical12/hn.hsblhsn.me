package readerviews

import (
	"context"
	"os"
	"time"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/grpc/readabilityclient"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/grpc/readabilityserver"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// readabilityClient returns a grpc client for the readability service.
//
//nolint:ireturn // readabilityclient.NewReadabilityClient returns interface.
func readabilityClient() readabilityclient.ReadabilityClient {
	addr := os.Getenv("READABILITY_SERVER_ADDR")
	if addr == "" {
		addr = readabilityserver.DefaultServerAddress
	}
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(errors.Wrap(
			err, "clients: could not open grpc connection for readability client",
		))
	}
	return readabilityclient.NewReadabilityClient(conn)
}

// isReadabilityClientReady returns true if the readability client is ready.
// It waits for the client to be ready for up to timeout.
func isReadabilityClientReady(ctx context.Context, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	ready := make(chan struct{})
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
	const retryInterval = time.Millisecond * 10
	client := readabilityClient()
	for {
		deadline, ok := ctx.Deadline()
		if ok {
			if time.Now().After(deadline) {
				break
			}
		}
		info, err := client.GetReadinessInfo(
			ctx, &readabilityclient.GetReadinessInfoRequest{
				Identifier: "readability-client",
			},
		)
		if err != nil || !info.GetReady() {
			time.Sleep(retryInterval)
			continue
		}
		close(resultChan)
		break
	}
}
