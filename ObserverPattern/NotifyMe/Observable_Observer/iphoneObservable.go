package observable

type IphoneObservable struct {
	ObserverList []IAlertObserver
	count        int
}

func NewIphoneObservable() *IphoneObservable {
	return &IphoneObservable{
		ObserverList: make([]IAlertObserver, 0),
		count:        0,
	}
}

func (obs *IphoneObservable) Add(product IAlertObserver) {
	obs.ObserverList = append(obs.ObserverList, product)
}

func (obs *IphoneObservable) Remove(deleteProduct IAlertObserver) {
	newList := []IAlertObserver{}

	for _, product := range obs.ObserverList {
		if product != deleteProduct {
			newList = append(newList, product)
		}
	}
	obs.ObserverList = newList
}

func (obs *IphoneObservable) SetData(count int) {
	if obs.count == 0 {
		obs.Notify()
	}
	obs.count = obs.count + count
}

func (obs *IphoneObservable) GetData() int {
	return obs.count
}

func (obs *IphoneObservable) Notify() {
	for _, observer := range obs.ObserverList {
		observer.Update()
	}
}
