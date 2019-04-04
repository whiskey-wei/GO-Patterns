package main

import "GO-Patterns/Observer/observer"

func main() {
	var EventCenter observer.Notifier
	EventCenter = observer.NewEventCenter()
	var Ober1, Ober2 observer.Observer
	Ober1 = &observer.EventReciver{}
	Ober2 = &observer.EventReciver{}

	EventCenter.Register(Ober1)
	EventCenter.Register(Ober2)

	EventCenter.Notify(observer.Event{Data: 1})

	EventCenter.Deregister(Ober2)

	EventCenter.Notify(observer.Event{Data: 2})
}
