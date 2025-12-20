package main

import "fmt"

type notifier interface {
	notify() string
}

type email string
type phone string

func main() {

	myEmail := email("anuroopps@10388@gmail.com") // Type Conversion: TargetType(ValueYouHave)
	myPhone := phone("3782392993")                // // Type Conversion: TargetType(ValueYouHave)
	sendNotification(myEmail)
	sendNotification(myPhone)
}

func (ml email) notify() string {
	return "Sending email to: " + string(ml) // // Type Conversion: TargetType(ValueYouHave)
}

func (ph phone) notify() string {
	return "Sending SMS to: " + string(ph) // Type Conversion: TargetType(ValueYouHave)
}

func sendNotification(n notifier) {
	fmt.Println(n.notify())
}
