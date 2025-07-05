package main

import "fmt"


type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	EmailAddress string
}

func (e EmailNotifier) Notify(message string) {
	fmt.Println("Notified via", e.EmailAddress, message)
}

type SMSNotifier struct {
	PhoneNumber string
}

func (s SMSNotifier) Notify(message string) {
	fmt.Println("Notified via", s.PhoneNumber, message)
}

func sendNotification(n Notifier, message string) {
	n.Notify(message)
}

func main() {

	message := "Hello via Notifier"

	email := EmailNotifier{EmailAddress: "hello@email.com"}
	sms := SMSNotifier{PhoneNumber: "0987654321"}

	sendNotification(email, message)
	sendNotification(sms, message)
}
