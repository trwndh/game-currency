package errors

import "testing"

func TestGetErrorCurrencyAlreadyExist(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "error happens",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetErrorCurrencyAlreadyExist(); (err != nil) != tt.wantErr {
				t.Errorf("GetErrorCurrencyAlreadyExist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetErrorInvalidPayload(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "error happens",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetErrorInvalidPayload(); (err != nil) != tt.wantErr {
				t.Errorf("GetErrorInvalidPayload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
