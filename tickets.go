package main

import (
    "fmt"
)

type Tickets interface {
	List(query *TicketsListQuery) (*TicketList, error)
	Create(options *TicketCreateOrUpdateOptions) (*Ticket, error)
	Read(ticketId string, query *TicketReadQuery) (*Ticket, error)
	Update(ticketId string, options *TicketCreateOrUpdateOptions) (*Ticket, error)
	Archive(ticketId string) (error)
	BatchArchive(ticketIds []string) (error)
	BatchCreate(options *TicketBatchCreateOptions) (*TicketBatchResults, error)
	BatchRead(options *TicketBatchReadOptions) (*TicketBatchResults, error)
	BatchUpdate(options *TicketBatchUpdateOptions) (*TicketBatchResults, error)
	Search(options *TicketSearchOptions) (*TicketSearchResults, error)
	Merge(options *MergeOptions) (*Ticket, error)
}

type tickets struct {
	client *Client
}

type TicketsListQuery struct {
	ListQuery
}

type TicketList struct {
	Results []Ticket `json:"results"`
	Pagination
}

type Ticket struct {
	Id         string           `json:"id"`
	Properties TicketProperties `json:"properties"`
}

type TicketProperties struct {
	CreateDate               string `json:"createdate"`
	HSLastModifiedDate       string `json:"hs_lastmodifieddate"`
	HSPipeline               string `json:"hs_pipeline"`
	HSPipelineStage          string `json:"hs_pipeline_stage"`
	HSTicketPriority         string `json:"hs_ticket_priority"`
	HubSpotOwnerId           string `json:"hubspot_owner_id"`
	Subject                  string `json:"subject"`
}

type TicketCreateOrUpdateOptions struct {
	Properties TicketCreateOrUpdateProperties `json:"properties"`
}

type TicketCreateOrUpdateProperties struct {
	HSPipeline               string `json:"hs_pipeline"`
	HSPipelineStage          string `json:"hs_pipeline_stage"`
	HSTicketPriority         string `json:"hs_ticket_priority"`
	HubSpotOwnerId           string `json:"hubspot_owner_id"`
	Subject                  string `json:"subject"`
}

type TicketReadQuery struct {
	Properties            []string `url:"properties,omitempty"`
	PropertiesWithHistory []string `url:"propertiesWithHistory,omitempty"`
	Associations          []string `url:"associations,omitempty"`
	Archived              bool     `url:"archived,omitempty"`
	idProperty            string   `url:"idProperty,omitempty"`
}

type TicketBatchCreateOptions struct {
	Inputs []TicketCreateOrUpdateOptions `json:"inputs"`
}

type TicketBatchResults struct {
	Status      string `json:"status"`
	Results     []Ticket `json:"results"`
	RequestedAt string `json:"requestedAt"`
	StartedAt   string `json:"startedAt"`
	CompletedAt string `json:"completedAt"`
}

type TicketBatchReadOptions struct {
	BatchReadOptions
}

type TicketBatchUpdateOptions struct {
	Inputs []Ticket `json:"inputs"`
}

type TicketSearchOptions struct {
	SearchOptions
}

type TicketSearchResults struct {
	Total   int64    `json:"total"`
	Results []Ticket `json:"results"`
}

func (z *tickets) List(query *TicketsListQuery) (*TicketList, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets")
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.ticket.List(): newHttpRequest(): %v", err)
	}

	tl := &TicketList{}

	err = z.client.do(req, tl)
	if err != nil {
		return nil, fmt.Errorf("client.tickets.ListAssociations(): do(): %v", err)
	}
	
	return tl, nil
}

func (z *tickets) Create(options *TicketCreateOrUpdateOptions) (*Ticket, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.ticket.Create(): newHttpRequest(): %v", err)
	}

	ticket := &Ticket{}

	err = z.client.do(req, ticket)
	if err != nil {
		return nil, fmt.Errorf("client.tickets.Create(): do(): %v", err)
	}
	
	return ticket, nil
}

func (z *tickets) Read(ticketId string, query *TicketReadQuery) (*Ticket, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/%s", ticketId)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.ticket.Read(): newHttpRequest(): %v", err)
	}

	ticket := &Ticket{}

	err = z.client.do(req, ticket)
	if err != nil {
		return nil, fmt.Errorf("client.tickets.Read(): do(): %v", err)
	}
	
	return ticket, nil
}

func (z *tickets) Update(ticketId string, options *TicketCreateOrUpdateOptions) (*Ticket, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/%s", ticketId)
	req, err := z.client.newHttpRequest("PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.ticket.Update(): newHttpRequest(): %v", err)
	}

	ticket := &Ticket{}

	err = z.client.do(req, ticket)
	if err != nil {
		return nil, fmt.Errorf("client.tickets.Update(): do(): %v", err)
	}
	
	return ticket, nil
}

func (z *tickets) Archive(ticketId string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/%s", ticketId)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.ticket.Update(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *tickets) BatchArchive(ticketIds []string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/batch/archive")

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, ticketId := range ticketIds{
		options.Inputs = append(options.Inputs, BatchInput{Id: ticketId})
	}
	
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return fmt.Errorf("client.ticket.BatchArchive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *tickets) BatchCreate(options *TicketBatchCreateOptions) (*TicketBatchResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/batch/create")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.ticket.BatchCreate(): newHttpRequest(): %v", err)
	}

	tbr := &TicketBatchResults{}

	err = z.client.do(req, tbr)
	if err != nil {
		return nil, fmt.Errorf("client.tickets.BatchCreate(): do(): %v", err)
	}
	
	return tbr, nil
}

func (z *tickets) BatchRead(options *TicketBatchReadOptions) (*TicketBatchResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/batch/read")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.ticket.BatchRead(): newHttpRequest(): %v", err)
	}

	tbr := &TicketBatchResults{}

	err = z.client.do(req, tbr)
	if err != nil {
		return nil, fmt.Errorf("client.tickets.BatchRead(): do(): %v", err)
	}
	
	return tbr, nil
}

func (z *tickets) BatchUpdate(options *TicketBatchUpdateOptions) (*TicketBatchResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/batch/update")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.ticket.BatchUpdate(): newHttpRequest(): %v", err)
	}

	tbr := &TicketBatchResults{}

	err = z.client.do(req, tbr)
	if err != nil {
		return nil, fmt.Errorf("client.tickets.BatchUpdate(): do(): %v", err)
	}
	
	return tbr, nil
}

func (z *tickets) Search(options *TicketSearchOptions) (*TicketSearchResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/search")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.ticket.Search(): newHttpRequest(): %v", err)
	}

	tsr := &TicketSearchResults{}

	err = z.client.do(req, tsr)
	if err != nil {
		return nil, fmt.Errorf("client.tickets.Search(): do(): %v", err)
	}
	
	return tsr, nil
}

func (z *tickets) Merge(options *MergeOptions) (*Ticket, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/merge")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.ticket.Merge(): newHttpRequest(): %v", err)
	}

	ticket := &Ticket{}

	err = z.client.do(req, ticket)
	if err != nil {
		return nil, fmt.Errorf("client.tickets.Merge(): do(): %v", err)
	}
	
	return ticket, nil
}