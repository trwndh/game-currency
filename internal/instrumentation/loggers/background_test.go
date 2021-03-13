package loggers

import (
	"reflect"
	"testing"

	"go.uber.org/zap"
)

func TestBg(t *testing.T) {
	tests := []struct {
		name string
		want *zap.Logger
	}{
		{
			name: "positive case",
			want: Bg(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bg(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bg() = %v, want %v", got, tt.want)
			}
		})
	}
}
