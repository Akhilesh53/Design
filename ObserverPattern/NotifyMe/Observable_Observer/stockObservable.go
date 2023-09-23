package observable

import "errors"

type IStockObservable interface {
	Add(product IAlertObserver)
	Remove(product IAlertObserver)
	SetData(count int)
	GetData() int
	Notify()
}

func NewObservable(observableType string) (IStockObservable, error) {
	switch observableType {
	case "Iphone":
		return NewIphoneObservable(), nil
	default:
		return nil, errors.New("unsupported observable type")

	}
}
