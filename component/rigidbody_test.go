package component

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRigidbody_SetData(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		data     map[string]interface{}
		wantErr  bool
		wantMass float64
	}{
		{
			name: "should set mass correctly",
			data: map[string]interface{}{
				"mass": 10.0,
			},
			wantErr:  false,
			wantMass: 10.0,
		},
		{
			name:    "should not return error when mass is missing",
			data:    map[string]interface{}{},
			wantErr: false,
		},
		{
			name: "should not return error when mass is not a float64",
			data: map[string]interface{}{
				"mass": "heavy",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			rb := &Rigidbody{}
			err := rb.SetData(tt.data)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantMass, rb.Mass)
			}
		})
	}
}
