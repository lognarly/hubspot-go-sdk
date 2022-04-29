package crm

import (
	"fmt"
)

type Contacts interface {
	ListAssociations(query *ContactAssociationsQuery, contactId string, toObjectType string) (*ContactAssociations, error)
	Associate(contactId string, toObjectType string, toObjectId string, associationType string) (*Contact, error)
	Disassociate(contactId string, toObjectType string, toObjectId string, associationType string) (error)
	List(query *ContactListQuery) (*ContactList, error)
	Create(options *ContactCreateOrUpdateOptions) (*Contact, error)
	Read(query *ContactReadQuery, contactId string) (*Contact, error)
	Update(contactId string, options *ContactCreateOrUpdateOptions) (*Contact, error)
	Archive(contactId string) (error)
	BatchArchive(contactIds []string) (error)
	BatchCreate(options *ContactBatchCreateOptions) (*ContactBatchOutput, error)
	BatchRead(options *ContactBatchReadOptions) (*ContactBatchOutput, error)
	BatchUpdate(options *ContactBatchUpdateOptions) (*ContactBatchOutput, error)
	GdprDelete(options *ContactGdprDeleteOptions) (error)
	Search(options *ContactSearchOptions) (*ContactSearchResults, error)
	Merge(options *ContactMergeOptions) (*Contact, error)
}

type contacts struct {
	client *Client
}

type ContactAssociationsQuery struct {
	ListAssociationsQuery
}

type ContactAssociations struct {
	Results []ContactAssociation `json:"results"`
	Pagination
}

type ContactAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type ContactListQuery struct {
	ListQuery
}

type ContactList struct {
	Contacts   []Contact `json:"results"`
	Pagination
}

type Contact struct {
	Id         string            `json:"id"`
	Properties ContactProperties `json:"properties"`
	CreatedAt  string            `json:"createdAt"`
	UpdatedAt  string            `json:"updatedAt"`
	Archived   bool              `json:"archived"`
}

type ContactCreateOrUpdateOptions struct {
	Properties ContactProperties `json:"properties"`
}

type ContactReadQuery struct {
	ReadQuery
}

type ContactBatchOutput struct {
	Status      string    `json:"status"`
	Results     []Contact `json:"results"`
	RequestedAt string    `json:"requestedAt"`
	StartedAt   string    `json:"startedAt"`
	CompletedAt string    `json:"completedAt"`
}

type ContactBatchReadOptions struct {
	BatchReadOptions
}

type ContactBatchCreateOptions struct {
	Inputs []ContactCreateOrUpdateOptions `json:"inputs"`
}

type ContactBatchUpdateOptions struct {
	Inputs []ContactBatchUpdateProperties `json:"inputs"`
}

type ContactBatchUpdateProperties struct {
	Id         string            `json:"id"`
	Properties ContactProperties `json:"properties"`
}

type ContactGdprDeleteOptions struct {
	ObjectId   string `json:"objectId"`
	IdProperty string `json:"idProperty"`
}

type ContactSearchOptions struct {
	SearchOptions
}

type ContactSearchResults struct {
	Total      int64 `json:"total"`
	Results    []Contact
	Pagination
}

type ContactMergeOptions struct {
	MergeOptions
}

func (z *contacts) ListAssociations(query *ContactAssociationsQuery, contactId string, toObjectType string) (*ContactAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/contacts/%s/associations/%s", contactId, toObjectType)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.ListAssociations(): newHttpRequest(): %v", err)
	}

	ca := &ContactAssociations{}

	err = z.client.do(req, ca)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.ListAssociations(): do(): %v", err)
	}
	
	return ca, nil
}

func (z *contacts) Associate(contactId string, toObjectType string, toObjectId string, associationType string) (*Contact, error) {
	u := fmt.Sprintf("/crm/v3/objects/contacts/%s/associations/%s/%s/%s", contactId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("PUT", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.Associate(): newHttpRequest(): %v", err)
	}

	contact := &Contact{}

	err = z.client.do(req, contact)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.Associate(): do(): %v", err)
	}
	
	return contact, nil
}

func (z *contacts) Disassociate(contactId string, toObjectType string, toObjectId string, associationType string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/contacts/%s/associations/%s/%s/%s", contactId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.contacts.Disassociate(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
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

func (z *contacts) Create(options *ContactCreateOrUpdateOptions) (*Contact, error) {
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

func (z *contacts) Read(query *ContactReadQuery, contactId string) (*Contact, error) {
	u := fmt.Sprintf("crm/v3/objects/contacts/%s", contactId)
	req, err := z.client.newHttpRequest("GET", u, query)
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

func (z *contacts) Update(contactId string, options *ContactCreateOrUpdateOptions) (*Contact, error) {
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

func (z *contacts) BatchCreate(options *ContactBatchCreateOptions) (*ContactBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/contacts/batch/create")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.BatchCreate(): newHttpRequest(): %v", err)
	}

	contacts := &ContactBatchOutput{}

	err = z.client.do(req, contacts)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.BatchCreate(): do(): %+v", err)
	}

	return contacts, nil
}

func (z *contacts) BatchRead(options *ContactBatchReadOptions) (*ContactBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/contacts/batch/read")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.BatchUpdate(): newHttpRequest(): %v", err)
	}

	contacts := &ContactBatchOutput{}

	err = z.client.do(req, contacts)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.BatchUpdate(): do(): %+v", err)
	}

	return contacts, nil
}

func (z *contacts) BatchUpdate(options *ContactBatchUpdateOptions) (*ContactBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/contacts/batch/update")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.BatchUpdate(): newHttpRequest(): %v", err)
	}

	contacts := &ContactBatchOutput{}

	err = z.client.do(req, contacts)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.BatchUpdate(): do(): %+v", err)
	}

	return contacts, nil
}

func (z *contacts) GdprDelete(options *ContactGdprDeleteOptions) (error) {
	u := fmt.Sprintf("/crm/v3/objects/contacts/gdpr-delete")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return fmt.Errorf("client.contacts.GdprDelete(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *contacts) Search(options *ContactSearchOptions) (*ContactSearchResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/contacts/search")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.Search(): newHttpRequest(): %v", err)
	}

	contacts := &ContactSearchResults{}

	err = z.client.do(req, contacts)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.Search(): do(): %+v", err)
	}

	return contacts, nil
}

func (z *contacts) Merge(options *ContactMergeOptions) (*Contact, error) {
	u := fmt.Sprintf("/crm/v3/objects/contacts/merge")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.Merge(): newHttpRequest(): %v", err)
	}

	company := &Contact{}

	err = z.client.do(req, company)
	if err != nil {
		return nil, fmt.Errorf("client.contacts.Merge(): do(): %+v", err)
	}

	return company, nil
}