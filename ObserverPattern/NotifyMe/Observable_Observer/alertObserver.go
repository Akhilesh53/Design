package observable

import "errors"

type IAlertObserver interface {
	Update()
	GetStockObservable() *IStockObservable
}

func NewAlertObserver(stock *IStockObservable, value, alertType string) (IAlertObserver, error) {
	switch alertType {
	case "Mobile":
		return NewMobileAlertObserver(value, stock), nil
	case "Email":
		return NewEmailAlertObserver(value, stock), nil
	default:
		return nil, errors.New("unknown alert type")
	}
}
