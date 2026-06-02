package context

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetString(t *testing.T) {
	testKey := Key("foo")
	testVal := "value"

	tests := []struct {
		name     string
		inputCtx context.Context
		want     string
	}{
		{
			name:     "single valid value",
			inputCtx: context.WithValue(context.Background(), testKey, testVal),
			want:     testVal,
		},
		{
			name:     "multiple valid values gets the latest",
			inputCtx: context.WithValue(context.WithValue(context.Background(), testKey, "1st value"), testKey, testVal),
			want:     testVal,
		},
		{
			name:     "invalid value type",
			inputCtx: context.WithValue(context.Background(), testKey, 1),
			want:     "",
		},
		{
			name:     "missing key",
			inputCtx: context.Background(),
			want:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetString(tt.inputCtx, testKey)
			assert.Equal(t, tt.want, got)
		})
	}
}
