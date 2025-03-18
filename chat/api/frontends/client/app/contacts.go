package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

const configFileName = "config.json"

type User struct {
	ID       common.Address
	Name     string
	Messages []string
}

type Contacts struct {
	filePath string
	me       User
	contacts map[common.Address]User
	mu       sync.RWMutex
	fileName string
}

func NewContacts(filePath string, id common.Address) (*Contacts, error) {
	os.MkdirAll(filepath.Join(filePath, "contacts"), os.ModePerm)

	fileName := filepath.Join(filePath, configFileName)

	var doc document

	_, err := os.Stat(fileName)
	switch {
	case err != nil:
		doc, err = createConfig(fileName, id)

	default:
		doc, err = readConfig(fileName)
		if doc.User.ID != id {
			return nil, fmt.Errorf("id mismatch")
		}
	}

	if err != nil {
		return nil, fmt.Errorf("config: %w", err)
	}

	contacts := make(map[common.Address]User, len(doc.Contacts))
	for _, user := range doc.Contacts {
		contacts[user.ID] = User{
			ID:   user.ID,
			Name: user.Name,
		}
	}

	cfg := Contacts{
		filePath: filePath,
		me: User{
			ID:   doc.User.ID,
			Name: doc.User.Name,
		},
		contacts: contacts,
		fileName: fileName,
	}

	return &cfg, nil
}

func (c *Contacts) My() User {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.me
}

func (c *Contacts) Contacts() []User {
	c.mu.RLock()
	defer c.mu.RUnlock()

	users := make([]User, 0, len(c.contacts))
	for _, user := range c.contacts {
		users = append(users, user)
	}

	return users
}

func (c *Contacts) LookupContact(id common.Address) (User, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	u, exists := c.contacts[id]
	if !exists {
		return User{}, fmt.Errorf("contact not found")
	}

	return u, nil
}

func (c *Contacts) AddContact(id common.Address, name string) (User, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	doc, err := readConfig(c.fileName)
	if err != nil {
		return User{}, fmt.Errorf("config read: %w", err)
	}

	du := docUser{
		ID:   id,
		Name: name,
	}

	c.contacts[id] = User{
		ID:   id,
		Name: name,
	}

	doc.Contacts = append(doc.Contacts, du)

	writeConfig(c.fileName, doc)

	u := User{
		ID:   id,
		Name: name,
	}

	return u, nil
}

func (c *Contacts) AddMessage(id common.Address, msg string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	u, exists := c.contacts[id]
	if !exists {
		return fmt.Errorf("contact not found")
	}

	u.Messages = append(u.Messages, msg)
	c.contacts[id] = u

	if err := c.writeMessage(id, msg); err != nil {
		return fmt.Errorf("write message: %w", err)
	}

	return nil
}

func (c *Contacts) ReadMessages(id common.Address) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	u, exists := c.contacts[id]
	if !exists {
		return fmt.Errorf("contact not found")
	}

	if len(u.Messages) > 0 {
		return nil
	}

	fileName := filepath.Join(c.filePath, "contacts", id.Hex()+".msg")

	f, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		u.Messages = append(u.Messages, s)
	}

	c.contacts[id] = u

	return nil
}

func (c *Contacts) writeMessage(id common.Address, msg string) error {
	var f *os.File

	fileName := filepath.Join(c.filePath, "contacts", id.Hex()+".msg")

	_, err := os.Stat(fileName)
	switch {
	case err != nil:

		f, err = os.Create(fileName)
		if err != nil {
			return fmt.Errorf("message file create: %w", err)
		}

	default:
		f, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("message file open: %w", err)
		}
	}

	defer f.Close()

	if _, err := f.WriteString(msg + "\n"); err != nil {
		return fmt.Errorf("message file write: %w", err)
	}

	return nil
}

// =============================================================================

type docUser struct {
	ID   common.Address `json:"id"`
	Name string         `json:"name"`
}

type document struct {
	User     docUser   `json:"user"`
	Contacts []docUser `json:"contacts"`
}

func readConfig(fileName string) (document, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return document{}, fmt.Errorf("id file open: %w", err)
	}
	defer f.Close()

	var doc document
	if err := json.NewDecoder(f).Decode(&doc); err != nil {
		return document{}, fmt.Errorf("config decode: %w", err)
	}

	return doc, nil
}

func writeConfig(fileName string, doc document) error {
	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("config file create: %w", err)
	}
	defer f.Close()

	jsonDoc, err := json.MarshalIndent(doc, "", "    ")
	if err != nil {
		return fmt.Errorf("config file marshal: %w", err)
	}

	if _, err := f.Write(jsonDoc); err != nil {
		return fmt.Errorf("config file write: %w", err)
	}

	return nil
}

func createConfig(fileName string, id common.Address) (document, error) {
	filePath := filepath.Dir(fileName)

	os.MkdirAll(filePath, os.ModePerm)

	f, err := os.Create(fileName)
	if err != nil {
		return document{}, fmt.Errorf("config file create: %w", err)
	}
	defer f.Close()

	doc := document{
		User: docUser{
			ID:   id,
			Name: "Anonymous",
		},
		Contacts: []docUser{
			{
				ID:   common.Address{},
				Name: "Sample Contact",
			},
		},
	}

	jsonDoc, err := json.MarshalIndent(doc, "", "    ")
	if err != nil {
		return document{}, fmt.Errorf("config file marshal: %w", err)
	}

	if _, err := f.Write(jsonDoc); err != nil {
		return document{}, fmt.Errorf("config file write: %w", err)
	}

	return doc, nil
}
