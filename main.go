package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Notification struct {
	id        int
	Message   string
	Timestamp time.Time
}

type MemoryStore struct {
	notification map[int]*Notification
	nextID       int
}

func (ms *MemoryStore) Add(notification *Notification) error {
	notification.id = ms.nextID
	ms.notification[notification.id] = notification
	ms.nextID++
	fmt.Println("Notification numeros :", notification.id)
	return nil
}

func (ms *MemoryStore) List() []*Notification {
	result := make([]*Notification, 0, len(ms.notification))
	for _, n := range ms.notification {
		result = append(result, n)
	}
	return result
}

type Notifier interface {
	Send(notification *Notification) error
}

type SMSNotifier struct {
	Store  *MemoryStore
	Number string
}

func (s SMSNotifier) Send(notification *Notification) error {
	if err := VerifyNumber(s.Number); err != nil {
		return err
	}
	fmt.Println("SMS envoyé à", s.Number, ":", notification.Message)
	s.Store.Add(notification)
	return nil
}

type EmailNotifier struct {
	Store *MemoryStore
	Email string
}

func (e EmailNotifier) Send(notification *Notification) error {
	fmt.Println("Email envoyé à", e.Email, ":", notification.Message)
	e.Store.Add(notification)
	return nil
}

type PushNotifier struct {
	Store  *MemoryStore
	Device string
}

func (p PushNotifier) Send(notification *Notification) error {
	fmt.Println("Push envoyé à", p.Device, ":", notification.Message)
	p.Store.Add(notification)
	return nil
}

func VerifyNumber(number string) error {
	if len(number) != 10 {
		return errors.New("numéro invalide : doit contenir 10 chiffres")
	}
	if !strings.HasPrefix(number, "06") && !strings.HasPrefix(number, "07") {
		return errors.New("numéro invalide : doit commencer par 06 ou 07")
	}
	return nil
}

func main() {
	// Créer le store
	store := &MemoryStore{
		notification: make(map[int]*Notification),
		nextID:       0,
	}

	// Créer les notifiers
	sms := SMSNotifier{Store: store, Number: "0612345678"}
	email := EmailNotifier{Store: store, Email: "test@example.com"}
	push := PushNotifier{Store: store, Device: "Test123"}

	// Créer une notification
	notif := &Notification{
		Message:   "Bonjour, test notification",
		Timestamp: time.Now(),
	}

	notifiers := []Notifier{sms, email, push}
	for _, n := range notifiers {
		err := n.Send(notif)
		if err != nil {
			fmt.Println("Erreur :", err)
		}
	}

	fmt.Println("Historique :")
	for _, n := range store.List() {
		fmt.Println(n.id, n.Timestamp, n.Message)
	}
}
