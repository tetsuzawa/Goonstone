package user

import "context"

// Provider TODO
type Provider struct {
	r Repository
}

// NewProvider TODO
func NewProvider(r Repository) *Provider {
	return &Provider{r}
}


// CreateRecipe - レシピを作成
func (p *Provider) CreateUser(ctx context.Context, user User) (User, error) {
	recipe, err := p.r.CreateUser(ctx, user)
	if err != nil {
		return User{}, err
	}

	return recipe, nil
}
