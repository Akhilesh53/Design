package main

import (
	pizza "pattern/Decorator/Pizza_Decorator"
	notify "pattern/ObserverPattern/NotifyMe"
)

func main() {
	// notify me button functionality :- Observer Pattern
	notify.CallNotifyMe()

	// pizza decorator function :- Decorator Pattern
	pizza.CreatePizzaDecorators()
}
