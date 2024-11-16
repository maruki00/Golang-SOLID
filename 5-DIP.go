// Dependency Inversion Principle (DIP)

// The Dependency Inversion Principle states that high-level
// modules should not depend on low-level modules,
// but both should depend on abstractions.
// This principle promotes loose coupling between components,
// making the code more maintainable and testable.

package main

type MessageSender interface {
	Send(to string, message string)
}

type EmailService struct{}

func (e *EmailService) Send(to string, message string) {
	// Send email
}

type SMSService struct{}

func (s *SMSService) Send(to string, message string) {
	// Send SMS
}

type NotificationService struct {
	messageSender MessageSender
}

func (n *NotificationService) Notify(to string, message string) {
	n.messageSender.Send(to, message)
}

func main() {}
