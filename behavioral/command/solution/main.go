package main

import (
	"fmt"
)

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: Email)", message)
}

type SmsNotifier struct{}

func (sms SmsNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: SMS)", message)
}

type NotiService struct {
	notifier Notifier
}

func (s NotiService) SendNoti(message string) {
	s.notifier.Send(message)
}

func main() {
	s := NotiService{
		notifier: SmsNotifier{},
	}
	s.notifier.Send("Hello")
}
