package core

// Provider TODO
type Provider struct {
	r Repository
}

// NewProvider TODO
func NewProvider(r Repository) *Provider {
	return &Provider{r}
}
