package observer

import "fmt"

type EventCenter struct {
	observers []Observer //观察者列表
}

func (this *EventCenter) Notify(event Event) {
	for _, v := range this.observers {
		v.OnNotify(event)
	}
}

func (this *EventCenter) Register(o Observer) {
	this.observers = append(this.observers, o)
}

func (this *EventCenter) Deregister(o Observer) {
	for i := 0; i < len(this.observers); i++ {
		if this.observers[i] == o {
			this.observers = append(this.observers[:i], this.observers[i+1:]...)
			break
		}
	}
}

func NewEventCenter() *EventCenter {
	return &EventCenter{make([]Observer, 0)}
}

type EventReciver struct {
}

func (this *EventReciver) OnNotify(event Event) {
	fmt.Printf("Event Recive:%d\n", event.Data)
}
