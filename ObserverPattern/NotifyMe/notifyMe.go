package notifyme
/*

 Design Question:
 ======================
 There is notify me button in amazon/ flipkart website that notifies the
 persons about the product when it comes back from out of stock.

 Implement the functionality for the same

*/

import (
	"fmt"
	ob "pattern/ObserverPattern/NotifyMe/Observable_Observer"
)

func CallNotifyMe() {

	fmt.Println("================================================")
	fmt.Println(" Calling Notify Me Observer Pattern Func")
	fmt.Println("================================================")

	defer fmt.Println("================================================")

	iphoneStockObserver, err := ob.NewObservable("Iphone")
	if err != nil {
		fmt.Println(err)
		return
	}

	alterObserver1, err := ob.NewAlertObserver(&iphoneStockObserver, "9906877909", "Mobile")
	if err != nil {
		fmt.Println(err)
		return
	}
	alterObserver2, err := ob.NewAlertObserver(&iphoneStockObserver, "abc@gmail.com", "Email")
	if err != nil {
		fmt.Println(err)
		return
	}

	iphoneStockObserver.Add(alterObserver1)
	iphoneStockObserver.Add(alterObserver2)

	iphoneStockObserver.SetData(12)
}
