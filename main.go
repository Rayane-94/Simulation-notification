package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Contact est notre structure de données centrale
type Contact struct {
	ID    int
	Name  string
	Email string
}

// Storer est un CONTRAT de stockage
// Il définit un ensemble de comportements (méthodes) que tout type
// de stockage doit respecter. On ne se soucie par du comment c'est fait
// (en mémoire, fichier, BDD...) seulement de ce qui peut être fait
type Storer interface {
	Add(contact *Contact) error
	GetAll() ([]*Contact, error)
	GetByID(id int) (*Contact, error)
	Update(id int, newName, newEmail string) error
	Delete(id int) error
}

type MemoryStore struct {
	contacts map[int]*Contact
	nextID   int
}

// NewMemoryStore est un constructeur qui initialise proprement notre storer
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[int]*Contact),
		nextID:   1,
	}
}

func (ms *MemoryStore) Add(contact *Contact) error {
	contact.ID = ms.nextID
	ms.contacts[contact.ID] = contact
	ms.nextID++
	return nil
}

func (ms *MemoryStore) GetAll() ([]*Contact, error) {
	var allContacts []*Contact
	for _, c := range ms.contacts {
		allContacts = append(allContacts, c)
	}
	return allContacts, nil
}

func (ms *MemoryStore) GetByID(id int) (*Contact, error) {
	contact, ok := ms.contacts[id]
	if !ok {
		return nil, errors.New("Contact not found")
	}
	return contact, nil
}

func (ms *MemoryStore) Update(id int, newName, newEmail string) error {
	contact, err := ms.GetByID(id)
	if err != nil {
		return err
	}
	if newName != "" {
		contact.Name = newName
	}
	if newEmail != "" {
		contact.Email = newEmail
	}
	return nil
}

func (ms *MemoryStore) Delete(id int) error {
	if _, ok := ms.contacts[id]; !ok {
		return errors.New("Contact not found")
	}
	delete(ms.contacts, id)
	return nil
}

func main() {

	var store Storer = NewMemoryStore()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Mini CRM v3!")

	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Add a contact")
		fmt.Println("2. List contacts")
		fmt.Println("3. Update a contact")
		fmt.Println("4. Delete a contact")
		fmt.Println("5. Exit")
		fmt.Print("Your choice: ")

		choice := readUserChoice(reader)

		switch choice {
		case 1:
			handleAddContact(reader, store)
		case 2:
			handleListContacts(store)
		case 3:
			handleUpdateContact(reader, store)
		case 4:
			handleDeleteContact(reader, store)
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, please try again")

		}
	}
}

// Les fonctions "handle..." s'occupent de l'interaction avec l'utilisateur
// et elles appellent la couche de stockage (store) pour effectuer les opérations.
// Elles sont découplées du stockage : elles fonctionnent avec n'importe quel storer

func handleAddContact(reader *bufio.Reader, storer Storer) {
	fmt.Print("Enter contact name: ")
	name := readLine(reader)

	fmt.Print("Enter contact email: ")
	email := readLine(reader)

	contact := &Contact{
		Name:  name,
		Email: email,
	}
	err := storer.Add(contact)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Contact '%s' added with ID %d.\n", contact.Name, contact.ID)
}

func handleListContacts(store Storer) {
	contacts, err := store.GetAll()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(contacts) == 0 {
		fmt.Println(" No contacts to display.")
		return
	}

	fmt.Println("\n--- Contact List ---")
	for _, contact := range contacts {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", contact.ID, contact.Name, contact.Email)
	}
}

func handleUpdateContact(reader *bufio.Reader, store Storer) {
	fmt.Print("Enter the ID of the contact to update: ")
	id := readInteger(reader)
	if id == -1 {
		return
	}

	// On vérifie que le contact existe avant de demander les nouvelles infos
	existingContact, err := store.GetByID(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Updating '%s'. Leave blank to keep current value.\n", existingContact.Name)

	fmt.Printf("New name (%s): ", existingContact.Name)
	newName := readLine(reader)

	fmt.Printf("New email (%s): ", existingContact.Email)
	newEmail := readLine(reader)

	err = store.Update(id, newName, newEmail)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Contact updated successfully.")
}

func handleDeleteContact(reader *bufio.Reader, store Storer) {
	fmt.Print("Enter the ID of the contact to delete: ")
	id := readInteger(reader)
	if id == -1 {
		return
	}

	err := store.Delete(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Contact with ID %d has been deleted.\n", id)
} // Fonctions utilitaires pour la saisie utilisateur

func readLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readUserChoice(reader *bufio.Reader) int {
	choice, err := strconv.Atoi(readLine(reader))
	if err != nil {
		return -1 // Renvoie -1 pour un choix invalide
	}
	return choice
}

func readInteger(reader *bufio.Reader) int {
	id, err := strconv.Atoi(readLine(reader))
	if err != nil {
		fmt.Println("Error: Invalid ID. Please enter a number.")
		return -1
	}
	return id
}
