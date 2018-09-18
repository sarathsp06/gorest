package google

import (
	"context"

	"github.com/sarathsp06/gorest/utils/routes"
	"googlemaps.github.io/maps"
)

// GetDistance tries to find distance given locations as lat ,long tuples
// distance will be in kilometers
func (g Client) GetDistance(ctx context.Context, origin, destination routes.Location) (float64, error) {
	r := &maps.DistanceMatrixRequest{
		Origins:      []string{origin.String()},
		Destinations: []string{destination.String()},
	}
	results, err := g.DistanceMatrix(ctx, r)
	if err != nil {
		return 0, err
	}
	if len(results.Rows) == 0 {
		return 0, routes.ErrNotFound
	}
	result := results.Rows[0]
	if result.Elements[0].Status != "OK" {
		return 0, routes.ErrNotFound
	}
	return float64(result.Elements[0].Distance.Meters) / 1000, nil
}
