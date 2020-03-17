package core

import "context"

// Provider - アプリケーションコアの構造体
type Provider struct {
	r Repository
}

// NewProvider - アプリケーションコアの構造体のコンストラクタ
func NewProvider(r Repository) *Provider {
	return &Provider{r}
}


// CreateUser - レシピを作成
func (p *Provider) CreateUser(ctx context.Context, user User) (User, error) {
	user, err := p.r.CreateUser(ctx, user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
