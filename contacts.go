package main

import (
	"fmt"

)

type Contacts interface {
	List(query *ContactListQuery) (*ContactList, error)
	Create(options *ContactCreateOptions) (*Contact, error)
	Read(contactId string) (*Contact, error)
	Update(contactId string, options *ContactUpdateOptions) (*Contact, error)
	Archive(contactId string) (error)
	BatchArchive(contactIds []string) (error)
	BatchCreate(options *ContactBatchCreateOptions) (*ContactBatchCreateOutput, error)
}

type contacts struct {
	client *Client
}

type ContactListQuery struct {
	ListQuery
}

type ContactList struct {
	Contacts   []Contact `json:"results"`
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

type ContactCreateOptions struct {
	Properties ContactCreateOrUpdateProperties `json:"properties"`
}

type ContactBatchCreateOutput struct {
	Status  string    `json:"status"`
	Results []Contact `json:"results"`
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

type ContactBatchCreateOptions struct {
	Inputs []ContactCreateOptions `json:"inputs"`
}

func (z *contacts) List(query *ContactListQuery) (*ContactList, error) {
	u := fmt.Sprintf("crm/v3/objects/contacts")
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.List(): newHttpRequest(): %v", err)
	}

	cl := &ContactList{}

	err = z.client.do(req, cl)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.List(): do(): %v", err)
	}
	
	return cl, nil
}

func (z *contacts) Create(options *ContactCreateOptions) (*Contact, error) {
	u := fmt.Sprintf("crm/v3/objects/contacts")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.Create(): newHttpRequest(): %+v", err)
	}

	contact := &Contact{}

	err = z.client.do(req, contact)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.Create(): do(): %+v", err)
	}

	return contact, nil
}

func (z *contacts) Read(contactId string) (*Contact, error) {
	u := fmt.Sprintf("crm/v3/objects/contacts/%s", contactId)
	req, err := z.client.newHttpRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.Read(): newHttpRequest(): %v", err)
	}

	contact := &Contact{}

	err = z.client.do(req, contact)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.Read(): do(): %+v", err)
	}

	return contact, nil
}

func (z *contacts) Update(contactId string, options *ContactUpdateOptions) (*Contact, error) {
	u := fmt.Sprintf("crm/v3/objects/contacts/%s", contactId)
	req, err := z.client.newHttpRequest("PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.Update(): newHttpRequest(): %v", err)
	}

	contact := &Contact{}

	err = z.client.do(req, contact)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.Update(): do(): %+v", err)
	}

	return contact, nil
}

func (z *contacts) Archive(contactId string) (error) {
	u := fmt.Sprintf("crm/v3/objects/contacts/%s", contactId)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.contacts.Archive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *contacts) BatchArchive(contactIds []string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/contacts/batch/archive")

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, contactId := range contactIds{
		options.Inputs = append(options.Inputs, BatchInput{Id: contactId})
	}

	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return fmt.Errorf("client.contacts.BatchArchive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *contacts) BatchCreate(options *ContactBatchCreateOptions) (*ContactBatchCreateOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/contacts/batch/create")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.BatchCreate(): newHttpRequest(): %v", err)
	}

	contacts := &ContactBatchCreateOutput{}

	err = z.client.do(req, contacts)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.BatchCreate(): do(): %+v", err)
	}

	return contacts, nil
}