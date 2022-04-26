package main

import (
	"fmt"
)

type Quotes interface {
	ListAssociations(quoteId string, toObjectType string, query *QuotesListAssociationsQuery) (*QuotesAssociationsList, error)
	List(query *QuotesListQuery) (*QuotesList, error)
	Read(quoteId string, query *QuotesReadQuery) (*Quote, error)
	BatchRead(options *QuotesBatchReadOptions) (*QuotesBatchReadResults, error)
	Search(options *QuotesSearchOptions) (*QuotesSearchResults, error)
}

type quotes struct {
	client *Client
}

type QuotesListAssociationsQuery struct {
	ListAssociationsQuery
}

type QuotesAssociationsList struct {
	Results []QuotesAssociations `json:"results"`
	Pagination
}

type QuotesAssociations struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type QuotesListQuery struct {
	ListQuery
}

type QuotesList struct {
	Results []Quote `json:"results"`
	Pagination
}

type Quote struct {
	Id         string               `json:"id,omitempty"`
	Properties QuotesListProperties `json:"properties"`
}

type QuotesListProperties struct {
	HSCreateDate     string `json:"hs_createdate"`
	HSExpirationDate string `json:"hs_expiration_date"`
	HSQuoteAmount    string `json:"hs_quote_amount"`
	HSQuoteNumber    string `json:"hs_quote_number"`
	HSStatus         string `json:"hs_status"`
	HSTerms          string `json:"hs_terms"`
	HSTitle          string `json:"hs_title"`
	HubSpotOwnerId   string `json:"hubspot_owner_id"`
}

type QuotesReadQuery struct {
	ReadQuery
}

type QuotesBatchReadQuery struct {
	BatchReadQuery
}

type QuotesBatchReadOptions struct {
	BatchReadOptions
}

type QuotesBatchReadResults struct {
	Status      string  `json:"status"`
	Results     []Quote `json:"results"`
	RequestedAt string  `json:"requestedAt"`
	StartedAt   string  `json:"startedAt"`
	CompletedAt string  `json:"completedAt"`
}

type QuotesSearchOptions struct {
	SearchOptions
}

type QuotesSearchResults struct {
	Total   int32 `json:"total"`
	Results []Quote `json:"results"`
}

func (z *quotes) ListAssociations(quoteId string, toObjectType string, query *QuotesListAssociationsQuery) (*QuotesAssociationsList, error) {
	u := fmt.Sprintf("/crm/v3/objects/quotes/%s/associations/%s", quoteId, toObjectType)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.quotes.ListAssociations(): newHttpRequest(): %v", err)
	}

	qal := &QuotesAssociationsList{}

	err = z.client.do(req, qal)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.ListAssociations(): do(): %v", err)
	}
	
	return qal, nil
}

func (z *quotes) List(query *QuotesListQuery) (*QuotesList, error) {
	u := fmt.Sprintf("/crm/v3/objects/quotes")
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.quotes.List(): newHttpRequest(): %v", err)
	}

	ql := &QuotesList{}

	err = z.client.do(req, ql)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.List(): do(): %v", err)
	}
	
	return ql, nil
}

func (z *quotes) Read(quoteId string, query *QuotesReadQuery) (*Quote, error) {
	u := fmt.Sprintf("/crm/v3/objects/quotes/%s", quoteId)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.quotes.Read(): newHttpRequest(): %v", err)
	}

	q := &Quote{}

	err = z.client.do(req, q)
	if err != nil {
		return nil, fmt.Errorf("client.quotes.Read(): do(): %v", err)
	}
	
	return q, nil
}

func (z *quotes) BatchRead(options *QuotesBatchReadOptions) (*QuotesBatchReadResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/quotes/batch/read")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.quotes.BatchRead(): newHttpRequest(): %v", err)
	}

	qbrr := &QuotesBatchReadResults{}

	err = z.client.do(req, qbrr)
	if err != nil {
		return nil, fmt.Errorf("client.quotes.BatchRead(): do(): %v", err)
	}
	
	return qbrr, nil
}

func (z *quotes) Search(options *QuotesSearchOptions) (*QuotesSearchResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/quotes/search")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.quotes.Search(): newHttpRequest(): %v", err)
	}

	qsr := &QuotesSearchResults{}

	err = z.client.do(req, qsr)
	if err != nil {
		return nil, fmt.Errorf("client.quotes.Search(): do(): %v", err)
	}
	
	return qsr, nil
}