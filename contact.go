package main

import (
	"fmt"

)

type Contacts interface {
	List(options *ContactListOptions) (*ContactList, error)
	Create(options *ContactCreateOptions) (*Contact, error)
	Read(contactId string) (*Contact, error)
	Update(contactId string, options *ContactUpdateOptions) (*Contact, error)
	Delete(contactId string) (error)
	BatchArchive(options *ContactBatchArchiveOptions) (error)
	BatchCreate(options *ContactBatchCreateOptions) (*ContactBatchCreateOutput, error)
}

type contacts struct {
	client *Client
}

type ContactQuery string

type ContactListOptions struct {
	ListOptions
	Properties  ContactQuery `url:"properties,omitempty"`
}

type ContactList struct {
	Contacts []Contact `json:"results"`
	Pagination
}

type Contact struct {
	ID         string            `json:"id"`
	Properties ContactProperties `json:"properties"`
	CreatedAt  string            `json:"createdAt"`
	UpdatedAt  string            `json:"updatedAt"`
	Archived   bool              `json:"archived"`
}

type ContactProperties struct {
	Company          string   `json:"company,omitempty"`
	CreateDate       string   `json:"createdate"`
	Email            string   `json:"email"`
	FirstName        string   `json:"firstname,omitempty"`
	HSObjectID       string   `json:"hs_object_id"`
	LastModifiedDate string   `json:"lastmodifieddate"`
	LastName         string   `json:"lastname,omitempty"`
	Phone            string   `json:"phone,omitempty"`
	Website          string   `json:"website,omitempty"`
}

type ContactBatchCreateOutput struct {
	Status  string    `json:"status"`
	Results []Contact `json:"results"`
}

type ContactCreateOptions struct {
	Properties ContactCreateOrUpdateProperties `json:"properties"`
}

type ContactUpdateOptions struct {
	Properties ContactCreateOrUpdateProperties `json:"properties"`
}

type ContactCreateOrUpdateProperties struct {
	Company          string   `json:"company,omitempty"`
	Email            string   `json:"email"`
	FirstName        string   `json:"firstname,omitempty"`
	LastName         string   `json:"lastname,omitempty"`
	Phone            string   `json:"phone,omitempty"`
	Website          string   `json:"website,omitempty"`
}

type ContactBatchArchiveOptions struct {
	Inputs []ArchiveContacts `json:"inputs"`
}

type ArchiveContacts struct {
	ContactId string `json:"id"`
}

type ContactBatchCreateOptions struct {
	Inputs []ContactCreateOptions `json:"inputs"`
}

func (s *contacts) List(options *ContactListOptions) (*ContactList, error) {
	u := fmt.Sprintf("crm/v3/objects/contacts")
	req, err := s.client.newHttpRequest("GET", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.contact.List(): newHttpRequest(): %v", err)
	}

	cl := &ContactList{}

	err = s.client.do(req, cl)
	if err != nil {
		return nil, fmt.Errorf("client.contact.List(): do(): %v", err)
	}
	
	return cl, nil
}

func (s *contacts) Create(options *ContactCreateOptions) (*Contact, error) {
	u := fmt.Sprintf("crm/v3/objects/contacts")
	req, err := s.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.contact.Create(): newHttpRequest(): %+v", err)
	}

	contact := &Contact{}

	err = s.client.do(req, contact)
	if err != nil {
		return nil, fmt.Errorf("client.contact.Create(): do(): %+v", err)
	}

	return contact, nil
}

func (s *contacts) Read(contactId string) (*Contact, error) {
	u := fmt.Sprintf("crm/v3/objects/contacts/%s", contactId)
	req, err := s.client.newHttpRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.contact.Read(): newHttpRequest(): %v", err)
	}

	contact := &Contact{}

	err = s.client.do(req, contact)
	if err != nil {
		return nil, fmt.Errorf("client.contact.Read(): do(): %+v", err)
	}

	return contact, nil
}

func (s *contacts) Update(contactId string, options *ContactUpdateOptions) (*Contact, error) {
	u := fmt.Sprintf("crm/v3/objects/contacts/%s", contactId)
	req, err := s.client.newHttpRequest("PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.contact.Update(): newHttpRequest(): %v", err)
	}

	contact := &Contact{}

	err = s.client.do(req, contact)
	if err != nil {
		return nil, fmt.Errorf("client.contact.Update(): do(): %+v", err)
	}

	return contact, nil
}

func (s *contacts) Delete(contactId string) (error) {
	u := fmt.Sprintf("crm/v3/objects/contacts/%s", contactId)
	req, err := s.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.contact.Delete(): newHttpRequest(): %v", err)
	}

	return s.client.do(req, nil)
}

func (s *contacts) BatchArchive(options *ContactBatchArchiveOptions) (error) {
	u := fmt.Sprintf("/crm/v3/objects/contacts/batch/archive")
	req, err := s.client.newHttpRequest("POST", u, options)
	if err != nil {
		return fmt.Errorf("client.contact.BatchArchive(): newHttpRequest(): %v", err)
	}

	return s.client.do(req, nil)
}

func (s *contacts) BatchCreate(options *ContactBatchCreateOptions) (*ContactBatchCreateOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/contacts/batch/create")
	req, err := s.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.contact.BatchCreate(): newHttpRequest(): %v", err)
	}

	contacts := &ContactBatchCreateOutput{}

	err = s.client.do(req, contacts)
	if err != nil {
		return nil, fmt.Errorf("client.contact.BatchCreate(): do(): %+v", err)
	}

	return contacts, nil
}