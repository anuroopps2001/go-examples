package main

import "fmt"

type Notifier interface {
	Notify(message string) string
}

type EmailNotifer struct {
	email string
}

type SMSNotifier struct {
	phone string
}

type WhatsAppNotifier struct {
	number string
}

func main() {
	EmailNotifer1 := EmailNotifer{email: "Anuroopps@112"}
	SMSNotifier1 := SMSNotifier{phone: "383793949"}
	WhatsAppNotifier1 := WhatsAppNotifier{number: "3947899394"}

	send(EmailNotifer1, "Hello, this is from EmailNotifier")
	send(SMSNotifier1, "Hello, this is from SMSNotifier")
	send(WhatsAppNotifier1, "Hello, this is from What'sAppNotifier")
}

func (e EmailNotifer) Notify(message string) string {

	return "sendind EMAIL to " + e.email + ": " + message
}

func (sms SMSNotifier) Notify(message string) string {

	return "Sending SMS to " + sms.phone + ": " + message

}

func (wh WhatsAppNotifier) Notify(message string) string {
	return "Sending Whatsapp message to " + wh.number + ": " + message
}

func send(n Notifier, message string) {
	fmt.Println(n.Notify(message))
}
