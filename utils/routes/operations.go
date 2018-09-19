package routes

import (
	"context"

	"github.com/sarathsp06/gorest/utils/metrics"
)

// GetDistance returns distance between two locations
// TODO : Cache the result and try using that ins
func GetDistance(ctx context.Context, origin, destination Location) (float64, error) {
	metrics.CaptureDelay("GetDistance")()
	//return 49.7, nil
	return route.GetDistance(ctx, origin, destination)
}
