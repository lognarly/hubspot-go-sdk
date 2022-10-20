package crm

import (
	"fmt"
)

type Emails interface {
	ListAssociations(query *EmailAssociationsQuery, emailId string, toObjectType string) (*EmailAssociations, error)
	Associate(emailId string, toObjectType string, toObjectId string, associationType string) (*Email, error)
	Disassociate(emailId string, toObjectType string, toObjectId string, associationType string) (error)
	List(query *EmailListQuery) (*EmailList, error)
	Create(options *EmailCreateOrUpdateOptions) (*Email, error)
	Read(query *EmailReadQuery, emailId string) (*Email, error)
	Update(options *EmailCreateOrUpdateOptions, emailId string) (*Email, error)
	Archive(emailId string) (error)
	BatchArchive(emailIds []string) (error)
	BatchCreate(options *EmailBatchCreateOptions) (*EmailBatchOutput, error)
	BatchRead(options *EmailBatchReadOptions) (*EmailBatchOutput, error)
	BatchUpdate(options *EmailBatchUpdateOptions) (*EmailBatchOutput, error)
	Search(options *EmailSearchOptions) (*EmailSearchResults, error)
	Merge(options *EmailMergeOptions) (*Email, error)
}

type emails struct {
	client *Client
}

type EmailAssociationsQuery struct {
	ListAssociationsQuery
}

type EmailAssociations struct {
	Results []EmailAssociation `json:"results"`
	Pagination
}

type EmailAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type EmailListQuery struct {
	ListQuery
}

type EmailList struct {
	Emails      []Email `json:"results"`
	Pagination
}

type Email struct {
	Id         string         `json:"id"`
	Properties EmailProperties `json:"properties"`
	CreatedAt  string         `json:"createdAt"`
	UpdatedAt  string         `json:"updatedAt"`
	Archived   bool           `json:"archived"`
}

type EmailProperties struct {
	CreateDate             string `json:"createdate"`
	HsEmailDirection       string `json:"hs_email_direction,omitempty"`
	HsEmailSenderEmail     string `json:"hs_email_sender_email,omitempty"`
	HsEmailSenderFirstName string `json:"hs_email_sender_firstname"`
	HsEmailSenderLastName  string `json:"hs_email_sender_lastname,omitempty"`
	HsEmailStatus          string `json:"hs_email_status,omitempty"`
	HsEmailSubject         string `json:"hs_email_subject,omitempty"`
	HsEmailText            string `json:"hs_email_text,omitempty"`
	HsEmailToEmail         string `json:"hs_email_to_email,omitempty"`
	HsEmailToFirstName     string `json:"hs_email_to_firstname,omitempty"`
	HsEmailToLastName      string `json:"hs_email_to_lastname,omitempty"`
	HsLastModifiedDate     string `json:"hs_lastmodifieddate,omitempty"`
	HsTimestamp            string `json:"hs_timestamp,omitempty"`
	HubSpotOwnerId         string `json:"hubspot_owner_id,omitempty"`
}

type EmailCreateOrUpdateOptions struct {
	Properties EmailCreateOrUpdateProperties `json:"properties"`
}

type EmailCreateOrUpdateProperties struct {
	HsEmailDirection       string `json:"hs_email_direction,omitempty"`
	HsEmailSenderEmail     string `json:"hs_email_sender_email,omitempty"`
	HsEmailSenderFirstName string `json:"hs_email_sender_firstname,omitempty"`
	HsEmailSenderLastName  string `json:"hs_email_sender_lastname,omitempty"`
	HsEmailStatus          string `json:"hs_email_status,omitempty"`
	HsEmailSubject         string `json:"hs_email_subject,omitempty"`
	HsEmailText            string `json:"hs_email_text,omitempty"`
	HsEmailToEmail         string `json:"hs_email_to_email,omitempty"`
	HsEmailToFirstName     string `json:"hs_email_to_firstname,omitempty"`
	HsEmailToLastName      string `json:"hs_email_to_lastname,omitempty"`
	HsTimestamp            string `json:"hs_timestamp,omitempty"`
	HubSpotOwnerId         string `json:"hubspot_owner_id,omitempty"`
}

type EmailReadQuery struct {
	ReadQuery
}

type EmailUpdateQuery struct {
	IdProperty string `url:"idProperty"`
}

type EmailBatchOutput struct {
	Status      string `json:"status"`
	Results     []Email `json:"results"`
	RequestedAt string `json:"requestedAt"`
	StartedAt   string `json:"startedAt"`
	CompletedAt string `json:"completedAt"`
}

type EmailBatchReadOptions struct {
	BatchReadOptions
}

type EmailBatchCreateOptions struct {
	Inputs []EmailCreateOrUpdateOptions `json:"inputs"`
}

type EmailBatchUpdateOptions struct {
	Inputs []EmailBatchUpdateProperties `json:"inputs"`
}

type EmailBatchUpdateProperties struct {
	Id         string                       `json:"id"`
	Properties EmailCreateOrUpdateProperties `json:"properties"`
}

type EmailSearchOptions struct {
	SearchOptions
}

type EmailSearchResults struct {
	Total      int64  `json:"total"`
	Results    []Email `json:"results"`
	Pagination
}

type EmailMergeOptions struct {
	MergeOptions
}

func (z *emails) ListAssociations(query *EmailAssociationsQuery, emailId string, toObjectType string) (*EmailAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/emails/%s/associations/%s", emailId, toObjectType)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.emails.ListAssociations(): newHttpRequest(): %v", err)
	}

	ca := &EmailAssociations{}

	err = z.client.do(req, ca)
	if err != nil {
		return nil, fmt.Errorf("client.emails.ListAssociations(): do(): %v", err)
	}
	
	return ca, nil
}

func (z *emails) Associate(emailId string, toObjectType string, toObjectId string, associationType string) (*Email, error) {
	u := fmt.Sprintf("/crm/v3/objects/emails/%s/associations/%s/%s/%s", emailId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("PUT", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.emails.Associate(): newHttpRequest(): %v", err)
	}

	email := &Email{}

	err = z.client.do(req, email)
	if err != nil {
		return nil, fmt.Errorf("client.emails.Associate(): do(): %v", err)
	}
	
	return email, nil
}

func (z *emails) Disassociate(emailId string, toObjectType string, toObjectId string, associationType string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/emails/%s/associations/%s/%s/%s", emailId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.emails.Disassociate(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *emails) List(query *EmailListQuery) (*EmailList, error) {
	u := fmt.Sprintf("/crm/v3/objects/emails")
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.emails.List(): newHttpRequest(): %v", err)
	}

	el := &EmailList{}

	err = z.client.do(req, el)
	if err != nil {
		return nil, fmt.Errorf("client.emails.List(): do(): %v", err)
	}
	
	return el, nil
}

func (z *emails) Create(options *EmailCreateOrUpdateOptions) (*Email, error) {
	u := fmt.Sprintf("/crm/v3/objects/emails")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.emails.Create(): newHttpRequest(): %v", err)
	}

	email := &Email{}

	err = z.client.do(req, email)
	if err != nil {
		return nil, fmt.Errorf("client.emails.Create(): do(): %v", err)
	}
	
	return email, nil
}

func (z *emails) Read(query *EmailReadQuery, emailId string) (*Email, error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems/%s", emailId)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.emails.Read(): newHttpRequest(): %v", err)
	}

	email := &Email{}

	err = z.client.do(req, email)
	if err != nil {
		return nil, fmt.Errorf("client.emails.Read(): do(): %+v", err)
	}

	return email, nil
}

func (z *emails) Update(options *EmailCreateOrUpdateOptions, emailId string) (*Email, error) {
	u := fmt.Sprintf("crm/v3/objects/emails/%s", emailId)
	req, err := z.client.newHttpRequest("PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.emails.Update(): newHttpRequest(): %v", err)
	}

	email := &Email{}

	err = z.client.do(req, email)
	if err != nil {
		return nil, fmt.Errorf("client.emails.Update(): do(): %+v", err)
	}

	return email, nil
}

func (z *emails) Archive(emailId string) (error) {
	u := fmt.Sprintf("crm/v3/objects/emails/%s", emailId)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.emails.Archive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *emails) BatchArchive(emailIds []string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/emails/batch/archive")

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, emailId := range emailIds{
		options.Inputs = append(options.Inputs, BatchInput{Id: emailId})
	}

	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return fmt.Errorf("client.emails.BatchArchive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *emails) BatchCreate(options *EmailBatchCreateOptions) (*EmailBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/emails/batch/create")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.emails.BatchCreate(): newHttpRequest(): %v", err)
	}

	emails := &EmailBatchOutput{}

	err = z.client.do(req, emails)
	if err != nil {
		return nil, fmt.Errorf("client.emails.BatchCreate(): do(): %+v", err)
	}

	return emails, nil
}

func (z *emails) BatchRead(options *EmailBatchReadOptions) (*EmailBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/emails/batch/read")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.emails.BatchUpdate(): newHttpRequest(): %v", err)
	}

	emails := &EmailBatchOutput{}

	err = z.client.do(req, emails)
	if err != nil {
		return nil, fmt.Errorf("client.emails.BatchUpdate(): do(): %+v", err)
	}

	return emails, nil
}

func (z *emails) BatchUpdate(options *EmailBatchUpdateOptions) (*EmailBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/emails/batch/update")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.emails.BatchUpdate(): newHttpRequest(): %v", err)
	}

	emails := &EmailBatchOutput{}

	err = z.client.do(req, emails)
	if err != nil {
		return nil, fmt.Errorf("client.emails.BatchUpdate(): do(): %+v", err)
	}

	return emails, nil
}

func (z *emails) Search(options *EmailSearchOptions) (*EmailSearchResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/emails/search")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.emails.Search(): newHttpRequest(): %v", err)
	}

	emails := &EmailSearchResults{}

	err = z.client.do(req, emails)
	if err != nil {
		return nil, fmt.Errorf("client.emails.Search(): do(): %+v", err)
	}

	return emails, nil
}

func (z *emails) Merge(options *EmailMergeOptions) (*Email, error) {
	u := fmt.Sprintf("/crm/v3/objects/emails/merge")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.emails.Merge(): newHttpRequest(): %v", err)
	}

	email := &Email{}

	err = z.client.do(req, email)
	if err != nil {
		return nil, fmt.Errorf("client.emails.Merge(): do(): %+v", err)
	}

	return email, nil
}