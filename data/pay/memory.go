package pay

type MemoryPay struct {
	orders map[string]*SuccCallBackParam
}

func NewMemoryPay() Pay {
	return &MemoryPay{
		orders: map[string]*SuccCallBackParam{},
	}

}

func (mp *MemoryPay) GetAllOrders() (map[string]*SuccCallBackParam, error) {
	return mp.orders, nil
}

func (mp *MemoryPay) CreateOrder(tpOrderID string) error {
	mp.orders[tpOrderID] = nil
	return nil
}

func (mp *MemoryPay) PaySucc(scbp *SuccCallBackParam) error {
	mp.orders[scbp.TpOrderID] = scbp
	return nil
}

func (mp *MemoryPay) GetOrder(tpOrderID string) (*SuccCallBackParam, error) {
	return mp.orders[tpOrderID], nil
}
