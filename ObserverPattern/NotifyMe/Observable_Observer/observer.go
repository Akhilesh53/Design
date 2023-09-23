package observable

import (
	"fmt"
)

type MobileAlertObserver struct {
	StockObservable *IStockObservable
	MobileNumber    string
}

func NewMobileAlertObserver(mobile string, stock *IStockObservable) *MobileAlertObserver {
	return &MobileAlertObserver{
		StockObservable: stock,
		MobileNumber:    mobile,
	}
}

func (m *MobileAlertObserver) Update() {
	m.sendMessage()
}

func (m *MobileAlertObserver) GetStockObservable() (*IStockObservable){
	return m.StockObservable
}

func (m *MobileAlertObserver) sendMessage() {
	fmt.Println("sending message to ", m.MobileNumber)
}

type EmailAlertObserver struct {
	StockObservable *IStockObservable
	Email           string
}

func NewEmailAlertObserver(email string, stock *IStockObservable) *EmailAlertObserver {
	return &EmailAlertObserver{
		StockObservable: stock,
		Email:           email,
	}
}

func (e *EmailAlertObserver) Update() {
	e.sendEmail()
}

func (e *EmailAlertObserver) sendEmail() {
	fmt.Println("sending email to ", e.Email)
}

func (m *EmailAlertObserver) GetStockObservable() (*IStockObservable){
	return m.StockObservable
}
//anthesupport@aesl.in
