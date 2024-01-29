package boxcollider

const (
	Type uint = 3
	Slug      = "boxCollider"
)

type BoxCollider struct{}

func (*BoxCollider) SetData(data map[string]any) error {
	return nil
}
