package main

import "fmt"

type NotificationService struct {
	notifierType string
}

func (s NotificationService) SendNoti(message string) {
	if s.notifierType == "email" {
		fmt.Printf("Sending message: %s (Sender: Email)\n", message)
	} else if s.notifierType == "sms" {
		fmt.Printf("Sending message: %s (Sender: SMS)\n", message)
	}
}

func main() {
	s := NotificationService{notifierType: "email"}
	s.SendNoti("Hello Dat")
}
