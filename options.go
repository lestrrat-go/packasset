package packasset

func (o option) Name() string {
	return o.name
}

func (o option) Value() interface{} {
	return o.value
}

func WithPackageName(s string) Option {
	return &option{
		name:  "package_name",
		value: s,
	}
}

func WithFiles(l []string) Option {
	return &option{
		name:  "files",
		value: l,
	}
}

func WithStripPrefix(s string) Option {
	return &option{
		name: "strip_prefix",
		value: s,
	}
}
