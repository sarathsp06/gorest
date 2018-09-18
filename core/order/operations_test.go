package order

import (
	"context"
	"reflect"
	"testing"

	"github.com/sarathsp06/gorest/utils/routes"
)

func TestCreate(t *testing.T) {
	type args struct {
		ctx         context.Context
		origin      routes.Location
		destination routes.Location
	}
	tests := []struct {
		name    string
		args    args
		want    *Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.ctx, tt.args.origin, tt.args.destination)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTake(t *testing.T) {
	type args struct {
		ctx    context.Context
		status Status
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Take(tt.args.ctx, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("Take() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestList(t *testing.T) {
	type args struct {
		ctx   context.Context
		page  int
		limit int
	}
	tests := []struct {
		name    string
		args    args
		want    []Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.ctx, tt.args.page, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}
