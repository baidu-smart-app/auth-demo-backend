package pay

type SuccCallBackParam struct {
	UserID         string
	OrderID        string
	UnitPrice      int64
	Count          int64
	TotalMoney     int64
	PayMoney       int64
	PromoMoney     int64
	HbMoney        int64
	HbBalanceMoney int64
	GiftCardMoney  int64
	DealID         string
	PayTime        int64
	PromoDetail    string
	PayType        int64
	PartnerID      int64
	Status         int64
	TpOrderID      string
	ReturnData     string
	RsaSign        string
}

type Pay interface {
	CreateOrder(tpOrderID string) error
	PaySucc(*SuccCallBackParam) error
	GetOrder(tpOrderID string) (*SuccCallBackParam, error)
	GetAllOrders() (map[string]*SuccCallBackParam, error)
}

var DefaulPay = NewMemoryPay()
