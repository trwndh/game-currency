package errors

import "testing"

func TestGetErrorConversionAlreadyExist(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "error message",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetErrorConversionAlreadyExist(); (err != nil) != tt.wantErr {
				t.Errorf("GetErrorConversionAlreadyExist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetErrorConversionNotFound(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "error message",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetErrorConversionNotFound(); (err != nil) != tt.wantErr {
				t.Errorf("GetErrorConversionNotFound() error = %v, wantErr %v", err, tt.wantErr)
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
			name:    "error message",
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

func TestGetErrorDatabase(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "error message",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetErrorDatabase(); (err != nil) != tt.wantErr {
				t.Errorf("GetErrorDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetErrorRateIsZero(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "error message",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetErrorRateIsZero(); (err != nil) != tt.wantErr {
				t.Errorf("GetErrorRateIsZero() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetErrorConvertingSameID(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "error message",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetErrorConvertingSameID(); (err != nil) != tt.wantErr {
				t.Errorf("GetErrorConvertingSameID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetErrorInvalidCurrencyID(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "error message",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetErrorInvalidCurrencyID(); (err != nil) != tt.wantErr {
				t.Errorf("GetErrorInvalidCurrencyID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
