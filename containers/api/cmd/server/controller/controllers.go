package controller

// Controllers - コントローラをまとめた構造体. DIで利用.
type Controllers struct {
	Ctrl *Controller
}

// NewControllers - コントローラをまとめた構造体のコンストラクタ.
func NewControllers(ctrl *Controller) *Controllers {
	return &Controllers{ctrl}
}
