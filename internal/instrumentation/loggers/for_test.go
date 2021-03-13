package loggers

import (
	"context"
	"reflect"
	"testing"

	"go.uber.org/zap"
)

func TestFor(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *zap.Logger
	}{
		{
			name: "positive case",
			args: args{
				ctx: context.Background(),
			},
			want: For(context.Background()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := For(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("For() = %v, want %v", got, tt.want)
			}
		})
	}
}
