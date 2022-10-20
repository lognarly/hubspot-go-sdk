package crm

import (
    "fmt"
)

type Tickets interface {
	ListAssociations(query *TicketAssociationsQuery, ticketId string, toObjectType string) (*TicketAssociations, error)
	Associate(ticketId string, toObjectType string, toObjectId string, associationType string) (*Ticket, error)
	Disassociate(ticketId string, toObjectType string, toObjectId string, associationType string) (error)
	List(query *TicketsListQuery) (*TicketList, error)
	Create(options *TicketCreateOrUpdateOptions) (*Ticket, error)
	Read(ticketId string, query *TicketReadQuery) (*Ticket, error)
	Update(ticketId string, options *TicketCreateOrUpdateOptions) (*Ticket, error)
	Archive(ticketId string) (error)
	BatchArchive(ticketIds []string) (error)
	BatchCreate(options *TicketBatchCreateOptions) (*TicketBatchOutput, error)
	BatchRead(options *TicketBatchReadOptions) (*TicketBatchOutput, error)
	BatchUpdate(options *TicketBatchUpdateOptions) (*TicketBatchOutput, error)
	Search(options *TicketSearchOptions) (*TicketSearchResults, error)
	Merge(options *MergeOptions) (*Ticket, error)
}

type tickets struct {
	client *Client
}

type TicketAssociationsQuery struct {
	ListAssociationsQuery
}

type TicketAssociations struct {
	Results    []TicketAssociation `json:"results"`
	Pagination
}

type TicketAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
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

type TicketCreateOrUpdateOptions struct {
	Properties TicketProperties `json:"properties"`
}

type TicketReadQuery struct {
	ReadQuery
}

type TicketBatchOutput struct {
	Status      string `json:"status"`
	Results     []Ticket `json:"results"`
	RequestedAt string `json:"requestedAt"`
	StartedAt   string `json:"startedAt"`
	CompletedAt string `json:"completedAt"`
}

type TicketBatchReadOptions struct {
	BatchReadOptions
}

type TicketBatchCreateOptions struct {
	Inputs []TicketCreateOrUpdateOptions `json:"inputs"`
}

type TicketBatchUpdateOptions struct {
	Inputs []TicketBatchUpdateProperties `json:"inputs"`
}

type TicketBatchUpdateProperties struct {
	Id         string           `json:"id"`
	Properties TicketProperties `json:"properties"`
}

type TicketSearchOptions struct {
	SearchOptions
}

type TicketSearchResults struct {
	Total   int64    `json:"total"`
	Results []Ticket `json:"results"`
}

type TicketMergeOptions struct {
	MergeOptions
}

func (z *tickets) ListAssociations(query *TicketAssociationsQuery, ticketId string, toObjectType string) (*TicketAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/%s/associations/%s", ticketId, toObjectType)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.tickets.ListAssociations(): newHttpRequest(): %v", err)
	}

	ta := &TicketAssociations{}

	err = z.client.do(req, ta)
	if err != nil {
		return nil, fmt.Errorf("client.tickets.ListAssociations(): do(): %v", err)
	}
	
	return ta, nil
}

func (z *tickets) Associate(ticketId string, toObjectType string, toObjectId string, associationType string) (*Ticket, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/%s/associations/%s/%s/%s", ticketId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("PUT", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.tickets.Associate(): newHttpRequest(): %v", err)
	}

	ticket := &Ticket{}

	err = z.client.do(req, ticket)
	if err != nil {
		return nil, fmt.Errorf("client.tickets.Associate(): do(): %v", err)
	}
	
	return ticket, nil
}

func (z *tickets) Disassociate(ticketId string, toObjectType string, toObjectId string, associationType string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/%s/associations/%s/%s/%s", ticketId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.tickets.Disassociate(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
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

func (z *tickets) BatchCreate(options *TicketBatchCreateOptions) (*TicketBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/batch/create")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.ticket.BatchCreate(): newHttpRequest(): %v", err)
	}

	tbr := &TicketBatchOutput{}

	err = z.client.do(req, tbr)
	if err != nil {
		return nil, fmt.Errorf("client.tickets.BatchCreate(): do(): %v", err)
	}
	
	return tbr, nil
}

func (z *tickets) BatchRead(options *TicketBatchReadOptions) (*TicketBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/batch/read")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.ticket.BatchRead(): newHttpRequest(): %v", err)
	}

	tbr := &TicketBatchOutput{}

	err = z.client.do(req, tbr)
	if err != nil {
		return nil, fmt.Errorf("client.tickets.BatchRead(): do(): %v", err)
	}
	
	return tbr, nil
}

func (z *tickets) BatchUpdate(options *TicketBatchUpdateOptions) (*TicketBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/batch/update")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.ticket.BatchUpdate(): newHttpRequest(): %v", err)
	}

	tbr := &TicketBatchOutput{}

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