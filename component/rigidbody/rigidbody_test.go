package rigidbody

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
		wantMass *float64
		wantType RidigbodyPhysicsType
	}{
		{
			name: "should set mass and type correctly",
			data: map[string]interface{}{
				"mass": 10.0,
				"type": "dynamic",
			},
			wantErr:  false,
			wantMass: func() *float64 { f := 10.0; return &f }(),
			wantType: RigidbodyTypeDynamic,
		},
		{
			name: "should not return error when mass is missing",
			data: map[string]interface{}{
				"type": "dynamic",
			},
			wantErr: false,
		},
		{
			name: "should return error when mass is not a float64",
			data: map[string]interface{}{
				"type": "dynamic",
				"mass": "heavy",
			},
			wantErr: true,
		},
		{
			name: "should return error when type is invalid",
			data: map[string]interface{}{
				"type": "invalid",
			},
			wantErr: true,
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
				if tt.wantMass != nil {
					assert.NotNil(t, rb.mass)
					assert.EqualValues(t, *tt.wantMass, *rb.mass)
				}
				assert.Equal(t, tt.wantType, rb.Type)
			}
		})
	}
}
