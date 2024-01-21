package component

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransform_SetData(t *testing.T) {
	t.Parallel()

	type fields struct {
		X int
		Y int
	}
	type args struct {
		data map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should set X and Y correctly",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				data: map[string]interface{}{
					"x": 10,
					"y": 20,
				},
			},
			wantErr: false,
		},
		{
			name: "should return error when X is missing",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				data: map[string]interface{}{
					"y": 20,
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when Y is missing",
			fields: fields{
				X: 0,
				Y: 0,
			},
			args: args{
				data: map[string]interface{}{
					"x": 10,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tr := &Transform{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			err := tr.SetData(tt.args.data)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.args.data["x"], tr.X)
				assert.Equal(t, tt.args.data["y"], tr.Y)
			}
		})
	}
}
