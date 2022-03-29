package transports

import (
	"context"
	"net/http"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/types"
)

func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	types.EncodeErrorMessage(ctx, err, w)
}
