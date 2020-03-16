package user

import "github.com/go-playground/validator/v10"

// Validate - バリデータの実体. 外部に公開する.
var Validate *validator.Validate

func init() {
	Validate = validator.New()
}
