package crm

import (
	"fmt"
)

type Quotes interface {
	ListAssociations(quoteId string, toObjectType string, query *QuoteListAssociationsQuery) (*QuoteAssociationsList, error)
	List(query *QuoteListQuery) (*QuoteList, error)
	Read(quoteId string, query *QuoteReadQuery) (*Quote, error)
	BatchRead(options *QuoteBatchReadOptions) (*QuoteBatchReadResults, error)
	Search(options *QuoteSearchOptions) (*QuoteSearchResults, error)
}

type quotes struct {
	client *Client
}

type QuoteListAssociationsQuery struct {
	ListAssociationsQuery
}

type QuoteAssociationsList struct {
	Results []QuoteAssociations `json:"results"`
	Pagination
}

type QuoteAssociations struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type QuoteListQuery struct {
	ListQuery
}

type QuoteList struct {
	Results []Quote `json:"results"`
	Pagination
}

type Quote struct {
	Id         string          `json:"id,omitempty"`
	Properties QuoteProperties `json:"properties"`
}

type QuoteProperties struct {
	HSCreateDate     string `json:"hs_createdate"`
	HSExpirationDate string `json:"hs_expiration_date,omitempty"`
	HSQuoteAmount    string `json:"hs_quote_amount,omitempty"`
	HSQuoteNumber    string `json:"hs_quote_number"`
	HSStatus         string `json:"hs_status,omitempty"`
	HSTerms          string `json:"hs_terms,omitempty"`
	HSTitle          string `json:"hs_title,omitempty"`
	HubSpotOwnerId   string `json:"hubspot_owner_id,omitempty"`
}

type QuoteReadQuery struct {
	ReadQuery
}

type QuoteBatchReadOptions struct {
	BatchReadOptions
}

type QuoteBatchReadResults struct {
	Status      string  `json:"status"`
	Results     []Quote `json:"results"`
	RequestedAt string  `json:"requestedAt"`
	StartedAt   string  `json:"startedAt"`
	CompletedAt string  `json:"completedAt"`
}

type QuoteSearchOptions struct {
	SearchOptions
}

type QuoteSearchResults struct {
	Total      int32   `json:"total"`
	Results    []Quote `json:"results"`
	Pagination
}

func (z *quotes) ListAssociations(quoteId string, toObjectType string, query *QuoteListAssociationsQuery) (*QuoteAssociationsList, error) {
	u := fmt.Sprintf("/crm/v3/objects/quotes/%s/associations/%s", quoteId, toObjectType)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.quotes.ListAssociations(): newHttpRequest(): %v", err)
	}

	qal := &QuoteAssociationsList{}

	err = z.client.do(req, qal)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.ListAssociations(): do(): %v", err)
	}
	
	return qal, nil
}

func (z *quotes) List(query *QuoteListQuery) (*QuoteList, error) {
	u := fmt.Sprintf("/crm/v3/objects/quotes")
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.quotes.List(): newHttpRequest(): %v", err)
	}

	ql := &QuoteList{}

	err = z.client.do(req, ql)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.List(): do(): %v", err)
	}
	
	return ql, nil
}

func (z *quotes) Read(quoteId string, query *QuoteReadQuery) (*Quote, error) {
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

func (z *quotes) BatchRead(options *QuoteBatchReadOptions) (*QuoteBatchReadResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/quotes/batch/read")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.quotes.BatchRead(): newHttpRequest(): %v", err)
	}

	qbrr := &QuoteBatchReadResults{}

	err = z.client.do(req, qbrr)
	if err != nil {
		return nil, fmt.Errorf("client.quotes.BatchRead(): do(): %v", err)
	}
	
	return qbrr, nil
}

func (z *quotes) Search(options *QuoteSearchOptions) (*QuoteSearchResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/quotes/search")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.quotes.Search(): newHttpRequest(): %v", err)
	}

	qsr := &QuoteSearchResults{}

	err = z.client.do(req, qsr)
	if err != nil {
		return nil, fmt.Errorf("client.quotes.Search(): do(): %v", err)
	}
	
	return qsr, nil
}