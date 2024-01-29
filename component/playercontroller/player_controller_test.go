package playercontroller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerController_SetData(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		data      map[string]interface{}
		wantErr   bool
		wantSpeed float64
	}{
		{
			name: "should set speed correctly",
			data: map[string]interface{}{
				"speed": 10.0,
			},
			wantErr:   false,
			wantSpeed: 10.0,
		},
		{
			name:    "should return error when speed is missing",
			data:    map[string]interface{}{},
			wantErr: true,
		},
		{
			name: "should return error when speed is not a float64",
			data: map[string]interface{}{
				"speed": "fast",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			pc := &PlayerController{}
			err := pc.SetData(tt.data)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantSpeed, pc.Speed)
			}
		})
	}
}
