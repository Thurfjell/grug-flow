package compose

type Resolver struct {
	layouts map[string]Layout
}

func NewResolver(layouts map[string]Layout) Resolver {
	return Resolver{
		layouts: layouts,
	}
}

func (r Resolver) Get(name string) Layout {
	return r.layouts[name]
}
