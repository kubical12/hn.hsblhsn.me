package images

import (
	"testing"

	"github.com/hsblhsn/hn.hsblhsn.me/featureflags"
)

// nolint:paralleltest // not a test
func TestProxyURL(t *testing.T) {
	t.Setenv("DOMAIN", "hn.hsblhsn.me")
	t.Setenv(featureflags.FeatureImgProxy.String(), "on")
	out := ProxiedURL("https://cdn.substack.com/image/fetch/w_2400,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fbucketeer-e05bbc84-baa3-437e-9518-adb32be77984.s3.amazonaws.com%2Fpublic%2Fimages%2F4f120181-9b0e-480b-a1ca-05fd8201c6c1_1600x900.png", ImageSizeFull)
	t.Log(out)
}
