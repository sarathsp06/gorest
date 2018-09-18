package google

import (
	"context"
	"testing"

	"github.com/sarathsp06/gorest/utils/routes"
)

func getClient() *Client {
	key := "AIzaSyD9iObJCBZaiFN_NosQAj-iVhcubBx_j-w"
	c, _ := New(key)
	return c
}

func TestClient_GetDistance(t *testing.T) {
	type args struct {
		ctx         context.Context
		origin      routes.Location
		destination routes.Location
	}
	tests := []struct {
		name    string
		client  *Client
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:   "Valid locations",
			client: getClient(),
			args: args{
				ctx:         context.Background(),
				origin:      routes.Location{Latitude: 9.238487, Longitude: 76.531479}, //mavelikkara
				destination: routes.Location{Latitude: 8.893212, Longitude: 76.614143}, //kollam
			},
			want:    49.35,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.client
			got, err := g.GetDistance(tt.args.ctx, tt.args.origin, tt.args.destination)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.GetDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
