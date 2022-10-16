package tracing

import (
	"context"
	"os"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"github.com/tasylab/hn.hsblhsn.me/backend/internal/featureflags"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func New(lc fx.Lifecycle, logger *zap.Logger) (trace.TracerProvider, error) {
	if featureflags.IsOn(featureflags.FeatureOpentelemetry, false) {
		return trace.NewNoopTracerProvider(), nil
	}
	ctx := context.Background()
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	exporter, err := texporter.New(texporter.WithProjectID(projectID))
	if err != nil {
		logger.Error("failed to create exporter", zap.Error(err))
		return trace.NewNoopTracerProvider(), nil
	}

	res, err := resource.New(ctx,
		resource.WithDetectors(gcp.NewDetector()),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(os.Getenv("K_SERVICE")),
			semconv.ServiceVersionKey.String(os.Getenv("K_REVISION")),
		),
	)
	if err != nil {
		logger.Error("failed to create resource", zap.Error(err))
		return trace.NewNoopTracerProvider(), nil
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return tp.ForceFlush(ctx)
		},
	})
	otel.SetTracerProvider(tp)
	return tp, nil
}
