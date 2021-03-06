package packasset

type Generator struct {
	packageName string
	files       []string
	stripPrefix string
}

type Option interface {
	Name() string
	Value() interface{}
}

type option struct {
	name  string
	value interface{}
}
