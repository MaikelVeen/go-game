package component

import (
	"testing"

	"github.com/MaikelVeen/go-game/assets"
	"github.com/stretchr/testify/assert"
)

func TestSpriteRender_SetData(t *testing.T) {
	assets.GlobalAssetRegistry = assets.NewAssetRegistry()
	assets.GlobalAssetRegistry.SpriteSheets["testSpriteSheet"] = assets.SpriteSheet{}

	type args struct {
		data map[string]interface{}
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should set SpriteSheet, X and Y correctly",
			args: args{
				data: map[string]interface{}{
					"spriteSheet": "testSpriteSheet",
					"x":           10,
					"y":           20,
				},
			},
			wantErr: false,
		},
		{
			name: "should return error when spriteSheet is missing",
			args: args{
				data: map[string]interface{}{
					"x": 10,
					"y": 20,
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when X is missing",
			args: args{
				data: map[string]interface{}{
					"spriteSheet": "testSpriteSheet",
					"y":           20,
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when Y is missing",
			args: args{
				data: map[string]interface{}{
					"spriteSheetName": "testSpriteSheet",
					"x":               10,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &SpriteRender{}
			err := sr.SetData(tt.args.data)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.args.data["x"], sr.X)
				assert.Equal(t, tt.args.data["y"], sr.Y)
				assert.Equal(t, assets.GlobalAssetRegistry.SpriteSheets[tt.args.data["spriteSheet"].(string)], sr.SpriteSheet)
			}
		})
	}
}
